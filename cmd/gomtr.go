package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gooops/mtr"
)

func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ".")
}

func main() {
	fmt.Println(mtr.DEFAULT_RETRIES)
	mtr.LocalAddr()
	destAddrs, _ := mtr.DestAddr("www.baidu.com")

	fmt.Println("ip:", convert(destAddrs[:]))

	c := make(chan mtr.TracerouteHop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			fmt.Println(hop.TTL, hop.Address, hop.AvgTime, hop.BestTime, hop.Loss)
		}
	}()
	options := mtr.TracerouteOptions{}
	_, err := mtr.Mtr(destAddrs, &options, c)
	if err != nil {
		fmt.Println(err)
	}

	mm, err := mtr.T("www.baidu.com", true, 0, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mm)

	// info, err := mtr.T("www.baidu.com", false, 0, 0, 0, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(info)
}
