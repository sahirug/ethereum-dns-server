package contractload

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sahirug/ethereum-dns/contracts"
	"github.com/sahirug/ethereum-dns/errorhandler"
)

func LoadContract(hexAddr string) *eddns.Eddns {
	client, err := ethclient.Dial("http://localhost:7545") // ganache
	//client, err := ethclient.Dial("http://localhost:8545") // docker
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