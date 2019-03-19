package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloud66-oss/cloud66"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	client       cloud66.Client
	clientId     string
	clientSecret string
	debugMode    bool   = false
	underTest    bool   = false
	VERSION      string = "dev"
	BUILD_DATE   string = ""
	tokenFile    string = "cx.json"
	fayeEndpoint string = "https://sockets.cloud66.com:443/push"
)

func main() {

	//var stencil cloud66.Stencil
	var templateFile string = "templates.json"
	var stencilFilename string  = "all-in-one.yml"

	template, _ := os.Open(templateFile)
/*	var templateData interface{}
	err := json.Unmarshal(template, &templateData)

	if err != nil {
		fmt.Printf("error while reading template.json: %s \n", err)
		return
	}
	fmt.Println("printing template.json: ")
	fmt.Println(templateData)
*/

	templateData, _ := ioutil.ReadAll(template)

	var result map[string]interface{}
	json.Unmarshal([]byte(templateData), &result)

	fmt.Println("printing template.json: ")
	result["name"] = "stencils-node-express"
	fmt.Println(result)


	stencilFile, _ := ioutil.ReadFile(stencilFilename)

	var stencilData interface{}
	err2 := yaml.Unmarshal(stencilFile, &stencilData)

	if err2 != nil {
		fmt.Printf("error while reading stencil file: %s \n", err2)
		return
	}
	fmt.Println("printing all-in-one.yml : ")
	fmt.Println(stencilData)
}