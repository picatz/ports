package ports

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func connect(protocol, ip string, port int, timeout time.Duration) *Result {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout(protocol, target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			return connect(protocol, ip, port, timeout)
		}
		return &Result{
			Open:  false,
			Error: err,
			IP:    ip,
			Port:  port,
		}
	}

	conn.Close()

	return &Result{
		Open:  true,
		Error: err,
		IP:    ip,
		Port:  port,
	}
}

func Scan(protocol, ip string, ports ...int) <-chan *Result {
	results := make(chan *Result)
	go func() {
		defer close(results)
		wg := sync.WaitGroup{}
		for _, port := range ports {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				results <- connect(protocol, ip, port, 1000*time.Millisecond)
			}(port)
		}
		wg.Wait()
	}()
	return results
}
