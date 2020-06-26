package udp

import (
	"net"
)

type Server struct {
	udpAddr *net.UDPAddr
	conn *net.UDPConn
}

func NewServer(address string) (*Server, error) {

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}

	return &Server{
		udpAddr: udpAddr,
		conn: conn,
	}, nil
}

func (s *Server) Read(b []byte) (int, error) {
	return s.conn.Read(b)
}

func (s* Server) Close() (error) {
	if s.conn == nil {
		return nil
	}

	return s.conn.Close()
}