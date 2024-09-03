package transactions

import (
	objectfns "github.com/Peersyst/xrpl-go/xrpl/utils/object-fns"
	typeoffns "github.com/Peersyst/xrpl-go/xrpl/utils/typeof-fns"
)

const MEMO_SIZE = 3

// IsMemo checks if the given object is a valid Memo object.
func IsMemo(obj FlatTransaction) bool {
	if obj == nil {
		return false
	}

	memo, ok := obj["Memo"].(map[string]interface{})
	if !ok {
		return false
	}

	size := len(objectfns.GetKeys(memo))

	// Check if MemoData is not a string or nil
	if memo["MemoData"] != nil && !typeoffns.IsString(memo["MemoData"]) {
		return false
	}
	validData := memo["MemoData"] == nil || typeoffns.IsHex(memo["MemoData"].(string))

	// Check if MemoFormat is not a string or nil
	if memo["MemoFormat"] != nil && !typeoffns.IsString(memo["MemoFormat"]) {
		return false
	}
	validFormat := memo["MemoFormat"] == nil || typeoffns.IsHex(memo["MemoFormat"].(string))

	// Check if MemoType is not a string or nil
	if memo["MemoType"] != nil && !typeoffns.IsString(memo["MemoType"]) {
		return false
	}
	validType := memo["MemoType"] == nil || typeoffns.IsHex(memo["MemoType"].(string))

	return size >= 1 && size <= MEMO_SIZE && validData && validFormat && validType && onlyHasFields(memo, []string{"MemoFormat", "MemoData", "MemoType"})
}

// onlyHasFields checks if the given object has only the specified fields or a subset of them.
func onlyHasFields(obj map[string]interface{}, fields []string) bool {
	fieldSet := make(map[string]struct{}, len(fields))
	for _, field := range fields {
		fieldSet[field] = struct{}{}
	}

	for key := range obj {
		if _, ok := fieldSet[key]; !ok {
			return false
		}
	}
	return true
}

const SIGNER_SIZE = 3

// IsSigner checks if the given object is a valid Signer object.
func IsSigner(obj map[string]interface{}) bool {
	signer, ok := obj["Signer"].(map[string]interface{})
	if !ok {
		return false
	}

	size := len(objectfns.GetKeys(signer))
	validAccount := signer["Account"] != nil && typeoffns.IsString(signer["Account"])
	validTxnSignature := signer["TxnSignature"] != nil && typeoffns.IsString(signer["TxnSignature"])
	validSigningPubKey := signer["SigningPubKey"] != nil && typeoffns.IsString(signer["SigningPubKey"])

	return size == SIGNER_SIZE && validAccount && validTxnSignature && validSigningPubKey

}

// IsAmount checks if the given object is a valid Amount object.
// It is a string for an XRP amount or a map for an IssuedCurrency amount.
func IsAmount(amount interface{}) bool {
	if typeoffns.IsString(amount) {
		return true
	}

	if amt, _ := typeoffns.IsMap(amount); IsIssuedCurrency(amt) {
		return true
	}

	return false
}

const ISSUED_CURRENCY_SIZE = 3

// IsIssuedCurrency checks if the given object is a valid IssuedCurrency object.
func IsIssuedCurrency(input map[string]interface{}) bool {
	_, isMap := typeoffns.IsMap(input)

	return isMap &&
		len(objectfns.GetKeys(input)) == ISSUED_CURRENCY_SIZE &&
		typeoffns.IsString(input["value"]) &&
		typeoffns.IsString(input["issuer"]) &&
		typeoffns.IsString(input["currency"])
}

// IsPathStep checks if the given map is a valid PathStep.
func IsPathStep(pathStep map[string]interface{}) bool {
	if account, ok := pathStep["account"]; ok && !typeoffns.IsString(account) {
		return false
	}
	if currency, ok := pathStep["currency"]; ok && !typeoffns.IsString(currency) {
		return false
	}
	if issuer, ok := pathStep["issuer"]; ok && !typeoffns.IsString(issuer) {
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

const STANDARD_CURRENCY_CODE_LEN = 3

// CheckIssuedCurrencyIsNotXrp checks if the given transaction map does not have an issued currenc as XRP.
func CheckIssuedCurrencyIsNotXrp(tx map[string]interface{}) {
	keys := objectfns.GetKeys(tx)
	for _, value := range keys {
		// Check if the value is an issued currency
		if result, ok := typeoffns.IsMap(tx[value]); ok && IsIssuedCurrency(result) {
			// Check if the issued currency is XRP (which is incorrect)
			currency := tx[value].(map[string]interface{})["currency"].(string)

			if len(currency) == STANDARD_CURRENCY_CODE_LEN && currency == "XRP" {
				panic("Cannot have an issued currency with a similar standard code as XRP.")
			}
		}
	}
}
