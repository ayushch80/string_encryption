package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

func GetProjectBundle(folderPath string) {
	// find the xcodeproj file in the folder
	xcodeprojPath := GetXcodeprojPath(folderPath)
	pbxprojPath := filepath.Join(folderPath, xcodeprojPath, "project.pbxproj")

	// read the file bytes
	data, err := os.ReadFile(pbxprojPath)
	if err != nil {
		log.Fatal("[-] Failed while trying to read the buffer form file : "+pbxprojPath+"\n", err.Error())
	}

	json := gjson.ParseBytes(data)
	fmt.Println(json)
}

func GetXcodeprojPath(folderPath string) string {
	dirs, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("[-] Error occurred while trying to read directory at path:", folderPath, "\n", err.Error())
	}

	for _, entry := range dirs {
		if entry.IsDir() && strings.Split(entry.Name(), ".")[1] == "xcodeproj" {
			return entry.Name()
		}
	}

	log.Fatal("[-] Unable to find xcodeproj inside the folder :", folderPath)
	return ""
}
