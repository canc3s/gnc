package runner

import (
	"github.com/canc3s/gnc/internal/rc4"
	"github.com/canc3s/gnc/internal/tcp"
	"net"
)

func CipWorker (conn net.Conn, options *Options) {
	CipherConn,_ := rc4.NewCipherConn(conn, options.Pass)
	tcp.CipTransferStreams(CipherConn, options.Exec)
}

func Process(options *Options)  {
	if options.Listen {
		conn := tcp.StartServer(options.Proto, options.Port)
		if options.Security {
			CipWorker(conn, options)
			return
		}
		tcp.TransferStreams(conn, options.Exec)
		return
	}
	conn := tcp.StartClient(options.Proto, options.Addr)
	if options.Security {
		CipWorker(conn, options)
		return
	}
	tcp.TransferStreams(conn,options.Exec)
	return
}
