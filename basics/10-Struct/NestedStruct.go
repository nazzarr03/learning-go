package main

import "fmt"

type Configuration struct {
	Env   string
	Proxy Proxy
}

type Proxy struct {
	Address string
	Port    string
}

func main() {
	c := &Configuration{
		Env: "DEBUG:TRUE",
		Proxy: Proxy{
			Address: "addr",
			Port:    "port",
		},
	}

	fmt.Println(c)
	fmt.Println(c.Proxy.Address)
}