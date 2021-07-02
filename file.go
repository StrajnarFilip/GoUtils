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

func JSONdeserialize(jsonString string) interface{} {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var deserialized interface{}
	err := json.Unmarshal([]byte(jsonString), deserialized)
	if err != nil {
		println(err.Error())
	}
	return deserialized
}
