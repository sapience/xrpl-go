package main

import (
	"fmt"
	"log"

	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
)

func main() {

	// init new config object with desired node address
	cfg, err := rpc.NewJsonRpcConfig("http://testnode/")
	if err != nil {
		log.Panicln(err)
	}

	// Initialise new json client with json config
	client := rpc.NewJsonRpcClient(cfg)

	// call the desired method
	var req *account.AccountChannelsRequest
	ac, err := client.GetAccountChannels(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Results mapped to struct: %v\n", ac)
}
