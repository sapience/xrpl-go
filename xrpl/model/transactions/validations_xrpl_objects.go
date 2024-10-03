package transactions

import (
	"errors"

	maputils "github.com/Peersyst/xrpl-go/pkg/map_utils"
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
)

const (
	// The Memos field includes arbitrary messaging data with the transaction.
	// It is presented as an array of objects. Each object has only one field, Memo,
	// which in turn contains another object with one or more of the following fields:
	// MemoData, MemoFormat, and MemoType. https://xrpl.org/docs/references/protocol/transactions/common-fields#memos-field
	MEMO_SIZE                  = 3
	SIGNER_SIZE                = 3
	ISSUED_CURRENCY_SIZE       = 3
	STANDARD_CURRENCY_CODE_LEN = 3
)

// IsMemo checks if the given object is a valid Memo object.
func IsMemo(memo Memo) (bool, error) {
	// Get the size of the Memo object.
	size := len(maputils.GetKeys(memo.Flatten()))

	if size == 0 {
		return false, errors.New("memo object should have at least one field, MemoData, MemoFormat or MemoType")
	}

	if size > MEMO_SIZE {
		return false, errors.New("memo object should have at most three fields, MemoData, MemoFormat and MemoType")
	}

	validData := memo.MemoData == "" || typecheck.IsHex(memo.MemoData)
	if !validData {
		return false, errors.New("memoData should be a hexadecimal string")
	}

	validFormat := memo.MemoFormat == "" || typecheck.IsHex(memo.MemoFormat)
	if !validFormat {
		return false, errors.New("memoFormat should be a hexadecimal string")
	}

	validType := memo.MemoType == "" || typecheck.IsHex(memo.MemoType)
	if !validType {
		return false, errors.New("memoType should be a hexadecimal string")
	}

	return true, nil
}

// IsSigner checks if the given object is a valid Signer object.
func IsSigner(signerData SignerData) (bool, error) {
	size := len(maputils.GetKeys(signerData.Flatten()))
	if size != SIGNER_SIZE {
		return false, errors.New("signers: Signer should have 3 fields: Account, TxnSignature, SigningPubKey")
	}

	// TODO: Update to check if the account is valid when the "isAccount" function exists
	validAccount := signerData.Account != "" && typecheck.IsString(signerData.Account.String())
	if !validAccount {
		return false, errors.New("signers: Account should be a string")
	}

	validTxnSignature := signerData.TxnSignature != "" && typecheck.IsString(signerData.TxnSignature)
	if !validTxnSignature {
		return false, errors.New("signers: TxnSignature should be a string")
	}

	validSigningPubKey := signerData.SigningPubKey != "" && typecheck.IsString(signerData.SigningPubKey)
	if !validSigningPubKey {
		return false, errors.New("signers: SigningPubKey should be a string")
	}

	return true, nil

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
	i, isMap := input.(map[string]interface{})
	if !isMap {
		return false
	}

	value, isValueString := i["value"].(string)
	_, isIssuerString := i["issuer"].(string)
	_, isCurrencyString := i["currency"].(string)

	result := len(maputils.GetKeys(i)) == ISSUED_CURRENCY_SIZE && isValueString && isIssuerString && isCurrencyString && typecheck.IsFloat(value)

	return result
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
