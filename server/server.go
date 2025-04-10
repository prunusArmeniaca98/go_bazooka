package server

import (
	"fmt"
	"net"
	"sync"
)

type AcceptError func(err error)

// OpenSuccess 打开端口成功的回调
type OpenSuccess func()

// ClientConnected 有新的客户机连接进入
type ClientConnected func(ip string, port int, clientAddr string)

func NewServer(port int) *BazookaServer {
	return &BazookaServer{
		port: port,
	}
}

// BazookaServer 服务端
type BazookaServer struct {
	port               int // 端口号
	listener           net.Listener
	acceptErrorHandler AcceptError //监听发生错误的回调
	//todo 还没用到
	clientMaxSize uint16 //最大允许客户端连接数量
	//todo 还没用到
	clientConnectedHandler ClientConnected //有新的客户机连接进入
	//todo 还没用到
	rwLock sync.RWMutex //读写锁
}

// BindAcceptErrorHandler 绑定监听发生错误的回调
func (b *BazookaServer) BindAcceptErrorHandler(acceptErrorHandler AcceptError) {
	b.acceptErrorHandler = acceptErrorHandler
}

// BindClientMaxSize 绑定最大允许客户端连接数量
func (b *BazookaServer) BindClientMaxSize(clientMaxSize uint16) {
	b.clientMaxSize = clientMaxSize
}

func (b *BazookaServer) BindClientConnected(clientConnected ClientConnected) {
	b.clientConnectedHandler = clientConnected
}

// Open 打开服务端
func (b *BazookaServer) Open() error {
	var err error
	b.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", b.port))
	if err != nil {
		return err
	}
	defer b.accept()
	return nil
}

// 监听
func (b *BazookaServer) accept() {
	for {
		conn, err := b.listener.Accept() // 接受客户端连接
		if err != nil {
			if b.acceptErrorHandler != nil {
				go b.acceptErrorHandler(err)
			}
			continue
		}
		go b.clientHandle(conn)
	}
}

// Close 关闭端口
func (b *BazookaServer) Close() error {
	return b.listener.Close()
}

func (b *BazookaServer) clientHandle(conn net.Conn) {

}
