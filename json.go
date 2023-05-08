package utils

import jsoniter "github.com/json-iterator/go"

func Obj2String(obj interface{}) string {
	objStr, err := jsoniter.MarshalToString(obj)
	if err != nil {
		return err.Error()
	}
	return objStr
}
