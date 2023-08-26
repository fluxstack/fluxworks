package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	bt, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

func MustMarshalToString(v interface{}) string {
	bt, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bt)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
