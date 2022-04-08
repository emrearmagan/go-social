/*
main.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package main

import (
	"go-social/config"
	"log"
)

func main() {
	//load
	if err := config.LoadConfig("./config/config_example.json"); err != nil {
		log.Fatalf("Failed to load config file: %s", err)
	}
}
