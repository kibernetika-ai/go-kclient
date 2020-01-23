package util

import (
	"encoding/base64"
	"github.com/Sirupsen/logrus"
	"github.com/kibernetika-ai/go-kclient/pkg/kclient"
	"io/ioutil"
	"reflect"
)

func Request(data map[string]interface{}, files map[string]string) kclient.ModelsArbitrary {

	var req = make(map[string]interface{})

	if data != nil {
		for k, v := range data {
			var val interface{}
			switch vv := v.(type) {
			case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
				val = vv
			default:
				logrus.Panicf("unsupported data type %s", reflect.TypeOf(vv).String())
			}
			req[k] = val
		}
	}

	if files != nil {
		for k, v := range files {
			if f, err := ioutil.ReadFile(v); err != nil {
				logrus.Panicf("unable to open file %s: %s", v, err.Error())
			} else {
				req[k] = base64.StdEncoding.EncodeToString(f)
			}
		}
	}

	return req

}
