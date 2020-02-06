package rcon

import "github.com/Tnze/go-mc/net"

var (
	conn         net.RCONClientConn
	Addr, Passwd string
)

//Open a rcon connection
func Open(addr, passwd string) (err error) {
	Addr, Passwd = addr, passwd
	conn, err = net.DialRCON(addr, passwd)
	if err != nil {
		return
	}
	conn.(*net.RCONConn).Close()
	return
}

//Cmd run a minecraft command on the server
func Cmd(cmd string) (resp string, err error) {
	if cmd == "" {
		return
	}
	conn, err = net.DialRCON(Addr, Passwd)
	if err != nil {
		return
	}
	if err = conn.Cmd(cmd); err != nil {
		return
	}
	resp, err = conn.Resp()
	return
}
