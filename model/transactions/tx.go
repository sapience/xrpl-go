package transactions

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/transactions/types"
)

type Tx interface {
	TxType() TxType
}

type TxHash string

func (*TxHash) TxType() TxType {
	return HashedTx
}

type Binary struct {
	TxBlob string `json:"tx_blob"`
}

func (tx *Binary) TxType() TxType {
	return BinaryTx
}

type BaseTx struct {
	/* The unique address of the transaction sender. */
	Account types.Address
	/*
		The type of transaction. Valid types include: `Payment`, `OfferCreate`,
		`TrustSet`, and many others.
	*/
	TransactionType TxType
	/*
		Integer amount of XRP, in drops, to be destroyed as a cost for
		distributing this transaction to the network. Some transaction types have
		different minimum requirements.
	*/
	Fee types.XRPCurrencyAmount `json:",omitempty"`
	/*
		The sequence number of the account sending the transaction. A transaction
		is only valid if the Sequence number is exactly 1 greater than the previous
		transaction from the same account. The special case 0 means the transaction
		is using a Ticket instead.
	*/
	Sequence uint `json:",omitempty"`
	/*
		Hash value identifying another transaction. If provided, this transaction
		is only valid if the sending account's previously-sent transaction matches
		the provided hash.
	*/
	AccountTxnID types.Hash256 `json:",omitempty"`
	/* Set of bit-flags for this transaction. */
	Flags uint `json:",omitempty"`
	/*
		Highest ledger index this transaction can appear in. Specifying this field
		places a strict upper limit on how long the transaction can wait to be
		validated or rejected.
	*/
	LastLedgerSequence uint `json:",omitempty"`
	/*
	   Additional arbitrary information used to identify this transaction.
	*/
	Memos []MemoWrapper `json:",omitempty"`
	/* The network id of the transaction. */
	NetworkId uint `json:",omitempty"`
	/*
		Array of objects that represent a multi-signature which authorizes this
		transaction.
	*/
	Signers []Signer `json:",omitempty"`
	/*
	   Arbitrary integer used to identify the reason for this payment, or a sender
	   on whose behalf this transaction is made. Conventionally, a refund should
	   specify the initial payment's SourceTag as the refund payment's
	   DestinationTag.
	*/
	SourceTag uint `json:",omitempty"`
	/*
	  Hex representation of the public key that corresponds to the private key
	  used to sign this transaction. If an empty string, indicates a
	  multi-signature is present in the Signers field instead.
	*/
	SigningPubKey string `json:",omitempty"`
	/*
	  The sequence number of the ticket to use in place of a Sequence number. If
	  this is provided, Sequence must be 0. Cannot be used with AccountTxnID.
	*/
	TicketSequence uint `json:",omitempty"`
	/*
	   The signature that verifies this transaction as originating from the
	   account it says it is from.
	*/
	TxnSignature string `json:",omitempty"`
}

/*
ValidateBaseTx validates the given BaseTx object.
Returns an error if any of the fields are missing.
TODO: Add validation for other fields.
*/
func ValidateBaseTx(tx *BaseTx) {
	if tx.Account == "" {
		panic("base transaction: missing Account")
	}
	if tx.TransactionType == "" {
		panic("base transaction: missing TransactionType")
	}

	// TODO: validate other fields
}

func (tx *BaseTx) TxType() TxType {
	return tx.TransactionType
}

// TODO AMM support
type AMMBid struct {
	BaseTx
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}

type AMMCreate struct {
	BaseTx
}

func (*AMMCreate) TxType() TxType {
	return AMMCreateTx
}

type AMMDeposit struct {
	BaseTx
}

func (*AMMDeposit) TxType() TxType {
	return AMMDepositTx
}

type AMMVote struct {
	BaseTx
}

func (*AMMVote) TxType() TxType {
	return AMMVoteTx
}

type AMMWithdraw struct {
	BaseTx
}

func (*AMMWithdraw) TxType() TxType {
	return AMMWithdrawTx
}

func UnmarshalTx(data json.RawMessage) (Tx, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if data[0] == '"' {
		var ret TxHash
		if err := json.Unmarshal(data, &ret); err != nil {
			return nil, err
		}
		return &ret, nil
	} else if data[0] != '{' {
		// TODO error verbosity/record failed json
		return nil, fmt.Errorf("unexpected tx format; must be tx object or hash string")
	}
	// TODO AMM endpoint support
	type txTypeParser struct {
		TransactionType TxType
		TxBlob          string `json:"tx_blob"`
	}
	var txType txTypeParser
	if err := json.Unmarshal(data, &txType); err != nil {
		return nil, err
	}
	if len(txType.TxBlob) > 0 && len(txType.TransactionType) == 0 {
		return &Binary{
			TxBlob: txType.TxBlob,
		}, nil
	}
	var tx Tx
	switch txType.TransactionType {
	case AccountSetTx:
		tx = &AccountSet{}
	case AccountDeleteTx:
		tx = &AccountDelete{}
	case CheckCancelTx:
		tx = &CheckCancel{}
	case CheckCashTx:
		tx = &CheckCash{}
	case CheckCreateTx:
		tx = &CheckCreate{}
	case DepositPreauthTx:
		tx = &DepositPreauth{}
	case EscrowCancelTx:
		tx = &EscrowCancel{}
	case EscrowCreateTx:
		tx = &EscrowCreate{}
	case EscrowFinishTx:
		tx = &EscrowFinish{}
	case NFTokenAcceptOfferTx:
		tx = &NFTokenAcceptOffer{}
	case NFTokenBurnTx:
		tx = &NFTokenBurn{}
	case NFTokenCancelOfferTx:
		tx = &NFTokenCancelOffer{}
	case NFTokenCreateOfferTx:
		tx = &NFTokenCreateOffer{}
	case NFTokenMintTx:
		tx = &NFTokenMint{}
	case OfferCreateTx:
		tx = &OfferCreate{}
	case OfferCancelTx:
		tx = &OfferCancel{}
	case PaymentTx:
		tx = &Payment{}
	case PaymentChannelClaimTx:
		tx = &PaymentChannelClaim{}
	case PaymentChannelCreateTx:
		tx = &PaymentChannelCreate{}
	case PaymentChannelFundTx:
		tx = &PaymentChannelFund{}
	case SetRegularKeyTx:
		tx = &SetRegularKey{}
	case SignerListSetTx:
		tx = &SignerListSet{}
	case TrustSetTx:
		tx = &TrustSet{}
	case TicketCreateTx:
		tx = &TicketCreate{}
	default:
		return nil, fmt.Errorf("unsupported transaction type %s", txType.TransactionType)
	}
	if err := json.Unmarshal(data, tx); err != nil {
		return nil, err
	}
	return tx, nil
}
