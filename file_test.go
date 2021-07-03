package GoUtils

import (
	"os"
	"testing"
)

func Pass() {
	println("\x1b[32m=== TEST PASSED ===\x1b[0m")
}
func Fail() {
	println("\x1b[31m=== TEST FAILED ===\x1b[0m")
}

type Str struct {
	Name string
	Year float64
	Note string
}

func TestJSONserialize(t *testing.T) {
	testing_object := &Str{Name: "John", Year: 1970, Note: "None"}
	if JSONserialize(testing_object) != `{"Name":"John","Year":1970,"Note":"None"}` {
		t.Error("JSON serialization failed")
		Fail()
		return
	}
	Pass()
}

func TestJSONdesrialize(t *testing.T) {
	json_string := `{"Name":"John","Year":1970,"Note":"None"}`
	var output Str
	JSONdeserialize(json_string, &output)
	if output.Name != "John" || output.Note != "None" || output.Year != float64(1970) {
		t.Error("JSON deserialization failed!")
		Fail()
		return
	}
	Pass()
}

func TestXMLserialize(t *testing.T) {
	testing_object := &Str{Name: "John", Year: 1970, Note: "None"}
	if XMLserialize(testing_object) != `<Str><Name>John</Name><Year>1970</Year><Note>None</Note></Str>` {
		t.Error("XML serialization failed")
		Fail()
		return
	}
	Pass()
}

func TestXMLdesrialize(t *testing.T) {
	json_string := `<Str><Name>John</Name><Year>1970</Year><Note>None</Note></Str>`
	var output Str
	XMLdeserialize(json_string, &output)
	if output.Name != "John" || output.Note != "None" || output.Year != float64(1970) {
		t.Error("XML deserialization failed!")
		Fail()
		return
	}
	Pass()
}

func TestFileReadWrite(t *testing.T) {
	file_name := ".testing_file_readwrite_obfuscated_name.xyz"
	FileWrite(file_name, []byte("Write good"))
	result := string(FileRead(file_name))
	os.Remove(file_name)
	if result != "Write good" {
		t.Error("Failed to read/write")
		Fail()
		return
	}
	Pass()
}
