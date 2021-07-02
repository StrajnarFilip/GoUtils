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
	println(output.Name, output.Note, output.Year)
}
