package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}

	fmt.Printf("\nport:8888\n")

	h := http.FileServer(http.Dir("E:/findproperty-android/findproperty/build/outputs/apk"))
	http.ListenAndServe(":8888", h)
}
