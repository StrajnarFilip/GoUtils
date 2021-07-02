/*
Copyright 2021 Filip Strajnar

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package GoUtils

import (
	"encoding/base64"

	jsoniter "github.com/json-iterator/go"
)

func JSONserialize(anything interface{}) string {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	serialized, err := json.Marshal(anything)
	if err != nil {
		println(err.Error())
	}
	return string(serialized)
}

// Do not forget to assert type of output!
// Reminder: assertion looks like: x.(T)
func JSONdeserialize(jsonString string) interface{} {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var deserialized interface{}
	err := json.Unmarshal([]byte(jsonString), deserialized)
	if err != nil {
		println(err.Error())
	}
	return deserialized
}

func ToBase64(bytes []byte) string {
	var result []byte
	base64.StdEncoding.Encode(result, bytes)
	return string(result)
}

func Base64ToString(encoded string) []byte {
	var decoded []byte
	base64.StdEncoding.Decode(decoded, []byte(encoded))
	return decoded
}
