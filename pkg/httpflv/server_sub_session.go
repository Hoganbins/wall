// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpflv

import (
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/q191201771/naza/pkg/connection"

	"github.com/q191201771/naza/pkg/nazalog"
	"github.com/q191201771/naza/pkg/unique"
)

var flvHTTPResponseHeaderStr = "HTTP/1.1 200 OK\r\n" +
	"Cache-Control: no-cache\r\n" +
	"Content-Type: video/x-flv\r\n" +
	"Connection: close\r\n" +
	"Expires: -1\r\n" +
	"Pragma: no-cache\r\n" +
	"\r\n"

var flvHTTPResponseHeader = []byte(flvHTTPResponseHeaderStr)

type SubSession struct {
	UniqueKey string

	StartTick  int64
	StreamName string
	AppName    string
	URI        string
	Headers    map[string]string

	IsFresh bool

	conn connection.Connection
}

func NewSubSession(conn net.Conn) *SubSession {
	uk := unique.GenUniqueKey("FLVSUB")
	nazalog.Infof("lifecycle new SubSession. [%s] remoteAddr=%s", uk, conn.RemoteAddr().String())
	return &SubSession{
		UniqueKey: uk,
		IsFresh:   true,
		conn: connection.New(conn, func(option *connection.Option) {
			option.ReadBufSize = readBufSize
			option.WriteChanSize = wChanSize
			option.WriteTimeoutMS = subSessionWriteTimeoutMS
		}),
	}
}

// TODO chef: read request timeout
func (session *SubSession) ReadRequest() (err error) {
	session.StartTick = time.Now().Unix()

	defer func() {
		if err != nil {
			session.Dispose()
		}
	}()

	var (
		requestLine string
		method      string
	)
	if requestLine, session.Headers, err = parseHTTPHeader(session.conn); err != nil {
		return
	}
	if method, session.URI, _, err = parseRequestLine(requestLine); err != nil {
		return
	}
	if method != "GET" {
		err = ErrHTTPFLV
		return
	}

	var urlObj *url.URL
	if urlObj, err = url.Parse(session.URI); err != nil {
		return
	}
	if !strings.HasSuffix(urlObj.Path, ".flv") {
		err = ErrHTTPFLV
		return
	}

	items := strings.Split(urlObj.Path, "/")
	if len(items) != 3 {
		err = ErrHTTPFLV
		return
	}
	session.AppName = items[1]
	items = strings.Split(items[2], ".")
	if len(items) < 2 {
		err = ErrHTTPFLV
		return
	}
	session.StreamName = items[0]

	return nil
}

func (session *SubSession) RunLoop() error {
	buf := make([]byte, 128)
	_, err := session.conn.Read(buf)
	return err
}

func (session *SubSession) WriteHTTPResponseHeader() {
	nazalog.Infof("<----- http response header. [%s]", session.UniqueKey)
	session.WriteRawPacket(flvHTTPResponseHeader)
}

func (session *SubSession) WriteFLVHeader() {
	nazalog.Infof("<----- http flv header. [%s]", session.UniqueKey)
	session.WriteRawPacket(FLVHeader)
}

func (session *SubSession) WriteTag(tag *Tag) {
	session.WriteRawPacket(tag.Raw)
}

func (session *SubSession) WriteRawPacket(pkt []byte) {
	_, _ = session.conn.Write(pkt)
}

func (session *SubSession) Dispose() {
	_ = session.conn.Close()
}
