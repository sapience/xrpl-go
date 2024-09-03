package transactions

type MemoWrapper struct {
	Memo Memo
}

type FlatMemoWrapper map[string]interface{}

type Memo struct {
	MemoData   string `json:",omitempty"`
	MemoFormat string `json:",omitempty"`
	MemoType   string `json:",omitempty"`
}

type FlatMemo map[string]interface{}

func (mw *MemoWrapper) Flatten() FlatMemoWrapper {
	if mw.Memo != (Memo{}) {
		flattened := make(FlatMemoWrapper)
		flattened["Memo"] = mw.Memo.Flatten()
		return flattened
	}
	return nil
}

func (m *Memo) Flatten() FlatMemo {
	flattened := make(FlatMemo)

	if m.MemoData != "" {
		flattened["MemoData"] = m.MemoData
	}

	if m.MemoFormat != "" {
		flattened["MemoFormat"] = m.MemoFormat
	}

	if m.MemoType != "" {
		flattened["MemoType"] = m.MemoType
	}

	return flattened
}
