package websocket

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/account"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/test"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func TestWebsocketClient_formatRequest(t *testing.T) {
	ws := &WebsocketClient{}
	tt := []struct {
		description string
		req         client.XRPLRequest
		id          int
		marker      any
		expected    string
		expectedErr error
	}{
		{
			description: "valid request",
			req: &account.AccountChannelsRequest{
				Account:            "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				Limit:              70,
			},
			id:     1,
			marker: nil,
			expected: `{
				"id": 1,
				"command":"account_channels",
				"account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"destination_account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"limit":70
			}`,
			expectedErr: nil,
		},
		{
			description: "valid request with marker",
			req: &account.AccountChannelsRequest{
				Account:            "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				Limit:              70,
			},
			id:     1,
			marker: "hdsohdaoidhadasd",
			expected: `{
				"id": 1,
				"command":"account_channels",
				"account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"destination_account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"limit":70,
				"marker":"hdsohdaoidhadasd"
			}`,
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			a, err := ws.formatRequest(tc.req, tc.id, tc.marker)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.JSONEq(t, tc.expected, string(a))
			}
		})
	}
}

func TestWebsocketClient_convertTransactionAddressToClassicAddress(t *testing.T) {
	ws := &WebsocketClient{}
	tests := []struct {
		name      string
		tx        map[string]interface{}
		fieldName string
		expected  map[string]interface{}
	}{
		{
			name: "No conversion for classic address",
			tx: map[string]interface{}{
				"Destination": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			fieldName: "Destination",
			expected: map[string]interface{}{
				"Destination": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
		},
		{
			name: "Field not present in transaction",
			tx: map[string]interface{}{
				"Amount": "1000000",
			},
			fieldName: "Destination",
			expected: map[string]interface{}{
				"Amount": "1000000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws.convertTransactionAddressToClassicAddress(&tt.tx, tt.fieldName)
			if reflect.DeepEqual(tt.expected, &tt.tx) {
				t.Errorf("expected %+v, result %+v", tt.expected, &tt.tx)
			}
		})
	}
}

func TestWebsocketClient_validateTransactionAddress(t *testing.T) {
	ws := &WebsocketClient{}
	tests := []struct {
		name         string
		tx           map[string]interface{}
		addressField string
		tagField     string
		expected     map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Valid classic address without tag",
			tx: map[string]interface{}{
				"Account": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			addressField: "Account",
			tagField:     "SourceTag",
			expected: map[string]interface{}{
				"Account": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			expectedErr: nil,
		},
		{
			name: "Valid classic address with tag",
			tx: map[string]interface{}{
				"Destination":    "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"DestinationTag": uint32(12345),
			},
			addressField: "Destination",
			tagField:     "DestinationTag",
			expected: map[string]interface{}{
				"Destination":    "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"DestinationTag": uint32(12345),
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ws.validateTransactionAddress(&tt.tx, tt.addressField, tt.tagField)
			
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			
			if !reflect.DeepEqual(tt.expected, tt.tx) {
				t.Errorf("Expected %v, but got %v", tt.expected, tt.tx)
			}
		})
	}
}

func TestWebsocketClient_setValidTransactionAddresses(t *testing.T) {
	tests := []struct {
		name        string
		tx          map[string]interface{}
		expected    map[string]interface{}
		expectedErr error
	}{
		{
			name: "Valid transaction with classic addresses",
			tx: map[string]interface{}{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
			},
			expected: map[string]interface{}{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
			},
			expectedErr: nil,
		},
		{
			name: "Transaction with additional address fields",
			tx: map[string]interface{}{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
				"Owner":       "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"RegularKey":  "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			expected: map[string]interface{}{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
				"Owner":       "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"RegularKey":  "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			expectedErr: nil,
		},
	}

	ws := &WebsocketClient{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ws.setValidTransactionAddresses(&tt.tx)
			
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			
			if !reflect.DeepEqual(tt.expected, tt.tx) {
				t.Errorf("Expected %v, but got %v", tt.expected, tt.tx)
			}
		})
	}
}

func TestWebsocketClient_setTransactionNextValidSequenceNumber(t *testing.T) {
	tests := []struct {
		name           string
		tx             map[string]interface{}
		serverMessages []map[string]any
		expected       map[string]interface{}
		expectedErr    error
	}{
		{
			name: "Valid transaction",
			tx: map[string]interface{}{
				"Account": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"account_data": map[string]any{
							"Sequence": json.Number("42"),
						},
						"ledger_current_index": json.Number("100"),
					},
				},
			},
			expected: map[string]interface{}{
				"Account":  "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Sequence": uint64(42),
			},
			expectedErr: nil,
		},
		{
			name: "Missing Account",
			tx:   map[string]interface{}{},
			serverMessages: []map[string]any{},
			expected:       map[string]interface{}{},
			expectedErr:    errors.New("missing Account in transaction"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &test.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()
			
			url, _ := test.ConvertHttpToWS(s.URL)
			cl := &WebsocketClient{
				cfg: WebsocketClientConfig{
					host: url,
				},
			}

			err := cl.setTransactionNextValidSequenceNumber(&tt.tx)

			if tt.expectedErr != nil {
				if !reflect.DeepEqual(err.Error(), tt.expectedErr.Error())  {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(tt.expected, tt.tx) {
				t.Errorf("Expected %v but got %v", tt.expected, tt.tx)
				
			}
		})
	}
}


func TestWebsocket_calculateFeePerTransactionType(t *testing.T) {
	tests := []struct {
		name            string
		tx              map[string]interface{}
		serverMessages  []map[string]any
		expectedFee     string
		expectedErr     error
		feeCushion 		float32
	}{
		{
			name: "Basic fee calculation",
			tx: map[string]interface{}{
				"TransactionType": transactions.PaymentTx,
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"info": map[string]any{
							"validated_ledger": map[string]any{
								"base_fee_xrp": float32(0.00001),
							},
							"load_factor": float32(1),
						},
					},
				},
			},
			expectedFee: "10",
			expectedErr: nil,
			feeCushion: 1,
		},
		{
			name: "Fee calculation with high load factor",
			tx: map[string]interface{}{
				"TransactionType": transactions.PaymentTx,
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"info": map[string]any{
							"validated_ledger": map[string]any{
								"base_fee_xrp": float32(0.00001),
							},
							"load_factor": float32(1000),
						},
					},
				},
			},
			expectedFee: "10000",
			expectedErr: nil,
			feeCushion: 1,
		},
		{
			name: "Fee calculation with max fee limit",
			tx: map[string]interface{}{
				"TransactionType": transactions.PaymentTx,
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"info": map[string]any{
							"validated_ledger": map[string]any{
								"base_fee_xrp": float32(1),
							},
							"load_factor": float32(1000),
						},
					},
				},
			},
			expectedFee: "2000000",
			expectedErr: nil,
			feeCushion: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &test.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()
			
			url, _ := test.ConvertHttpToWS(s.URL)
			cl := &WebsocketClient{
				cfg: WebsocketClientConfig{
					host: url,
					feeCushion: tt.feeCushion,
					maxFeeXRP: DEFAULT_MAX_FEE_XRP,
				},
			}

			err := cl.calculateFeePerTransactionType(&tt.tx)

			if tt.expectedErr != nil {
				if !reflect.DeepEqual(err.Error(), tt.expectedErr.Error()) {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(tt.expectedFee, tt.tx["Fee"]) {
					t.Errorf("Expected fee %v, but got %v", tt.expectedFee, tt.tx["Fee"])
				}
			}
		})
	}
}

func TestWebsocketClient_setLastLedgerSequence(t *testing.T) {
	tests := []struct {
		name            string
		serverMessages  []map[string]any
		tx              map[string]interface{}
		expectedTx      map[string]interface{}
		expectedErr     error
	}{
		{
			name: "Successfully set LastLedgerSequence",
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]interface{}{
						"ledger_index": 1000,
					},
				},
			},
			tx:              map[string]interface{}{},
			expectedTx:      map[string]interface{}{"LastLedgerSequence": uint32(1000 + LEDGER_OFFSET)},
			expectedErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &test.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()
			
			url, _ := test.ConvertHttpToWS(s.URL)
			cl := &WebsocketClient{
				cfg: WebsocketClientConfig{
					host: url,
				},
			}
			err := cl.setLastLedgerSequence(&tt.tx)

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(tt.expectedTx, tt.tx) {
					t.Errorf("Expected tx %v, but got %v", tt.expectedTx, tt.tx)
				}
			}
		})
	}
}

func TestWebsocketClient_checkAccountDeleteBlockers(t *testing.T) {
	tests := []struct {
		name            string
		address         types.Address
		serverMessages  []map[string]any
		expectedErr     error
	}{
		{
			name:    "No blockers",
			address: "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"account":          "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
						"account_objects":  []any{},
						"ledger_hash":      "4BC50C9B0D8515D3EAAE1E74B29A95804346C491EE1A95BF25E4AAB854A6A651",
						"ledger_index":     30,
						"validated":        true,
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &test.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()
			
			url, _ := test.ConvertHttpToWS(s.URL)
			cl := &WebsocketClient{
				cfg: WebsocketClientConfig{
					host: url,
				},
			}

			err := cl.checkAccountDeleteBlockers(tt.address)

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestWebsocketClient_setTransactionFlags(t *testing.T) {
	tests := []struct {
		name     string
		tx       map[string]interface{}
		expected uint32
		wantErr  bool
	}{
		{
			name: "No flags set",
			tx: map[string]interface{}{
				"TransactionType": string(transactions.PaymentTx),
			},
			expected: 0,
			wantErr:  false,
		},
		{
			name: "Flags already set",
			tx: map[string]interface{}{
				"TransactionType": string(transactions.PaymentTx),
				"Flags":           uint32(1),
			},
			expected: 1,
			wantErr:  false,
		},
		{
			name: "Missing TransactionType",
			tx: map[string]interface{}{
				"Flags": uint32(1),
			},
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WebsocketClient{}
			err := c.setTransactionFlags(&tt.tx)

			if (err != nil) != tt.wantErr {
				t.Errorf("setTransactionFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if flags, ok := tt.tx["Flags"].(uint32); !ok || flags != tt.expected {
					t.Errorf("setTransactionFlags() got = %v, want %v", flags, tt.expected)
				}
			}
		})
	}
}