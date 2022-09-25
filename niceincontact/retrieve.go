package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/grokify/mogo/bytes/bytesutil"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/spectrum/openapi2"
)

func GetBodyBytes(reqURL string) ([]byte, error) {
	emptyBytes := []byte("")
	resp, err := http.Get(reqURL)
	if err != nil {
		return emptyBytes, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyBytes, err
	}
	return bytesutil.TrimUTF8BOM(body), nil
}

func GetWriteBodyBytes(reqURL, storePath string, fileMode os.FileMode) error {
	bodyBytes, err := GetBodyBytes(reqURL)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(storePath, bodyBytes, fileMode)
}

func main() {
	apisString := "admin,agent,authentication,patron,realtimedata,reporting"
	apis := strings.Split(apisString, ",")

	rootDir := filepath.Join(os.Getenv("GOPATH"), "/src/github.com/grokify/api-specs/incontact/")

	rx := regexp.MustCompile(`^([^/]+)/[^/#]+`)
	urlFormat := "https://developer.incontact.com/content/apis/%v"

	for _, apiSlug := range apis {
		if apiSlug == "admin" {
			continue
		}
		filename := fmt.Sprintf("%v-api-docs", apiSlug)
		apiURL := fmt.Sprintf(urlFormat, filename)

		body, err := GetBodyBytes(apiURL)
		logutil.FatalErr(err)

		spec, err := openapi2.NewSpecificationFromBytes(body)
		logutil.FatalErr(err)

		apiVersionDir := filepath.Join(rootDir, apiSlug, fmt.Sprintf("v%v", spec.Info.Version))

		err = os.MkdirAll(apiVersionDir, 0755)
		logutil.FatalErr(err)

		specPath := filepath.Join(apiVersionDir, filename)

		err = ioutil.WriteFile(specPath, body, 0600)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("WROTE: %v\n", specPath)

		subPaths := map[string]string{}
		subDirs := map[string]int{}

		for _, pathInfo := range spec.Paths {
			ref := pathInfo.Ref
			m := rx.FindStringSubmatch(ref)
			fmtutil.MustPrintJSON(m)
			if len(m) > 1 {
				subDirs[m[1]] = 1
				subPaths[m[0]] = m[1]
			}
		}
		fmtutil.MustPrintJSON(subPaths)
		fmtutil.MustPrintJSON(subDirs)

		for subDir, _ := range subDirs {
			specSubDir := filepath.Join(apiVersionDir, subDir)
			fmt.Printf("CREATE_DIR: %v\n", specSubDir)
			err := os.MkdirAll(specSubDir, 0755)
			logutil.FatalErr(err)
		}

		for subPath, _ := range subPaths {
			subURL := fmt.Sprintf(urlFormat, subPath)
			fmt.Println(subURL)
			subPathFile := filepath.Join(apiVersionDir, subPath)
			err := GetWriteBodyBytes(subURL, subPathFile, 0644)
			logutil.FatalErr(err)
		}
	}

	fmt.Println("DONE")
}
