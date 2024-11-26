package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOracle_EntryType(t *testing.T) {
	oracle := &Oracle{}
	assert.Equal(t, OracleEntry, oracle.EntryType())
}

func TestPriceData_Flatten(t *testing.T) {
	testcases := []struct {
		name      string
		priceData *PriceData
		expected  FlatPriceData
	}{
		{
			name:      "pass - empty",
			priceData: &PriceData{},
			expected: FlatPriceData{
				"Scale": uint8(0),
			},
		},
		{
			name: "pass - complete",
			priceData: &PriceData{
				BaseAsset:  "XRP",
				QuoteAsset: "USD",
				AssetPrice: 740,
				Scale:      3,
			},
			expected: FlatPriceData{
				"BaseAsset":  "XRP",
				"QuoteAsset": "USD",
				"AssetPrice": uint64(740),
				"Scale":      uint8(3),
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			assert.Equal(t, testcase.priceData.Flatten(), testcase.expected)
		})
	}
}

func TestPriceData_Validate(t *testing.T) {
	testcases := []struct {
		name      string
		priceData *PriceData
		expected  error
	}{
		{
			name:      "fail - empty",
			priceData: &PriceData{},
			expected:  ErrPriceDataBaseAsset,
		},
		{
			name: "fail - empty quote asset",
			priceData: &PriceData{
				BaseAsset: "XRP",
			},
			expected: ErrPriceDataQuoteAsset,
		},
		{
			name: "fail - scale greater than max",
			priceData: &PriceData{
				BaseAsset:  "XRP",
				QuoteAsset: "USD",
				Scale:      11,
			},
			expected: ErrPriceDataScale,
		},
		{
			name: "fail - asset price and scale not set together",
			priceData: &PriceData{
				BaseAsset:  "XRP",
				QuoteAsset: "USD",
				AssetPrice: 740,
			},
			expected: ErrPriceDataAssetPriceAndScale,
		},
		{
			name: "pass - complete",
			priceData: &PriceData{
				BaseAsset:  "XRP",
				QuoteAsset: "USD",
				AssetPrice: 740,
				Scale:      3,
			},
			expected: nil,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			assert.Equal(t, testcase.priceData.Validate(), testcase.expected)
		})
	}
}
