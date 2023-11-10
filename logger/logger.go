// Package logger contains a logger implementation.
package logger

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/gookit/color"
)

// Logger is a log handler.
type Logger struct {
	level Level

	destinations []destination
	mutex        sync.Mutex
}

// New allocates a log handler.
func New(level Level, destinations []Destination, filePath string) (*Logger, error) {
	lh := &Logger{
		level: level,
	}

	for _, destType := range destinations {
		switch destType {
		case DestinationStdout:
			lh.destinations = append(lh.destinations, newDestionationStdout())

		case DestinationFile:
			dest, err := newDestinationFile(filePath)
			if err != nil {
				lh.Close()
				return nil, err
			}
			lh.destinations = append(lh.destinations, dest)

		case DestinationSyslog:
			dest, err := newDestinationSyslog()
			if err != nil {
				lh.Close()
				return nil, err
			}
			lh.destinations = append(lh.destinations, dest)
		}
	}

	return lh, nil
}

// New allocates a log handler.
func New2(level Level, destinations []Destination, filePath string, udp_server string, server_name string) (*Logger, error) {
	lh := &Logger{
		level: level,
	}

	for _, destType := range destinations {
		switch destType {
		case DestinationStdout:
			lh.destinations = append(lh.destinations, newDestionationStdout())

		case DestinationFile:
			dest, err := newDestinationFile(filePath)
			if err != nil {
				lh.Close()
				return nil, err
			}
			lh.destinations = append(lh.destinations, dest)

		case DestinationSyslog:
			dest, err := newDestinationSyslog()
			if err != nil {
				lh.Close()
				return nil, err
			}
			lh.destinations = append(lh.destinations, dest)
		case DestinationUdplog:
			dest, err := newDestinationUdpLog(udp_server, server_name)
			if err != nil {
				lh.Close()
				return nil, err
			}
			lh.destinations = append(lh.destinations, dest)
		}
	}

	return lh, nil
}

func (p *Logger) EnableUDPLogServer(sv string, server_name string) bool {
	dest, err := newDestinationUdpLog(sv, server_name)
	if err != nil {
		return false
	}
	p.destinations = append(p.destinations, dest)
	return true
}

func removea[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func (p *Logger) DisableUDPLogServer() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for i, dest := range p.destinations {
		if dest.Type() == 3 {
			dest.close()
			p.destinations = removea(p.destinations, i)
			break
		}
	}
}

// Close closes a log handler.
func (lh *Logger) Close() {
	for _, dest := range lh.destinations {
		dest.close()
	}
}

// https://golang.org/src/log/log.go#L78
func itoa(i int, wid int) []byte {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	return b[bp:]
}

func writeTime(buf *bytes.Buffer, t time.Time, useColor bool) {
	var intbuf bytes.Buffer

	// date
	year, month, day := t.Date()
	intbuf.Write(itoa(year, 4))
	intbuf.WriteByte('/')
	intbuf.Write(itoa(int(month), 2))
	intbuf.WriteByte('/')
	intbuf.Write(itoa(day, 2))
	intbuf.WriteByte(' ')

	// time
	hour, min, sec := t.Clock()
	intbuf.Write(itoa(hour, 2))
	intbuf.WriteByte(':')
	intbuf.Write(itoa(min, 2))
	intbuf.WriteByte(':')
	intbuf.Write(itoa(sec, 2))
	intbuf.WriteByte(' ')

	if useColor {
		buf.WriteString(color.RenderString(color.Gray.Code(), intbuf.String()))
	} else {
		buf.WriteString(intbuf.String())
	}
}

func writeLevel(buf *bytes.Buffer, level Level, useColor bool) {
	switch level {
	case Debug:
		if useColor {
			buf.WriteString(color.RenderString(color.Debug.Code(), "DEB"))
		} else {
			buf.WriteString("DEB")
		}

	case Info:
		if useColor {
			buf.WriteString(color.RenderString(color.Green.Code(), "INF"))
		} else {
			buf.WriteString("INF")
		}

	case Warn:
		if useColor {
			buf.WriteString(color.RenderString(color.Warn.Code(), "WAR"))
		} else {
			buf.WriteString("WAR")
		}

	case Error:
		if useColor {
			buf.WriteString(color.RenderString(color.Error.Code(), "ERR"))
		} else {
			buf.WriteString("ERR")
		}
	}
	buf.WriteByte(' ')
}

func writeContent(buf *bytes.Buffer, format string, args []interface{}) {
	buf.Write([]byte(fmt.Sprintf(format, args...)))
	buf.WriteByte('\n')
}

// Log writes a log entry.
func (lh *Logger) Log(level Level, format string, args ...interface{}) {
	if level < lh.level {
		return
	}

	lh.mutex.Lock()
	defer lh.mutex.Unlock()

	t := time.Now()

	for _, dest := range lh.destinations {
		dest.log(t, level, format, args...)
	}
}
