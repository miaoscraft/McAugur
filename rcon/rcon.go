package rcon

import "github.com/Tnze/go-mc/net"

var (
	conn         net.RCONClientConn
	addr, passwd string
)

//Open a rcon connection
func Open(addr1, passwd1 string) error {
	addr, passwd = addr1, passwd1
	conn, err := net.DialRCON(addr, passwd)
	if err != nil {
		return err
	}
	conn.(*net.RCONConn).Close()
	return nil
}

//Cmd run a minecraft command on the server
func Cmd(cmd string) (resp string, err error) {
	if cmd == "" {
		return
	}
	conn, err = net.DialRCON(addr, passwd)
	if err != nil {
		return
	}
	if err = conn.Cmd(cmd); err != nil {
		return
	}
	resp, err = conn.Resp()
	return
}
