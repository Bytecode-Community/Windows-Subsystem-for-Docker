package main

//conferir https://go.dev/play/p/YbZ1niXyFBR
import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func jsonAdd(jsonLink string, name string, commandLine string, startingDirectory string, icon string) string {

	jsonFile, err := os.Open(jsonLink)
	if err != nil {
		return fmt.Sprint(err)
	}

	//unmarshal
	byteValueJSON, _ := io.ReadAll(jsonFile)
	var data map[string]interface{}
	json.Unmarshal(byteValueJSON, &data)
	defer jsonFile.Close()

	list := data["profiles"].(map[string]interface{})["list"].([]interface{})

	wsdID := 1.0

wsdIDLoop:
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].(map[string]interface{})["wsdID"] != nil {
			wsdID = list[i].(map[string]interface{})["wsdID"].(float64) + 1
			break wsdIDLoop
		}
	}

	newItem := make(map[string]interface{})
	newItem["wsdID"] = wsdID
	newItem["name"] = name
	if commandLine != "" {
		newItem["commandLine"] = commandLine
	}
	if startingDirectory != "" {
		newItem["startingDirectory"] = startingDirectory
	}
	if icon != "" {
		newItem["icon"] = icon
	}

	data["profiles"].(map[string]interface{})["list"] = append(list, newItem)

	//marshal
	byteValueJSON, err = json.Marshal(data)
	if err != nil {
		return fmt.Sprint(err)
	}

	err = os.WriteFile(jsonLink, byteValueJSON, 0644)
	if err != nil {
		return fmt.Sprint(err)
	}
	defer jsonFile.Close()

	return "success!"
}

func jsonDelete(jsonLink string, jsonValue interface{}) string {

	jsonFile, err := os.Open(jsonLink)
	if err != nil {
		return fmt.Sprint(err)
	}
	byteValueJSON, _ := io.ReadAll(jsonFile)

	//unmarshal
	var data map[string]interface{}
	json.Unmarshal(byteValueJSON, &data)
	defer jsonFile.Close()

	list := data["profiles"].(map[string]interface{})["list"].([]interface{})
	comparerParam := "name"

	switch jsonValue.(type) {
	case int:
		comparerParam = "wsdID"
	case string:
		comparerParam = "name"
	}

	var comparerValue interface{}
	if comparerParam == "wsdID" {
		comparerValue = float64(jsonValue.(int))
	}
	if comparerParam == "name" {
		comparerValue = jsonValue
	}
wsdIDLoop:
	for i := len(list) - 1; i >= 0; i-- {
		listComp := list[i].(map[string]interface{})
		if listComp[comparerParam] == comparerValue && listComp["wsdID"] != nil {
			list = append(list[:i], list[i+1:]...)
			break wsdIDLoop
		}
	}

	data["profiles"].(map[string]interface{})["list"] = list

	//marshal
	byteValueJSON, err = json.Marshal(data)
	if err != nil {
		return fmt.Sprint(err)
	}

	err = os.WriteFile(jsonLink, byteValueJSON, 0644)
	if err != nil {
		return fmt.Sprint(err)
	}
	defer jsonFile.Close()

	return "success!"
}

func main() {
	var output string
	//output = jsonAdd(`H:\\Documentos\\Projects\\Json-Golang\\message.json`, "jimmy neutron africano", "", "", "")
	/*jsonAdd parâmetros:
		caminho do JSON,
		propriedade "name",
		propriedade "commandLine" ("" se não houver),
		propriedade "startingDirectory" ("" se não houver),
		propriedade "icon" ("" se não houver)
	retorno: string (o erro ou um "success!")
	*/

	//output = jsonDelete(`H:\\Documentos\\Projects\\Json-Golang\\message.json`, "teste")
	/*jsonDelete parâmetros:
		caminho do JSON,
		wsdid(int), ou name(string)
	retorno: string (o erro ou um "success!")
	*/
	fmt.Println(output)
}
