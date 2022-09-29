package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/KETI-Hybrid/hcp-apiserver-v1/pkg/util"

	"github.com/Jeffail/gabs"
)

func OpenAndReadJsonFile(jsonFileName string) []byte {
	jsonFile, _ := os.Open(jsonFileName)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	defer jsonFile.Close()

	//unmarshal Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func PrintOutput(bytes []byte) error {
	var output util.Output
	err := json.Unmarshal(bytes, &output)
	if output.Stdout != nil {
		fmt.Print(string(output.Stdout))
	}
	if output.Stderr != nil {
		fmt.Print(string(output.Stderr))
	}
	return err
}

func PrintOutputReplaceStr(bytes []byte, old string, new string) error {
	var output util.Output
	err := json.Unmarshal(bytes, &output)
	if output.Stdout != nil {
		fmt.Print(string(output.Stdout))
	}
	if output.Stderr != nil {
		str := string(output.Stderr)
		str = strings.ReplaceAll(str, old, new)
		fmt.Print(str)
	}
	return err
}

func PrintErrMsg(bytes []byte) error {
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		return nil
	}
	outerrmsg := jsonParsed.Path("Message_").Data()
	inerrmsg := jsonParsed.Path("ErrStatus").Path("message").Data()
	if outerrmsg != nil {
		fmt.Println(outerrmsg)
	} else if inerrmsg != nil {
		fmt.Println(inerrmsg)
	} else {
		return nil
	}
	return nil
}
