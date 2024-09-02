package validations

import (
	"errors"
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/model/utils"
	typeoffns "github.com/Peersyst/xrpl-go/xrpl/utils/typeof-fns"
)

func ValidateBaseTransaction(tx map[string]interface{}) error {
	ValidateRequiredField(tx, "TransactionType", typeoffns.IsString)
	ValidateRequiredField(tx, "Account", typeoffns.IsString)

	// optional fields
	ValidateOptionalField(tx, "Fee", typeoffns.IsUint64)
	ValidateOptionalField(tx, "Sequence", typeoffns.IsUint32)
	ValidateOptionalField(tx, "AccountTxnID", typeoffns.IsString)
	ValidateOptionalField(tx, "LastLedgerSequence", typeoffns.IsUint)
	ValidateOptionalField(tx, "SourceTag", typeoffns.IsUint)
	ValidateOptionalField(tx, "SigningPubKey", typeoffns.IsString)
	ValidateOptionalField(tx, "TicketSequence", typeoffns.IsUint)
	ValidateOptionalField(tx, "TxnSignature", typeoffns.IsString)
	ValidateOptionalField(tx, "NetworkID", typeoffns.IsUint)

	// memos
	validateMemos(tx)

	// signers
	validateSigners(tx)

	return nil
}

func ValidateRequiredField(tx map[string]interface{}, field string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if _, ok := tx[field]; !ok {
		return errors.New(field + " is missing")
	}

	// Check if the field is valid.
	if !checkValidity(tx[field]) {
		transactionType, _ := tx["TransactionType"].(string)
		return fmt.Errorf("%s: invalid field %s", transactionType, field)
	}

	return nil
}

// ValidateOptionalField validates an optional field in the transaction map.
func ValidateOptionalField(tx map[string]interface{}, paramName string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if value, ok := tx[paramName]; ok {
		// Check if the field is valid.
		if !checkValidity(value) {
			transactionType, _ := tx["TransactionType"].(string)
			return fmt.Errorf("%s: invalid field %s", transactionType, paramName)
		}
	}
	return nil
}

func validateMemos(tx map[string]interface{}) error {
	if tx["Memos"] != nil {
		memos, ok := tx["Memos"].([]map[string]interface{})
		if !ok {
			return errors.New("BaseTransaction: invalid Memos")
		}
		for _, memo := range memos {
			if !utils.IsMemo(memo) {
				return errors.New("BaseTransaction: invalid Memos")
			}
		}
	}
	return nil
}

func validateSigners(tx map[string]interface{}) error {
	if tx["Signers"] != nil {
		signers, ok := tx["Signers"].([]map[string]interface{})
		if !ok {
			return errors.New("BaseTransaction: invalid Signers")
		}
		for _, signer := range signers {
			if !utils.IsSigner(signer) {
				return errors.New("BaseTransaction: invalid Signers")
			}
		}
	}
	return nil
}
