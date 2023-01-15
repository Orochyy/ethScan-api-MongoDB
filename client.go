package main

import (
	"ethScan-api-MongoDB/accounts"
	"ethScan-api-MongoDB/blocks"
	"ethScan-api-MongoDB/contracts"
	"ethScan-api-MongoDB/gas"
	"ethScan-api-MongoDB/httpapi"
	"ethScan-api-MongoDB/logs"
	"ethScan-api-MongoDB/proxy"
	"ethScan-api-MongoDB/stats"
	"ethScan-api-MongoDB/tokens"
	"ethScan-api-MongoDB/transactions"
)

// Client is the main etherscan client.
type Client struct {
	Accounts     accounts.AccountsClient
	Contracts    contracts.ContractsClient
	Transactions transactions.TransactionsClient
	Blocks       blocks.BlocksClient
	Logs         logs.LogsClient
	Proxy        proxy.ProxyClient
	Tokens       tokens.TokensClient
	Gas          gas.GasClient
	Stats        stats.StatsClient
}

// Params are construction parameters for the etherscan Client.
type Params = httpapi.Params

// New constructs a new etherscan Client.
func New(params *Params) *Client {
	api := httpapi.New(params)
	return &Client{
		Accounts:     accounts.AccountsClient{API: api},
		Contracts:    contracts.ContractsClient{API: api},
		Transactions: transactions.TransactionsClient{API: api},
		Blocks:       blocks.BlocksClient{API: api},
		Logs:         logs.LogsClient{API: api},
		Proxy:        proxy.ProxyClient{API: api},
		Tokens:       tokens.TokensClient{API: api},
		Gas:          gas.GasClient{API: api},
		Stats:        stats.StatsClient{API: api},
	}
}
