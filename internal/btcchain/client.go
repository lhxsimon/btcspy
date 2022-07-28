package btcchain

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	conn        *rpcclient.Client
	chainParams *chaincfg.Params
}

func NewClient(host string, user string, password string) (*Client, error) {
	config := &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         password,
		DisableTLS:   true,
		HTTPPostMode: true,
	}
	btcClient, err := rpcclient.New(config, nil)
	if err != nil {
		return nil, err
	}
	log.Infof("try to connect to %s ...", host)
	err = btcClient.Ping()
	if err != nil {
		return nil, err
	}
	blockchainInfo, err := btcClient.GetBlockChainInfo()
	if err != nil {
		return nil, err
	}
	var chainParams *chaincfg.Params
	switch blockchainInfo.Chain {
	case "main":
		chainParams = &chaincfg.MainNetParams
	case "test":
		chainParams = &chaincfg.TestNet3Params
	default:
		return nil, fmt.Errorf("unsupport chain %s", blockchainInfo.Chain)
	}
	log.Infof("success connect to bitcoin %s net", blockchainInfo.Chain)
	return &Client{
		conn:        btcClient,
		chainParams: chainParams,
	}, nil
}
