package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type UdpLog struct {
	host string
	port int
	ServerName string
	chMsg chan string
	done chan struct{}
	conn *net.UDPConn
}
type LogMsg struct {
	Project string `json:"project"`
	Msg     string `json:"msg"`
}


func NewUdplog(config *udpConfig) (*UdpLog, error) {
    p := UdpLog{
    	host: config.host,
    	port: config.port,
    	ServerName: config.ServerName,
    	chMsg: make(chan string,10000),
    	done: make(chan struct{}),
	}
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%v",p.host,p.port))
	if err != nil {
		return nil,err
	}
	// Dial to the address with UDP
	p.conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil,err
	}
	for i := 0; i < 1; i++ {
		go p.run()
	}
	return &p,nil
}
func (p *UdpLog)Log(msg string){
	p.chMsg <- msg
}

func (p *UdpLog) run() {
loop:
	for{
		select {
		case <-p.done:
			break loop
		case m := <- p.chMsg:
			if p.conn != nil {
				msg := LogMsg{
					Project: p.ServerName,
					Msg:m,
				}
				marshal, err := json.Marshal(&msg)

				if err != nil {
					continue
				}
				p.conn.Write(marshal)
				bufio.NewReader(p.conn).ReadString('\n')
			}
		}
	}
}