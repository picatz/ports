package ports

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	ip := "192.168.33.10"

	for result := range Scan("tcp", ip, 22, 2222) {
		fmt.Println(result)
	}
}
