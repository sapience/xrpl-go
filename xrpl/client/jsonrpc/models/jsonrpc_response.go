package jsonrpcmodels

import (
	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/mitchellh/mapstructure"
)

type JsonRpcResponse struct {
	Result    AnyJson                      `json:"result"`
	Warning   string                       `json:"warning,omitempty"`
	Warnings  []client.XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                         `json:"forwarded,omitempty"`
}

type AnyJson transactions.FlatTransaction

type ApiWarning struct {
	Id      int         `json:"id"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (r JsonRpcResponse) GetResult(v any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json",
		Result: &v, DecodeHook: mapstructure.TextUnmarshallerHookFunc()})

	if err != nil {
		return err
	}
	err = dec.Decode(r.Result)
	if err != nil {
		return err
	}
	return nil
}
