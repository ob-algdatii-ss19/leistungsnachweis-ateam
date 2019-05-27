package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAdjMat(t *testing.T) {

	//read json-data from request
	//<nil> <nil> false true {0 0} false false false
	//{{0} {{false true true 0} {false false false 0} {false true false 0} {false false false 0}}}
	var jsonString="{{0} {{false false false 0} {false false false 0} {true false false 0} {true false true 0}}}";
	var r = ioutil.NopCloser(bytes.NewReader([]byte(jsonString))) // r type is io.ReadCloser

	decoder := json.NewDecoder(r)
	var receivedData GuiRequestData
	decoder.Decode(&receivedData)
	graphObject := buildGraphObjectFromJSON(receivedData)

	//graphObject="[[] [false false false false false] [false false false true false] [false true true false false] [false false false false false]]";

	fmt.Println("test output SHOULD BE", "[[] [false false false false false] [false false false true false] [false true true false false] [false false false false false]]")
	fmt.Println("test output in TestAdjMat", graphObject)
}
