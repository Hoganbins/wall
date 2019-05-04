package httpflv

import (
	"bufio"
	"fmt"
	"github.com/q191201771/lal/log"
	"github.com/q191201771/lal/util"
	"io"
	"net"
	"net/url"
	"strings"
	"sync"
	"time"
)

var flvHeaderSize = 13

var flvPrevTagFieldSize = 4

type PullSessionStat struct {
	readCount int64
	readByte  int64
}

type PullSession struct {
	StartTick int64

	obs  PullSessionObserver
	Conn *net.TCPConn // after Connect success, can direct visit net.TCPConn, useful for set socket options.
	rb   *bufio.Reader

	stat      PullSessionStat
	prevStat  PullSessionStat
	statMutex sync.Mutex

	closeOnce sync.Once

	UniqueKey string
}

type PullSessionObserver interface {
	ReadHttpRespHeaderCb()
	ReadFlvHeaderCb(flvHeader []byte)
	ReadTagCb(tag *Tag) // after cb, PullSession won't use this tag data
}

func NewPullSession(obs PullSessionObserver) *PullSession {
	uk := util.GenUniqueKey("FLVPULL")
	log.Infof("lifecycle new PullSession. [%s]", uk)
	return &PullSession{
		obs:       obs,
		UniqueKey: uk,
	}
}

// @param timeout: timeout for connect operate. if 0, then no timeout
func (session *PullSession) Connect(rawurl string, timeout time.Duration) error {
	session.StartTick = time.Now().Unix()

	url, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	if url.Scheme != "http" || !strings.HasSuffix(url.Path, ".flv") {
		return fxxkErr
	}

	host := url.Host
	// TODO chef: uri with url.RawQuery?
	uri := url.Path

	var addr string
	if strings.Contains(host, ":") {
		addr = host
	} else {
		addr = host + ":80"
	}

	var conn net.Conn
	if timeout == 0 {
		conn, err = net.Dial("tcp", addr)
	} else {
		conn, err = net.DialTimeout("tcp", addr, timeout)
	}
	if err != nil {
		return err
	}
	session.Conn = conn.(*net.TCPConn)
	session.rb = bufio.NewReaderSize(session.Conn, readBufSize)

	_, err = fmt.Fprintf(session.Conn,
		"GET %s HTTP/1.0\r\nAccept: */*\r\nRange: byte=0-\r\nConnection: close\r\nHost: %s\r\nIcy-MetaData: 1\r\n\r\n",
		uri, host)
	if err != nil {
		return err
	}

	return nil
}

func (session *PullSession) RunLoop() error {
	err := session.runReadLoop()
	session.Dispose(err)
	return err
}

func (session *PullSession) Dispose(err error) {
	session.closeOnce.Do(func() {
		log.Infof("lifecycle dispose PullSession. [%s] reason=%v", session.UniqueKey, err)
		if err := session.Conn.Close(); err != nil {
			log.Error("conn close error. [%s] err=%v", session.UniqueKey, err)
		}
	})
}

func (session *PullSession) GetStat() (now PullSessionStat, diff PullSessionStat) {
	session.statMutex.Lock()
	defer session.statMutex.Unlock()
	now = session.stat
	diff.readCount = session.stat.readCount - session.prevStat.readCount
	diff.readByte = session.stat.readByte - session.prevStat.readByte
	session.prevStat = session.stat
	return
}

func (session *PullSession) runReadLoop() error {
	if err := session.readHttpRespHeader(); err != nil {
		return err
	}
	session.obs.ReadHttpRespHeaderCb()

	flvHeader, err := session.readFlvHeader()
	if err != nil {
		return err
	}
	session.obs.ReadFlvHeaderCb(flvHeader)

	for {
		tag, err := session.readTag()
		if err != nil {
			return err
		}
		session.obs.ReadTagCb(tag)
	}
}

func (session *PullSession) readHttpRespHeader() error {
	firstLine, headers, err := parseHttpHeader(session.rb)
	if err != nil {
		return err
	}

	if !strings.Contains(firstLine, "200") || len(headers) == 0 {
		return fxxkErr
	}
	log.Infof("-----> http response header. [%s]", session.UniqueKey)

	return nil
}

func (session *PullSession) readFlvHeader() ([]byte, error) {
	flvHeader := make([]byte, flvHeaderSize)
	_, err := io.ReadAtLeast(session.rb, flvHeader, flvHeaderSize)
	if err != nil {
		return flvHeader, err
	}
	log.Infof("-----> http flv header. [%s]", session.UniqueKey)

	// TODO chef: check flv header's value
	return flvHeader, nil
}

func (session *PullSession) readTag() (*Tag, error) {
	header, rawHeader, err := readTagHeader(session.rb)
	if err != nil {
		return nil, err
	}
	session.addStat(tagHeaderSize)

	needed := int(header.DataSize) + flvPrevTagFieldSize
	tag := &Tag{}
	tag.Header = header
	tag.Raw = make([]byte, tagHeaderSize+needed)
	copy(tag.Raw, rawHeader)

	if _, err := io.ReadAtLeast(session.rb, tag.Raw[tagHeaderSize:], needed); err != nil {
		return nil, err
	}
	session.addStat(needed)

	return tag, nil
}

func (session *PullSession) addStat(readByte int) {
	session.statMutex.Lock()
	defer session.statMutex.Unlock()
	session.stat.readByte += int64(readByte)
	session.stat.readCount++
}
