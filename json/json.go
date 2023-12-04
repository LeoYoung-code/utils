package json

import json "github.com/bytedance/sonic"

func Obj2String(obj interface{}) string {
	objStr, err := json.MarshalToString(obj)
	if err != nil {
		return err.Error()
	}
	return objStr
}
