package btcchain

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (c *Client) GetBalance(account string) error {
	amount, err := c.conn.GetBalance(account)
	if err != nil {
		return err
	}
	log.Info("balance %s amount: %s BTC", account, amount.String())
	return nil
}

func (c *Client) GetTransaction(rawHash string) error {
	txHash, err := chainhash.NewHashFromStr(rawHash)
	if err != nil {
		return err
	}
	txData, err := c.conn.GetRawTransaction(txHash)
	if err != nil {
		return err
	}
	log.Info(strings.Repeat("-", 30))
	for _, el := range txData.MsgTx().TxOut {

		pkScript, err := txscript.ParsePkScript(el.PkScript)
		if err != nil {
			continue
		}
		address, err := pkScript.Address(c.chainParams)
		if err != nil {
			log.Error(err)
		}
		log.Infof("%s -> %f BTC", address, btcutil.Amount(el.Value).ToBTC())
	}
	log.Info(strings.Repeat("-", 30))
	return nil
}

func (c *Client) GetBlock(blockHeight int64) error {
	if blockHeight == 0 {
		var err error
		blockHeight, err = c.conn.GetBlockCount()
		if err != nil {
			return err
		}
	}
	blockHash, err := c.conn.GetBlockHash(blockHeight)
	if err != nil {
		return err
	}
	blockData, err := c.conn.GetBlock(blockHash)
	if err != nil {
		return err
	}
	log.Info(strings.Repeat("-", 30))
	log.Infof("block height %d has %d transactions", blockHeight, len(blockData.Transactions))
	for _, el := range blockData.Transactions {
		log.Infof("transaction: %s", el.TxHash().String())
	}
	log.Info(strings.Repeat("-", 30))
	return nil
}
