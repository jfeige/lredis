package lredis

import (
	"net"
	"bufio"
	"errors"
)

var (
	CRLF = "\r\n"
)
type Conn struct {
	address string
	conn    net.Conn
	dst     []byte
	br 		*bufio.Reader
	bw 		*bufio.Writer
	pending int
}



func NewConn(address string,pwd ...string)(*Conn,error){
	c,err := net.Dial("tcp",address)
	if err != nil{
		return nil,err
	}

	conn := &Conn{
		address:address,
		conn:c,
		dst:make([]byte,0),
		br:bufio.NewReader(c),
		bw:bufio.NewWriter(c),
		pending:0,
	}

	if len(pwd) > 0 &&  pwd[0] != ""{
		auth,err := conn.Do("AUTH",pwd[0])
		if err != nil{
			return nil,err
		}
		if auth.(string) != "OK"{
			return nil,errors.New("ERR invalid password")
		}
	}

	return conn,nil
}