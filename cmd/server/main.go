package main

import (
	"apis/configs"
)

func main() {
	config := configs.LoadConfig()
	println(config.DBDriver)
}
