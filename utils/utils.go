package utils

import (
	// "fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetProjectBundle(folderPath string) string {
	// find the xcodeproj file in the folder
	xcodeprojPath := GetXcodeprojPath(folderPath)
	pbxprojPath := filepath.Join(folderPath, xcodeprojPath, "project.pbxproj")

	// read the file bytes
	d, err := os.ReadFile(pbxprojPath)
	if err != nil {
		log.Fatal("[-] Failed while trying to read the buffer form file :", pbxprojPath, "\n", err.Error())
	}

	data := strings.Split(string(d), "\n")

	section := false
	foundReleaseConfig := false
	l_no := 0

	var releaseBuildConfig string
	var productBundle string

	for i := 0; i < len(data); i++ {
		line := data[i]

		if !foundReleaseConfig {
			if strings.Contains(line, "/* Begin XCConfigurationList section */") {
				section = true
			}

			if section {
				if l_no != 14 {
					l_no++
					continue
				}
				releaseBuildConfig = strings.TrimSpace(strings.Split(line, "/")[0])
				foundReleaseConfig = true
				i = 0
			}
			continue
		}

		if strings.Contains(line, releaseBuildConfig) {
			section = true
		}

		if section {
			if strings.Contains(line, "PRODUCT_BUNDLE_IDENTIFIER") {
				productBundle = strings.Split(strings.TrimSpace(strings.Split(line, "=")[1]), ";")[0]
			}
		}

	}

	return productBundle
}

func GetXcodeprojPath(folderPath string) string {
	dirs, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("[-] Error occurred while trying to read directory at path:", folderPath, "\n", err.Error())
	}

	for _, entry := range dirs {
		if entry.IsDir() {
			parts := strings.Split(entry.Name(), ".")
			if len(parts) > 1 && parts[len(parts)-1] == "xcodeproj" {
				return entry.Name()
			}
		}
	}

	log.Fatal("[-] Unable to find xcodeproj inside the folder :", folderPath)
	return ""
}

func ReadSwiftFiles (folderPath string) []string {
	// swift files starts with .swift extension
	dirs, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("[-] Error occurred while trying to read directory at path:", folderPath, "\n", err.Error())
	}

	var swiftFiles []File

	for _, entry := range dirs {
		if !entry.IsDir() && strings.Contains(entry.Name(), ".swift") {
			path := filepath.Join(folderPath, entry.Name())
			code, err := os.ReadFile(path)
			if err != nil {
				log.Fatal("[-] Failed while trying to read the buffer form file :", path, "\n", err.Error())
			}

			f := File{
				Name: entry.Name(),
				Code: code,
				Path: path,
			}
			swiftFiles = append(swiftFiles, f)
		}
	}

	return swiftFilePaths
}

func FindStrings(code []byte) {

}