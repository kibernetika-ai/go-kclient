package util

import (
	"encoding/base64"
	"github.com/Sirupsen/logrus"
	"github.com/kibernetika-ai/go-kclient/pkg/kclient"
	"io/ioutil"
	"reflect"
)

func TFRequest(data map[string]interface{}, files map[string]string) kclient.ModelsArbitrary {

	var inputs = make(map[string]interface{})

	if data != nil {
		for k, v := range data {
			var dtype int
			var dval interface{}
			switch vv := v.(type) {
			case string:
				dtype = 7 // DT_STRING
				dval = base64.StdEncoding.EncodeToString([]byte(vv))
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				dtype = 9 // DT_INT64
				dval = vv
			case float32, float64:
				dtype = 2 // DT_DOUBLE
				dval = vv
			default:
				logrus.Panicf("unsupported data type %s", reflect.TypeOf(vv).String())
			}
			inputs[k] = map[string]interface{}{
				"dtype": dtype,
				"data":  dval,
			}
		}
	}

	if files != nil {
		for k, v := range files {
			if f, err := ioutil.ReadFile(v); err != nil {
				logrus.Panicf("unable to open file %s: %s", v, err.Error())
			} else {
				inputs[k] = map[string]interface{}{
					"dtype": 7, // DT_STRING,
					"data":  base64.StdEncoding.EncodeToString(f),
				}
			}
		}
	}

	return map[string]interface{}{
		"raw_input": true,
		"options": map[string]interface{}{
			"noCache": true,
		},
		"inputs": inputs,
	}

}
