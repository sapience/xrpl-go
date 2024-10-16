package types

import (
	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

type Issue struct{}

func (i *Issue) FromJson(json any) ([]byte, error) {
	_, accountID, err := addresscodec.DecodeClassicAddressToAccountID(json.(string))

	if err != nil {
		return nil, err
	}

	return accountID, nil
}

func (i *Issue) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	if opts == nil {
		return nil, ErrNoLengthPrefix
	}
	b, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, err
	}
	return addresscodec.Encode(b, []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength), nil
}
