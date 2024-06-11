package main

import (
	"csbparser/convertor"
	"csbparser/csbparser"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	flatbuffers "github.com/google/flatbuffers/go"
	uuid "github.com/satori/go.uuid"
)

func usage() {
	fmt.Println(`
usage:
    csbparser b2d|b2j|j2b inFile outFile`)
}

func main() {
	if len(os.Args) != 4 {
		usage()
		return
	}
	switch os.Args[1] {
	case "b2d":
		csb2csd(os.Args[2], os.Args[3])
	case "b2j":
		csb2json(os.Args[2], os.Args[3])
	case "j2b":
		json2csb(os.Args[2], os.Args[3])
	default:
		usage()
		return
	}
}

func csb2csd(inFile, outFile string) {
	data, err := os.ReadFile(inFile)
	if err != nil {
		fmt.Println("Can not read data from file `" + inFile + "`.")
		return
	}

	parser := csbparser.GetRootAsCSParseBinary(data, 0)
	content := convertor.Csb2Csd(parser)

	typeStr := "Node"
	if objectData := content.FindElement("ObjectData"); objectData != nil {
		switch objectData.SelectAttrValue("Name", "") {
		case "Scene":
			typeStr = "Scene"
		case "Layer":
			typeStr = "Layer"
			objectData.CreateAttr("ctype", "GameLayerObjectData")
		}
	}

	name := filepath.Base(inFile)
	ext := filepath.Ext(name)
	name = name[:len(name)-len(ext)]
	propertyGroup := etree.NewElement("PropertyGroup")
	propertyGroup.CreateAttr("Name", name)
	propertyGroup.CreateAttr("Type", typeStr)
	propertyGroup.CreateAttr("ID", uuid.NewV4().String())
	propertyGroup.CreateAttr("Version", "3.10.0.0")

	content_ := etree.NewElement("Content")
	content_.CreateAttr("ctype", "GameProjectContent")
	content_.AddChild(content)

	root := etree.NewElement("GameFile")
	root.AddChild(propertyGroup)
	root.AddChild(content_)

	xmlDoc := etree.NewDocumentWithRoot(root)
	xmlDoc.Indent(2)
	err = xmlDoc.WriteToFile(outFile)
	if err != nil {
		fmt.Println("Write csd data to file error:", err)
	}
}

func csb2json(inFile, outFile string) {
	data, err := os.ReadFile(inFile)
	if err != nil {
		fmt.Println("Can not read data from file `" + inFile + "`.")
		return
	}

	parser := csbparser.GetRootAsCSParseBinary(data, 0)
	jData := convertor.Csb2Json(parser)
	data, err = json.MarshalIndent(jData, "", "  ")
	if err != nil {
		fmt.Println("Encode data to json error:", err)
		return
	}

	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Can not open file `" + outFile + "` for write.")
		return
	}
	defer file.Close()
	file.Write(data)
}

func json2csb(inFile, outFile string) {
	data, err := os.ReadFile(inFile)
	if err != nil {
		fmt.Println("Can not read data from file `" + inFile + "`.")
		return
	}
	var jData map[string]interface{}
	err = json.Unmarshal(data, &jData)
	if err != nil {
		fmt.Println("Decode json from data error:", err)
		return
	}

	builder := flatbuffers.NewBuilder(0)
	builder.Finish(convertor.Json2Csb(builder, jData))
	data = builder.FinishedBytes()

	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Can not open file `" + outFile + "` for write.")
		return
	}
	defer file.Close()
	file.Write(data)
}
