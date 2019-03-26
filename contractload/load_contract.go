package contractload

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sahirug/ethereum-dns/contracts"
	"github.com/sahirug/ethereum-dns/errorhandler"
)

func LoadContract(hexAddr string, networkType string) *eddns.Eddns {

	var rawurl string

	if networkType == "ganache" {
		rawurl = "http://localhost:8545"
	} else {
		rawurl = "http://localhost:8545"
	}

	client, err := ethclient.Dial(rawurl)

	if err != nil {
		errorhandler.HandleErr(err, 7)
	}
	address := common.HexToAddress(hexAddr)
	instance, err := eddns.NewEddns(address, client)

	if err != nil {
		errorhandler.HandleErr(err, 8)
	}

	return instance
}