package main

import (
	"fmt"
	"github.com/vignesh-j-shetty/GoDFS/internal/client"
)

func main() {
	conf, _ := client.LoadConfig()
	fmt.Printf(conf.MetadataServer)
}