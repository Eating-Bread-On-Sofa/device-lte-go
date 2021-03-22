package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

var sendData = make(map[string]int)

func process(conn net.Conn){
	defer conn.Close()
	for{
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recv := buf[:n]
		// fmt.Println("发送的数据值：",recv)
		send := string(recv)
		flag := recv[0] == 0xFE
		if flag {
			tem1 := int(recv[3])
			tem2 := int(recv[4])
			tem3 := int(recv[5])
			tem4 := int(recv[6])
			tem5 := int(recv[7])
			tem6 := int(recv[8])
			tem := tem1 << 40 | tem2 << 32 | tem3 << 24 | tem4 << 16 | tem5 << 8 | tem6
			temp := fmt.Sprintf("%X",tem)
			Id := string(temp)
			fmt.Println("设备ID为:",Id)
			tem7 := int(recv[16])
			tem8 := int(recv[17])
			tem9 := int(recv[18])
			tem10 := int(recv[19])
			value := tem7 << 24 | tem8 << 16 | tem9 << 8 | tem10
			sendData[Id] = value
			fmt.Println(sendData[Id])
			_, err = conn.Write([]byte("ok"))
			if err != nil{
				fmt.Println("发送数据失败，错误：",err)
				continue
			}
		}else if send == "Light"{
			_, err = conn.Write([]byte("ok"))
			if err != nil{
				fmt.Println("发送数据失败，错误：",err)
				continue
			}
		}else {
			str := strconv.Itoa(sendData[send])
			backData := []byte(str)
			_, err = conn.Write(backData)
			if err != nil{
				fmt.Println("发送数据失败，错误：",err)
				continue
			}
		}
	}
}

func main(){
	listen,err := net.Listen("tcp","0.0.0.0:8891")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n",err)
		return
	}
	for{
		conn,err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		go process(conn)
	}
}