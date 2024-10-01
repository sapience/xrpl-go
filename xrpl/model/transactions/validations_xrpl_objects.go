package transactions

import (
	"errors"

	maputils "github.com/Peersyst/xrpl-go/pkg/map_utils"
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/utils"
)

const (
	MEMO_SIZE                  = 3
	SIGNER_SIZE                = 3
	ISSUED_CURRENCY_SIZE       = 3
	STANDARD_CURRENCY_CODE_LEN = 3
)

// IsMemo checks if the given object is a valid Memo object.
func IsMemo(obj map[string]interface{}) bool {
	// Check if the object is not nil and if it has a Memo field.
	if obj == nil || obj["Memo"] == nil {
		return false
	}

	// Check if the Memo field is a map.
	memo, isMap := obj["Memo"].(map[string]interface{})
	if !isMap {
		return false
	}

	// Get the size of the Memo object.
	size := len(maputils.GetKeys(memo))

	validData := memo["MemoData"] == nil || typecheck.IsHex(memo["MemoData"].(string))
	validFormat := memo["MemoFormat"] == nil || typecheck.IsHex(memo["MemoFormat"].(string))
	validType := memo["MemoType"] == nil || typecheck.IsHex(memo["MemoType"].(string))

	return size >= 1 && size <= MEMO_SIZE && validData && validFormat && validType && utils.ObjectOnlyHasFields(memo, []string{"MemoFormat", "MemoData", "MemoType"})
}

// IsSigner checks if the given object is a valid Signer object.
func IsSigner(obj map[string]interface{}) bool {
	signer, ok := obj["Signer"].(map[string]interface{})
	if !ok {
		return false
	}

	size := len(maputils.GetKeys(signer))
	validAccount := signer["Account"] != nil && typecheck.IsString(signer["Account"])
	validTxnSignature := signer["TxnSignature"] != nil && typecheck.IsString(signer["TxnSignature"])
	validSigningPubKey := signer["SigningPubKey"] != nil && typecheck.IsString(signer["SigningPubKey"])

	return size == SIGNER_SIZE && validAccount && validTxnSignature && validSigningPubKey

}

// IsAmount checks if the given object is a valid Amount object.
// It is a string for an XRP amount or a map for an IssuedCurrency amount.
func IsAmount(amount interface{}) bool {
	if typecheck.IsString(amount) {
		return true
	}

	amt, ok := amount.(map[string]interface{})
	if !ok {
		return false
	}

	if IsIssuedCurrency(amt) {
		return true
	}

	return false
}

// IsIssuedCurrency checks if the given object is a valid IssuedCurrency object.
func IsIssuedCurrency(input interface{}) bool {
	// Type the input as a map.
	i, ok := input.(map[string]interface{})
	if !ok {
		return false
	}

	return len(maputils.GetKeys(i)) == ISSUED_CURRENCY_SIZE &&
		typecheck.IsString(i["value"]) &&
		typecheck.IsString(i["issuer"]) &&
		typecheck.IsString(i["currency"])
}

// IsPathStep checks if the given map is a valid PathStep.
func IsPathStep(pathStep map[string]interface{}) bool {
	if account, ok := pathStep["account"]; ok && !typecheck.IsString(account) {
		return false
	}
	if currency, ok := pathStep["currency"]; ok && !typecheck.IsString(currency) {
		return false
	}
	if issuer, ok := pathStep["issuer"]; ok && !typecheck.IsString(issuer) {
		return false
	}
	if _, ok := pathStep["account"]; ok {
		if _, ok := pathStep["currency"]; !ok {
			if _, ok := pathStep["issuer"]; !ok {
				return true
			}
		}
	}

	// check if the path step has either a currency or an issuer
	_, hasCurr := pathStep["currency"]
	_, hasIssuer := pathStep["issuer"]

	if !hasCurr && !hasIssuer {
		return true
	}
	return false
}

// IsPath checks if the given slice of maps is a valid Path.
func IsPath(path []map[string]interface{}) bool {
	for _, pathStep := range path {
		if !IsPathStep(pathStep) {
			return false
		}
	}
	return true
}

// IsPaths checks if the given slice of slices of maps is a valid Paths.
func IsPaths(paths [][]map[string]interface{}) bool {
	if len(paths) == 0 {
		return false
	}

	for _, path := range paths {
		if len(path) == 0 {
			return false
		}

		if !IsPath(path) {
			return false
		}
	}

	return true
}

// CheckIssuedCurrencyIsNotXrp checks if the given transaction map does not have an issued currenc as XRP.
func CheckIssuedCurrencyIsNotXrp(tx map[string]interface{}) error {
	keys := maputils.GetKeys(tx)
	for _, value := range keys {
		result, isFlatTxn := (tx[value]).(map[string]interface{})

		// Check if the value is an issued currency
		if isFlatTxn && IsIssuedCurrency(result) {
			// Check if the issued currency is XRP (which is incorrect)
			currency := tx[value].(map[string]interface{})["currency"].(string)

			if len(currency) == STANDARD_CURRENCY_CODE_LEN && currency == "XRP" {
				return errors.New("cannot have an issued currency with a similar standard code as XRP")
			}
		}
	}

	return nil
}
