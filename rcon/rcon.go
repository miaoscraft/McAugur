package rcon

import "github.com/Tnze/go-mc/net"

var conn net.RCONClientConn

//Open a rcon connection
func Open(addr, passwd string) (err error) {
	conn, err = net.DialRCON(addr, passwd)
	return
}

//Cmd run a minecraft command on the server
func Cmd(cmd string) (resp string, err error) {
	if cmd == "" {
		return "", nil
	}
	if err = conn.Cmd(cmd); err != nil {
		return
	}
	resp, err = conn.Resp()
	return
}
