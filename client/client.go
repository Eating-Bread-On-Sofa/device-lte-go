package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp","127.0.0.1:8891")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
	}
	data := []byte{0xFE,0xDC,0x01,0x17,0x1C,0xB7,0x40,0x84,0x63,0x00,0x00,0x00,0x00,0x03,0x00,0x10,0x00,0x00,0x00,0x2C,0x00,0x00,0x00,0x88,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x14,0x00}
	_, e := conn.Write(data)
	if e != nil {
		fmt.Printf("send failed, err:%v\n", e)
		return
	}
	var buf [1024]byte
	n,er := conn.Read(buf[:])
	if er != nil{
		fmt.Printf("read failed, err:%v\n", e)
		return
	}
	fmt.Println(string(buf[:n]))
}