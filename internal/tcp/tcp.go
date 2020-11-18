package tcp

import (
	"fmt"
	"github.com/canc3s/gnc/internal/gologger"
	"github.com/canc3s/gnc/internal/rc4"
	"io"
	"net"
	"os"
	"os/exec"
)

type ExecResult struct {
	Cmd			*exec.Cmd
	Stdin 		io.WriteCloser
	Stdout 		io.ReadCloser
	Stderr		io.ReadCloser
}

type Progress struct {
	bytes uint64
}

func CipTransferStreams(con *rc4.CipherConn,execs string) {
	c := make(chan Progress)

	// Read from Reader and write to Writer until EOF
	copy := func(r io.ReadCloser, w io.WriteCloser) {
		defer func() {
			r.Close()
			w.Close()
		}()
		n, err := io.Copy(w, r)
		if err != nil {
			gologger.Infof("[%s]: ERROR: %s\n", con.RemoteAddr(), err)
		}
		c <- Progress{bytes: uint64(n)}
	}

	if execs != "" {
		var execr ExecResult
		Exec(execs, &execr)
		go copy(con.Rwc, execr.Stdin)
		go copy(execr.Stdout, con.Rwc)
		go copy(execr.Stderr, con.Rwc)
	} else {
		go copy(con.Rwc, os.Stdout)
		go copy(os.Stdin, con.Rwc)
	}

	p := <-c
	gologger.Printf("[%s]: Connection has been closed by remote peer, %d bytes has been received\n", con.RemoteAddr(), p.bytes)
}


func TransferStreams(con net.Conn,execs string) {
	c := make(chan Progress)

	// Read from Reader and write to Writer until EOF
	copy := func(r io.ReadCloser, w io.WriteCloser) {
		defer func() {
			r.Close()
			w.Close()
		}()
		n, err := io.Copy(w, r)
		if err != nil {
			gologger.Infof("[%s]: ERROR: %s\n", con.RemoteAddr(), err)
		}
		c <- Progress{bytes: uint64(n)}
	}



	if execs != "" {
		var execr ExecResult
		Exec(execs, &execr)
		go copy(con, execr.Stdin)
		go copy(execr.Stdout, con)
		go copy(execr.Stderr, con)
	} else {
		go copy(con, os.Stdout)
		go copy(os.Stdin, con)
	}

	p := <-c
	gologger.Printf("[%s]: Connection has been closed by remote peer, %d bytes has been received\n", con.RemoteAddr(), p.bytes)
}

func StartServer(proto string, port int) net.Conn {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen(proto, addr)
	if err != nil {
		gologger.Fatalf("%s\n", err)
	}
	gologger.Printf("Listening on %s\n", proto, port)
	con, err := ln.Accept()
	if err != nil {
		gologger.Fatalf("%s\n", err)
	}
	gologger.Printf("[%s]: Connection has been opened\n", con.RemoteAddr())
	return con
}

func StartClient(proto string, addr string) net.Conn {
	con, err := net.Dial(proto, addr)
	if err != nil {
		gologger.Fatalf("%s\n", err)
	}
	gologger.Printf("Connected to %s\n", addr)
	return con
}

func Exec(execs string, execr *ExecResult){
	var err error
	execr.Cmd = exec.Command(execs)
	execr.Stdin, err = execr.Cmd.StdinPipe()
	if err != nil {
		gologger.Fatalf("%s\n", err)
	}
	execr.Stdout, err = execr.Cmd.StdoutPipe()
	if err != nil {
		gologger.Fatalf("%s\n", err)
	}
	execr.Stderr, err = execr.Cmd.StderrPipe()
	if err != nil {
		gologger.Fatalf("%s\n", err)
	}
	if err := execr.Cmd.Start(); err != nil {
		gologger.Fatalf("%s\n", err)
	}
}