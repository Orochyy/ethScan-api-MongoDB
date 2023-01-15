package handler

import (
	"context"
	"ethScan-api-MongoDB/proxy"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/nanmu42/etherscan-api"
	"net/http"
)

type EthScanHandler struct {
	*etherscan.Client
	p proxy.ProxyClient
}

func (e EthScanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	module := vars["module"]

	w.Header().Set("Content-Type", "application/json")

	switch module {
	case "hs":

		txHash := common.HexToHash("0x1e2910a262b1008d0616a0beb24c1a491d78771baa54a33e66065e03b1f46bc1")

		fmt.Println(txHash)

		hs, _ := e.p.GetTransactionByHash(context.Background(), txHash)

		fmt.Println(hs)

	}
}
