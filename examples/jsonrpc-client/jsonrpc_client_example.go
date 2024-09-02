package main

import (
	"fmt"
	"log"

	"github.com/Peersyst/xrpl-go/xrpl/client"
	jsonrpcclient "github.com/Peersyst/xrpl-go/xrpl/client/jsonrpc"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/account"
)

func main() {

	// init new config object with desired node address
	cfg, err := client.NewJsonRpcConfig("http://testnode/")
	if err != nil {
		log.Panicln(err)
	}

	// Initialise new json client with json config
	client := jsonrpcclient.NewClient(cfg)

	// call the desired method
	var req *account.AccountChannelsRequest
	ac, xrplRes, err := client.Account.GetAccountChannels(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Full XRPL response: %v\n", xrplRes)
	fmt.Printf("Results mapped to struct: %v\n", ac)
}
