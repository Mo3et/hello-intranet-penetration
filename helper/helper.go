package helper

import (
	"log"
	"net"
	"time"

	"github.com/mo3et/hello-intranet-penetration/define"
)

// CreateListen
func CreateListen(serverAddr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListen, nil
}

// Create Conn
func CreateConn(serverAddr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return conn, err
}

// KeepAlive 设置连接保活
func KeepAlive(conn *net.TCPConn) {
	for {
		_, err := conn.Write([]byte(define.KeepAliveStr))
		if err != nil {
			log.Printf("[KeepAlive] Error %s", err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}

// Get Data from Connection 获取Connection中的数据
func GetDataFromConnection(bufSize int, conn *net.TCPConn) ([]byte, error) {
	b := make([]byte, 0)
	for {
		buf := make([]byte, bufSize)
		n, err := conn.Read(buf)
		if err != nil {
			return nil, err
		}
		b = append(b, buf...)
		if n < bufSize {
			break
		}

	}
	return b, nil
}
