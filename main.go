package main

import (
	"fmt"
	"golang-be/routers"
	"golang-be/utils/helper"
)

func main() {
	listenAddress := helper.GetEnv("SERVER_PORT", ":40001")
	fmt.Println("Starting listen address: ", listenAddress)
	routers.Server(listenAddress)
}
