package main

import (
	"github.com/sahirug/ethereum-dns/contractdeploy"
	"github.com/sahirug/ethereum-dns/contractload"
	"github.com/sahirug/ethereum-dns/dns"
	"log"
)

func main() {
	log.Println("============================================================== STARTING ETHEREUM DNS ==============================================================")
	log.Println("================================================= Author: Sahiru Gunawardene <sahirug@gmail.com> ==================================================")
	//log.Println("=========================================================== Status: Deploying Contract ============================================================")
	log.Println("Deploying contract....")
	//err := errors.New("fuck")
	//log.Fatalf("Error occurred while deploying contract: %s", err)
	hexAddr, txHash := contractdeploy.DeployContract()
	log.Printf("Contract deployed at %s", hexAddr)
	log.Printf("Transaction hash hex: %s", txHash.Hex())

	log.Println("Loading contract....")

	instance := contractload.LoadContract(hexAddr)

	log.Println("Contract loaded!")

	dns.Start(instance)

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
