package mtr

import (
	"fmt"
	"syscall"
)

func Udp(socketAddr, destAddr [4]byte, ttl, port int, tv syscall.Timeval, p []byte) (hop TracerouteReturn, err error) {
	fmt.Println("socketAddr: ", socketAddr, " destAddr:  ", destAddr, " ttl:  ", ttl, " port:", port, " tv: ", tv, " p: ", p)
	// start := time.Now()

	return hop, err
}
