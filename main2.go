package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
)

func main() {
	ip, err := hexutil.Decode("0x312e322e332e340000000000000000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ip))
	// 0x312e322e332e340000000000000000
}
