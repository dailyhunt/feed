package util

type KV struct {
	key, value []byte
}

func (kv *KV) GetKey() []byte {
	return kv.key
}

func (kv *KV) GetStringKey() string {
	return string(kv.key)
}

func (kv *KV) GetValue() []byte {
	return kv.value
}

func (kv *KV) GetStringValue() string {
	if kv.value == nil {
		return ""
	}

	return string(kv.value)
}

func NewKV(k, v []byte) *KV {
	return &KV{k, v}
}
