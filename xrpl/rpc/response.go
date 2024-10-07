package rpc

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/mitchellh/mapstructure"
)

type JsonRpcResponse struct {
	Result    AnyJson               `json:"result"`
	Warning   string                `json:"warning,omitempty"`
	Warnings  []XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                  `json:"forwarded,omitempty"`
}

type XRPLResponseWarning struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

type AnyJson transaction.FlatTransaction

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

type JsonRpcXRPLResponse interface {
	GetResult(v any) error
}
