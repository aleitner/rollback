package udp

import (
	"errors"
	"net"
)

type Client struct {
	udpAddr *net.UDPAddr
	conn *net.UDPConn
}

func NewClient(address string) (*Client, error) {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		return nil, err
	}

	return &Client{
		udpAddr: udpAddr,
		conn: conn,
	}, nil
}

func (c* Client) Write(b []byte) (int, error) {
	if c.conn == nil {
		return 0, errors.New("No connection. cannot write to server")
	}

	return c.conn.Write(b)
}

func (c* Client) Close() (error) {
	if c.conn == nil {
		return nil
	}

	return c.conn.Close()
}