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
	"encoding/xml"

	jsoniter "github.com/json-iterator/go"
)

// Do not forget to pass a pointer to input!
// Passing function parameter without "&" will fail!
func JSONserialize(input interface{}) string {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	serialized, err := json.Marshal(input)
	if err != nil {
		println(err.Error())
	}
	return string(serialized)
}

// Do not forget to pass a pointer to output!
// Passing function parameter without "&" will fail!
func JSONdeserialize(jsonString string, output interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(jsonString), output)
	if err != nil {
		println(err.Error())
	}
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

// Do not forget to pass a pointer to input!
// Passing function parameter without "&" will fail!
func XMLserialize(input interface{}) string {
	result, err := xml.Marshal(input)
	if err != nil {
		println(err.Error())
	}
	return string(result)
}

// Do not forget to pass a pointer to output!
// Passing function parameter without "&" will fail!
func XMLdeserialize(xmlString string, output interface{}) {
	err := xml.Unmarshal([]byte(xmlString), output)
	if err != nil {
		println(err.Error())
	}
}
