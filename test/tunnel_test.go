package test

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"strconv"
	"testing"
)

// 作为中间商 处理 Client 和 server 之间的转发。
//  例如 Client 发送请求到 server, server 再返回响应给 Client

const (
	serverAddr = "0.0.0.0:22310"
	BufSize    = 10
	tunnelAddr = "0.0.0.0:22311"
)

// server
func TestServer(t *testing.T) {
	// Get the parsed address
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		t.Fatal(err)
	}
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}
	for {
		// when find new Connect
		tcpConn, err := tcpListen.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}
		b := make([]byte, 0)
		fmt.Println(reflect.TypeOf(b))
		// Read Data (need Buffer)
		for {
			var buf [bufSize]byte
			// buf := make([]byte, 1024)
			n, err := tcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			// Write data about buf to n
			b = append(b, buf[:n]...)
			// t.Log(string(buf[:n]))
			if n < bufSize {
				break
			}
		}

		// Write data
		// tcpConn.Write([]byte("Hello world,世界"))
		i, err := strconv.Atoi(string(b))
		if err != nil {
			t.Fatal(err)
		}
		i += 2
		if _, err := tcpConn.Write([]byte(strconv.Itoa(i))); err != nil {
			t.Fatal(err)
		}

	}
}

// client
func TestClient(t *testing.T) {
	// tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	tcpAddr, err := net.ResolveTCPAddr("tcp", tunnelAddr)
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
	if _, err := tcpConn.Write([]byte("1500")); err != nil {
		t.Fatal(err)
	}

	// Read Data
	b := make([]byte, 0)
	for {
		var buf [bufSize]byte
		// buf := make([]byte, 1024)
		n, err := tcpConn.Read(buf[:])
		if err != nil {
			t.Fatal(err)
		}
		b = append(b, buf[:n]...)
		if n < bufSize {
			break
		}
	}
	fmt.Println(string(b))
}

// tunnel

func TestTunnel(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", tunnelAddr)
	if err != nil {
		t.Fatal(err)
	}
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}
	for {
		// client tcp Conn
		clientTcpConn, err := tcpListen.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}

		// get user send data(client)

		// // bufio.NewReader(clientTcpConn)
		// b := make([]byte, 0) //b is client send data
		// for {
		// 	var buf [BufSize]byte
		// 	n, err := clientTcpConn.Read(buf[:])
		// 	if err != nil {
		// 		t.Fatal(err)
		// 	}
		// 	b = append(b, buf[:n]...)
		// 	if n < bufSize {
		// 		break
		// 	}
		// }

		// Server Conn
		addr, err := net.ResolveTCPAddr("tcp", serverAddr)
		if err != nil {
			t.Fatal(err)
		}

		serverTcpConn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			t.Fatal(err)
		}

		// if _, err = serverTcpConn.Write(b); err != nil { //write client data
		// 	t.Fatal(err)
		// }

		// get server response data

		// b2 := make([]byte, 0)
		// for {
		// 	var buf [bufSize]byte
		// 	n, err := serverTcpConn.Read(buf[:])
		// 	if err != nil {
		// 		t.Fatal(err)
		// 	}
		// 	b2 = append(b2, buf[:n]...)
		// 	if n < bufSize {
		// 		break
		// 	}
		// }
		// if _, err := clientTcpConn.Write(b2); err != nil {
		// 	t.Fatal(err)
		// }

		// Data exchange(Swap)
		go func() {
			if _, err := io.Copy(serverTcpConn, clientTcpConn); err != nil {
				panic(err)
			}
		}()
		go func() {
			if _, err := io.Copy(clientTcpConn, serverTcpConn); err != nil {
				panic(err)
			}

		}()
	}
}
