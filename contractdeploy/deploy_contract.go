package contractdeploy

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sahirug/ethereum-dns/contracts"
	"github.com/sahirug/ethereum-dns/errorhandler"
	"math/big"
)

func DeployContract() (string, common.Hash) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		errorhandler.HandleErr(err, 1)
	}

	//privateKey, err := crypto.HexToECDSA("bc5b578e0dcb2dbf98dd6e5fe62cb5a28b84a55e15fc112d4ca88e1f62bd7c35") // docker
	privateKey, err := crypto.HexToECDSA("a8ecc1d94b93080a347a994a35b8598ebe73eea21073739af847b7580230539a") //truffle
	if err != nil {
		errorhandler.HandleErr(err, 2)
	}

	publicKey := privateKey.Public()

	publicKeyEcdsa, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		errorhandler.HandleErr(errors.New("pub key not of type ecdsa.PublicKey"), 3)
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyEcdsa)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)

	if err != nil {
		errorhandler.HandleErr(err, 4)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		errorhandler.HandleErr(err, 5)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(4700000)
	auth.GasPrice = gasPrice

	address, tx, instance, err := eddns.DeployEddns(auth, client)

	if err != nil {
		errorhandler.HandleErr(err, 6)
	}

	//fmt.Println(address.Hex())
	//fmt.Println(tx.Hash().Hex())

	_ = instance
	return address.Hex(), tx.Hash()
}

//func handleError(err error, codeblock uint) {
//	log.Printf("Error at codeblock %d", codeblock)
//	log.Fatalf("Error occurred while deploying contract: %s", err)
//}
