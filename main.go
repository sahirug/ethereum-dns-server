package main

import (
	"flag"
	"github.com/sahirug/ethereum-dns/contractdeploy"
	"github.com/sahirug/ethereum-dns/contractload"
	"github.com/sahirug/ethereum-dns/dns"
	"log"
)

func main() {

	// flags
	networkType := flag.String("network", "private", "the network type")

	flag.Parse()

	log.Println("============================================================== STARTING ETHEREUM DNS ==============================================================")
	log.Println("================================================= Author: Sahiru Gunawardene <sahirug@gmail.com> ==================================================")
	if *networkType == "ganache" {
		log.Println("Using Ganache Test Network")
	} else if *networkType == "private" {
		log.Println("Using private Ethereum network")
	} else {
		log.Fatal("Unknown network type. Exiting...")
	}
	log.Println("Deploying contract....")
	hexAddr, txHash, auth := contractdeploy.DeployContract(*networkType)
	log.Printf("Contract deployed at %s", hexAddr)
	log.Printf("Transaction hash hex: %s", txHash.Hex())

	log.Println("Loading contract....")

	instance := contractload.LoadContract(hexAddr, *networkType)

	log.Println("Contract loaded!")

	dns.Start(instance, auth)

	//var tld [12]byte
	//
	//copy(tld[:], []byte("com"))
	//
	//ip, err := instance.GetIp(nil, []byte("google"), tld)
	//
	//if err != nil {
	//	errorhandler.HandleErr(err, 9)
	//}
	//
	//log.Printf("IP: %d", len(ip))
	//log.Printf("IP: %v", ip)

}
