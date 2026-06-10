package main

import (
	"fmt"

	"stream-service/services"
)

func main() {
	services.StartStream()
	fmt.Println("🚀 Quote Stream Worker running...")
}