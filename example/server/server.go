package main

import (
	"github.com/devfeel/connx"
	"fmt"
	"encoding/gob"
	"time"
)

type LoginInfo struct{
	UserName string
	Password string
	LoginFrom string
}

func init(){
	gob.Register(LoginInfo{})
	connx.SetHeadFlag(0x1000)
}

func main(){

	server, err := connx.NewServer("127.0.0.1:7069", onConnHandler)
	if err != nil{
		fmt.Println("GetNewServer error", err)
		return
	}
	fmt.Println("GetNewServer begin listen")
	go func(){
		for{
			time.Sleep(time.Second * 10)
			fmt.Println("ConnCount:", server.GetConnectionCount())
		}
	}()

	server.Start()
}


func onConnHandler(conn *connx.Connection) error{
	msg, err := conn.ParseMessage()
	fmt.Println(msg, err)

	msgReply := connx.ResponseMessage("echo ok")
	conn.SendMessage(msgReply)
	return nil
}