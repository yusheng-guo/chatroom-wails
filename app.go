package main

import (
	"context"
	"log"
	"net"
	"net/url"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// App struct
type App struct {
	ctx  context.Context
	conn net.Conn
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	// 连接服务器
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/",
	}
	conn, _, _, err := ws.Dial(ctx, u.String())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("----------------Connected----------------")
	a.conn = conn
	a.ctx = ctx
}

// shutdown 释放内存和执行关闭任务
func (a *App) shutdown(ctx context.Context) {
	// 关闭连接
	if a.conn != nil {
		a.conn.Close()
		a.conn = nil
		log.Println("----------------Closed----------------")
	}
}

// 接收消息
func (a *App) ReceiveMessage() string {
	if a.conn == nil {
		log.Println("Conn Failure")
		return ""
	}
	msg, _, err := wsutil.ReadServerData(a.conn)
	if err != nil {
		log.Println("Read error:", err)
	}
	log.Println("----------------Receive Successful!----------------")
	log.Println("----------------", string(msg), "----------------")
	return string(msg)
}

// 发送消息
func (a *App) SendMessage(msg string) {
	if a.conn != nil {
		err := wsutil.WriteClientText(a.conn, []byte(msg))
		if err != nil {
			log.Println("Send error:", err)
		}
		log.Println("----------------Send Successful!----------------")
	}
}
