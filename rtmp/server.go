package rtmp

import (
	"github.com/q191201771/lal/log"
	"net"
)

type ServerObserver interface {
	NewRTMPPubSessionCB(session *ServerSession)
	NewRTMPSubSessionCB(session *ServerSession)
}

type Server struct {
	obs  ServerObserver
	addr string
	ln   net.Listener
}

func NewServer(obs ServerObserver, addr string) *Server {
	return &Server{
		obs:  obs,
		addr: addr,
	}
}

func (server *Server) RunLoop() error {
	var err error
	server.ln, err = net.Listen("tcp", server.addr)
	if err != nil {
		return err
	}
	log.Infof("start rtmp listen. addr=%s", server.addr)
	for {
		conn, err := server.ln.Accept()
		if err != nil {
			return err
		}
		go server.handleConnect(conn)
	}
}

func (server *Server) Dispose() {
	if err := server.ln.Close(); err != nil {
		log.Error(err)
	}
}

func (server *Server) handleConnect(conn net.Conn) {
	log.Infof("accept a rtmp connection. remoteAddr=%v", conn.RemoteAddr())
	session := NewServerSession(server, conn)
	// TODO chef
	session.RunLoop()
}

func (server *Server) NewRTMPPubSessionCB(session *ServerSession) {
	server.obs.NewRTMPPubSessionCB(session)
}

func (server *Server) NewRTMPSubSessionCB(session *ServerSession) {
	server.obs.NewRTMPSubSessionCB(session)
}

func (server *Server) ReadAVMessageCB(t int, timestampAbs int, message []byte) {

}
