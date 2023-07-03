// https://www.elephdev.com/golang/285.html?ref=addtabs&lang=en
// Modified: prr, azul software
// Date: 3/7/2023
// copyright 2023 prr, azul software
//

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
//	"os"
//	"strconv"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {

	defIpStr := "89.116.30.49"
	portStr := "10600"

	wsStr := "ws://" + defIpStr + ":" + portStr +"/" 
	log.Printf("Client dialing %s\n", wsStr)

	for icount:=1; icount<5;icount++ {
		conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), wsStr)
		if err != nil {
			log.Printf("Cannot connect: %v\n", err)
			log.Printf("sleeping for 5 seconds\n")
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		log.Printf("Connection %d established to server at %s\n", icount, wsStr)

		for i := 0; i <10; i++ {
//			randomNumber := strconv.Itoa(rand.Intn(100))
			msgStr := fmt.Sprintf("msg[%d]: hello %d\n", i+1, rand.Intn(100))
			msg := []byte(msgStr)

			err = wsutil.WriteClientMessage(conn, ws.OpText, msg)
			if err != nil {
				fmt.Println("Cannot send: "+ err.Error())
				continue
			}
			log.Printf("Client send msg[%d]: %s\n", i+1, msgStr)
			rcvMsg, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				log.Printf("Cannot receive data: %v\n", err)
				continue
			}
			log.Printf("from server: %s\n", string(rcvMsg))
			time.Sleep(time.Duration(5) * time.Second)
		}
		err = conn.Close()
		if err != nil {
			log.Fatalf("Cannot close the connection: %v", err)
		}
		log.Println("Disconnected from server")
	}

}
