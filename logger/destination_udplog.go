package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type destinationUdpLog struct {
	chMsg chan string
	buf   bytes.Buffer
	done  chan struct{}
	conn  *net.UDPConn

	server_address string
	server_name    string
}
type LogMsg struct {
	Project string `json:"project"`
	Msg     string `json:"msg"`
}

func newDestinationUdpLog(server_address string, serverName string) (destination, error) {

	udpAddr, err := net.ResolveUDPAddr("udp", server_address)
	if err != nil {
		return nil, err
	}
	_con, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, err
	}
	_con.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	p := &destinationUdpLog{
		server_address: server_address,
		server_name:    serverName,
		chMsg:          make(chan string, 100_000),
		done:           make(chan struct{}),
		conn:           _con,
	}

	go p.run()
	return p, nil
}

func (d *destinationUdpLog) log(t time.Time, level Level, format string, args ...interface{}) {
	d.buf.Reset()
	writeContent(&d.buf, format, args)
	d.chMsg <- d.buf.String()
}

func (d *destinationUdpLog) close() {
	close(d.done)
}

func (p *destinationUdpLog) run() {
loop:
	for {
		select {
		case <-p.done:
			break loop
		case m := <-p.chMsg:
			if p.conn != nil {
				msg := LogMsg{
					Project: p.server_name,
					Msg:     m,
				}
				marshal, err := json.Marshal(&msg)
				if err != nil {
					fmt.Println("%v", err.Error())
					continue
				}
				_, err = p.conn.Write(marshal)
				if err != nil {
					fmt.Println("Reconect udp server %v", err.Error())
					udpAddr, err := net.ResolveUDPAddr("udp", p.server_address)
					if err != nil {
						fmt.Println("%v", err.Error())
						continue

					}
					_con, err := net.DialUDP("udp", nil, udpAddr)
					if err != nil {
						fmt.Println("%v", err.Error())
						continue
					}
					_con.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
					p.conn = _con
					continue
				}
			}
		}
	}
}
