package main

import (
	"ethScan-api-MongoDB/handler"
	"github.com/gorilla/mux"
	"github.com/nanmu42/etherscan-api"
	"log"
	"net/http"
)

func main() {

	client := etherscan.New(etherscan.Mainnet, "9PE85ZN1F87E9YMDP9NBF1MXWU4EA1QSYP")

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/eth/{module}", handler.EthScanHandler{Client: client}.ServeHTTP)

	log.Fatal(http.ListenAndServe(":8080", r))

}
