package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/privval"

	"github.com/jdkanani/web3-example/contracts/rootchain"
)

func main() {
	client, err := ethclient.Dial("https://kovan.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// with no 0x
	rootchainAddress := "24e01716a6ac34d5f2c4c082f553d86a557543a7"

	// token address
	tokenAddress := "670568761764f53E6C10cd63b71024c31551c9EC"

	// with no 0x
	// priv := "117bbcf6bdc3a8e57f311a2b4f513c25b20e3ad4606486d7a927d8074872c2af"

	rootchainClient, err := rootchain.NewRootchain(common.HexToAddress(rootchainAddress), client)

	// private key
	// priv := "117bbcf6bdc3a8e57f311a2b4f513c25b20e3ad4606486d7a927d8074872c2af"

	privVal := privval.LoadFilePV(config.DefaultBaseConfig().PrivValidatorFile())
	ecdsaPk, _ := crypto.ToECDSA(privVal.PrivKey.Bytes()[:])
	auth := bind.NewKeyedTransactor(ecdsaPk)

	/**
	 * Calling contract method
	 */
	var amount big.Int
	amount.SetUint64(0)
	tx, err := rootchainClient.Deposit(auth, common.HexToAddress(tokenAddress), common.BytesToAddress(privVal.Address.Bytes()), &amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

}
