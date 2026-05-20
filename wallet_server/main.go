package main

import (
	"log"
	"os"
	"fmt"
	"strconv"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
	log.SetPrefix("Wallet Server: ")
}

func main() {
	
	var port int
	var err error

	portInStr := os.Getenv("WALLET_SERVER_PORT")
	if portInStr == ""{
		port = 8080
	}else{
		port, err = strconv.Atoi(portInStr)
		if err != nil {
			log.Fatalf("Invalid Wallet Port set as environment variable: Needs to be in digits")
		}
	}

	blc_server_host := os.Getenv("BLOCKCHAIN_SERVER_HOST")
	if blc_server_host == "" {
		blc_server_host = "http://127.0.0.1"
	}
	blc_server_port := os.Getenv("BLOCKCHAIN_SERVER_PORT")
	if blc_server_port == "" {
		blc_server_host = "5000"
	}
	
	gateway := fmt.Sprintf("%s:%s", blc_server_host, blc_server_port)

	app := newWalletServer(uint16(port), gateway)
	app.Run()
}
