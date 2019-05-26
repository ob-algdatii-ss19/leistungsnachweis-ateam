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
	var jsonString="<nil> <nil> false true {0 0} false false false";
	var r = ioutil.NopCloser(bytes.NewReader([]byte(jsonString))) // r type is io.ReadCloser

	decoder := json.NewDecoder(r)
	var receivedData GuiRequestData
	err := decoder.Decode(&receivedData)
	if err != nil {
		panic(err)
	}
	graphObject := buildGraphObjectFromJSON(receivedData)

	fmt.Println("test output in TestAdjMat", graphObject)


}
