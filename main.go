package main

import (
	"github.com/sahirug/ethereum-dns/contractdeploy"
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

}
