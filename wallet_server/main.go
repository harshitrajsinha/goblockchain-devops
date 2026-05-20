package main

import (
	"flag"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
	log.SetPrefix("Wallet Server: ")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Wallet Server")
	blc_server_host := os.Getenv("BLOCKCHAIN_SERVER_HOST")
	blc_server_port := os.Getenv("BLOCKCHAIN_SERVER_PORT")
	
	gateway := flag.String("gateway", fmt.Sprintf("%s:%s", blcServerHost, blcServerPort), "Blockchain Gateway")
	flag.Parse()

	app := NewWalletServer(uint16(*port), *gateway)
	app.Run()
}
