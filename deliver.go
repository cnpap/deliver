package deliver

import (
	"io"
	"net"
	"strconv"
)

// Deliver >> A courier
type Deliver struct {
	Addr string
}

type server interface {
}

// Context >> Request Context
type Context struct {
}

// Middleware >> Func type
type Middleware func(*Context) bool

// Use >> Add middleware
func (_d *Deliver) Use(_m *Middleware) {
}

// Listen >> Start listening port
func (_d *Deliver) Listen() error {
	listen, err := net.Listen("tcp", _d.Addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		go func() {
			var body [1024]byte
			_, err := conn.Read(body[:])
			if err != nil {
				return
			}
			if body[0] == 0x04 {
				host := net.IPv4(body[4], body[5], body[6], body[7]).String()
				port := strconv.Itoa(int(body[2])<<8 | int(body[3]))
				request, err := net.Dial("tcp", net.JoinHostPort(host, port))
				if err != nil {
					return
				}
				conn.Write([]byte{0x00, 0x5a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
				go func() {
					io.Copy(conn, request)
					conn.Close()
					request.Close()
				}()
				io.Copy(request, conn)
				conn.Close()
				request.Close()
			}
		}()
	}
}
