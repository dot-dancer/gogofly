package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

var stRootDir02 string
var stSeparator02 string
var iRootNode Node

const stJsonFileName02 = "dir.json"

func loadJson02() {
	stSeparator02 = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir02 = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator02)]

	gnJsonFileBytes, _ := os.ReadFile(stWorkDir + stSeparator02 + stJsonFileName02)
	err := json.Unmarshal(gnJsonFileBytes, &iRootNode)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}
}

func parseNode(iNode Node, stParentDir string) {
	if iNode.Text != "" {
		createDir02(iNode, stParentDir)
	}

	if stParentDir != "" {
		stParentDir = stParentDir + stSeparator02
	}

	if iNode.Text != "" {
		stParentDir = stParentDir + iNode.Text
	}

	for _, v := range iNode.Children {
		parseNode(v, stParentDir)
	}
}

func createDir02(iNode Node, stParentDir string) {
	stDirPath := stRootDir02 + stSeparator02
	if stParentDir != "" {
		stDirPath = stDirPath + stParentDir + stSeparator02
	}
	stDirPath = stDirPath + iNode.Text

	err := os.MkdirAll(stDirPath, fs.ModePerm)
	if err != nil {
		panic("Create Dir Error: " + err.Error())
	}
}
func TestGenerateDir02(t *testing.T) {
	loadJson02()
	parseNode(iRootNode, "")
}
