package test

import (
	"net"
	"testing"
)

// # TCP Create Listen
// - parse addr and store to variable(tcpAddr)
// - create tcp listen, use `listenTCP(network,addr)`
// - use loop to provide function(read/write data)
// - Tips: Read data need to create buffer to store data, use `tcpConn.Read(buf[:])` Assign to n. use `t.Log(string(buf[:n]))` to print.

// # TCP Create Conn
// - parse addr and store to variable(tcpAddr)
// - use `DialTCP("tcp", nil, tcpAddr)` to create `tcpConn`
// - use loop to provide function (First write, second Read data)
// - Same Tips.

const (

	// addr="0.0.0.0:22222"
	addr    = "127.0.0.1:22222"
	bufSize = 10
)

// listen (Server)
func TestTcpListen(t *testing.T) {
	// Get the parsed address
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}
	for {
		tcpConn, err := tcpListen.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}

		// Read Data (need Buffer)
		for {
			var buf [bufSize]byte
			// buf := make([]byte, 1024)
			n, err := tcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(buf[:n]))
			if n < bufSize {
				break
			}
		}

		// Write data
		// tcpConn.Write([]byte("Hello world,世界"))
		if _, err := tcpConn.Write([]byte("Hello world,世界")); err != nil {
			t.Fatal(err)
		}

	}

}

// Create Connect (Client)
func TestCreateTcp(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	// Use dial package create TCP connect
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	// Write data
	// tcpConn.Write([]byte("Client ==> hello world, 世界！"))
	if _, err := tcpConn.Write([]byte("Client ==> hello world, 世界！")); err != nil {
		t.Fatal(err)
	}

	// Read Data
	for {
		var buf [bufSize]byte
		// buf := make([]byte, 1024)
		n, err := tcpConn.Read(buf[:])
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(buf[:n]))
		if n < bufSize {
			break
		}
	}
}
