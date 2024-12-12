package xrpl

import (
	"errors"
	"fmt"
	"sort"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
)

var (
	ErrNoTxToMultisign = errors.New("no transaction to multisign")
)

func Multisign(blobs ...string) (string, error) {
	if len(blobs) == 0 {
		return "", ErrNoTxToMultisign
	}

	signers := make([]interface{}, 0)
	for _, blob := range blobs {
		tx, err := binarycodec.Decode(blob)
		if err != nil {
			return "", err
		}

		signers = append(signers, tx["Signers"].([]interface{})...)
	}

	tx, err := binarycodec.Decode(blobs[0])
	if err != nil {
		return "", err
	}

	tx["Signers"] = sortSigners(signers)

	blob, err := binarycodec.Encode(tx)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}

	return blob, nil
}

func sortSigners(signers []interface{}) []interface{} {
	sort.Slice(signers, func(i, j int) bool {
		iSigner := signers[i].(map[string]interface{})["Signer"].(map[string]interface{})
		jSigner := signers[j].(map[string]interface{})["Signer"].(map[string]interface{})

		return iSigner["Account"].(string) > jSigner["Account"].(string)
	})
	return signers
}
