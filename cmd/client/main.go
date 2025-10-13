package main

import (
	"fmt"

	"github.com/vignesh-j-shetty/GoDFS/internal/config"
)

func main() {
	conf, _ := config.LoadConfig()
	fmt.Printf(conf.MetadataServer)
}