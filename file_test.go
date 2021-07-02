package GoUtils

import (
	"testing"
)

type Str struct {
	Name string
	Year float64
	Note string
}

func TestJSONserialize(t *testing.T) {
	testing_object := &Str{Name: "John", Year: 1970, Note: "None"}
	if JSONserialize(testing_object) != `{"Name":"John","Year":1970,"Note":"None"}` {
		t.Error("JSON serialization failed")
	}
}

func TestJSONdesrialize(t *testing.T) {
	json_string := `{"Name":"John","Year":1970,"Note":"None"}`
	var output Str
	JSONdeserialize(json_string, &output)
	if output.Name != "John" || output.Note != "None" || output.Year != float64(1970) {
		t.Error("JSON deserialization failed!")
	}
}

func TestXMLserialize(t *testing.T) {
	testing_object := &Str{Name: "John", Year: 1970, Note: "None"}
	if XMLserialize(testing_object) != `<Str><Name>John</Name><Year>1970</Year><Note>None</Note></Str>` {
		t.Error("XML serialization failed")
	}
}

func TestXMLdesrialize(t *testing.T) {
	json_string := `<Str><Name>John</Name><Year>1970</Year><Note>None</Note></Str>`
	var output Str
	XMLdeserialize(json_string, &output)
	if output.Name != "John" || output.Note != "None" || output.Year != float64(1970) {
		t.Error("XML deserialization failed!")
	}
}
