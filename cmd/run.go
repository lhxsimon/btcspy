package main

import (
	"btcspy/internal/btcchain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

var rpcHost = "127.0.0.1:8332"
var rpcUser = "bitcoin"
var rpcPassword = "bitcoin"

func main() {
	// prepare
	prepare()
	err := runCmd()
	if err != nil {
		log.Fatal(err)
	}
}

func runCmd() error {
	// 设置命令行
	var rootCmd = &cobra.Command{
		Use:   "btcspy",
		Short: "BTC 工具箱",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	var balanceCmd = &cobra.Command{
		Use:     "balance",
		Short:   "查询余额",
		Example: "btc balance <btc_address>",
		Args:    cobra.ExactArgs(1),
		Run:     cmdBalance,
	}

	var transactionCmd = &cobra.Command{
		Use:     "transaction",
		Short:   "查询交易",
		Example: "btc transaction <btc_transaction>",
		Args:    cobra.ExactArgs(1),
		Run:     cmdTransaction,
	}

	var blockCmd = &cobra.Command{
		Use:     "block",
		Short:   "查询区块",
		Example: "btc block <btc_height>",
		Args:    cobra.MaximumNArgs(1),
		Run:     cmdBlock,
	}

	rootCmd.PersistentFlags().StringP("host", "H", rpcHost, "RPC Host")
	rootCmd.PersistentFlags().StringP("user", "u", rpcUser, "RPC User")
	rootCmd.PersistentFlags().StringP("password", "p", rpcPassword, "RPC Password")
	rootCmd.AddCommand(balanceCmd)
	rootCmd.AddCommand(transactionCmd)
	rootCmd.AddCommand(blockCmd)
	return rootCmd.Execute()
}

func cmdBalance(cmd *cobra.Command, args []string) {
	client, err := btcchain.NewClient(rpcHost, rpcUser, rpcPassword)
	if err != nil {
		log.Fatal(err)
	}
	err = client.GetBalance(args[0])
	if err != nil {
		log.Fatal(err)
	}
}

func cmdTransaction(cmd *cobra.Command, args []string) {
	client, err := btcchain.NewClient(rpcHost, rpcUser, rpcPassword)
	if err != nil {
		log.Fatal(err)
	}
	err = client.GetTransaction(args[0])
	if err != nil {
		log.Fatal(err)
	}
}

func cmdBlock(cmd *cobra.Command, args []string) {
	var blockHeight int64
	var err error
	client, err := btcchain.NewClient(rpcHost, rpcUser, rpcPassword)
	if err != nil {
		log.Fatal(err)
	}
	if len(args) == 0 {
		blockHeight = 0
	} else {
		blockHeight, err = strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = client.GetBlock(blockHeight)
	if err != nil {
		log.Fatal(err)
	}
}
