package main

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
	log.SetPrefix("Blockchain: ")
}

func main() {
	var port int
	var err error
	portInStr := os.Getenv("BLOCKCHAIN_SERVER_PORT")
	if portInStr == ""{
		port = 5000
	}else{
		port, err = strconv.Atoi(portInStr)
		if err != nil {
			log.Fatalf("Invalid Blockchain Port set as environment variable: Needs to be in digits")
		}
	}
	app := newBlockchainServer(uint16(port))
	app.Run()
}
