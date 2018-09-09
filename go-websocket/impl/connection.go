package impl

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn *websocket.Conn
	// 读协程
	inChan chan []byte
	// 写协程
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
		isClosed:  false,
	}
	//  启动读协程
	go conn.readLoop()
	go conn.writeLoop()
	return
	// fmt.Println("vim-go")
}

// API
func (conn *Connection) ReadMessage(data []byt, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}

	return
}

func (conn *Connection) WriteMessage(data []byt, err error) {
	// data = <-conn.inChan
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) Close() {
	// 线程安全的close
	// 可多次调用，可重入的
	conn.wsConn.Close()
	// 这一行只能执行一次
	conn.mutex.Lock()
	if !conn.isClosed {

		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		// 阻塞在这里，等待inChan有空闲的位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			//closeChan关闭
			goto ERR

		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}
