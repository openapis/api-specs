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

	"github.com/grokify/gotilla/bytes/bytesutil"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/swaggman/swagger2"
)

func GetBodyBytes(reqUrl string) ([]byte, error) {
	emptyBytes := []byte("")
	resp, err := http.Get(reqUrl)
	if err != nil {
		return emptyBytes, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyBytes, err
	}
	return bytesutil.TrimUTF8BOM(body), nil
}

func GetWriteBodyBytes(reqUrl, storePath string, fileMode os.FileMode) error {
	bodyBytes, err := GetBodyBytes(reqUrl)
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
		apiUrl := fmt.Sprintf(urlFormat, filename)

		body, err := GetBodyBytes(apiUrl)
		if err != nil {
			log.Fatal(err)
		}

		spec, err := swagger2.NewSpecificationFromBytes(body)
		if err != nil {
			log.Fatal(err)
		}

		apiVersionDir := filepath.Join(rootDir, apiSlug, fmt.Sprintf("v%v", spec.Info.Version))

		os.MkdirAll(apiVersionDir, 0755)

		specPath := filepath.Join(apiVersionDir, filename)

		err = ioutil.WriteFile(specPath, body, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("WROTE: %v\n", specPath)

		subPaths := map[string]string{}
		subDirs := map[string]int{}

		for _, pathInfo := range spec.Paths {
			ref := pathInfo.Ref
			m := rx.FindStringSubmatch(ref)
			fmtutil.PrintJSON(m)
			if len(m) > 1 {
				subDirs[m[1]] = 1
				subPaths[m[0]] = m[1]
			}
		}
		fmtutil.PrintJSON(subPaths)
		fmtutil.PrintJSON(subDirs)

		for subDir, _ := range subDirs {
			specSubDir := filepath.Join(apiVersionDir, subDir)
			fmt.Printf("CREATE_DIR: %v\n", specSubDir)
			err := os.MkdirAll(specSubDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}

		for subPath, _ := range subPaths {
			subURL := fmt.Sprintf(urlFormat, subPath)
			fmt.Println(subURL)
			subPathFile := filepath.Join(apiVersionDir, subPath)
			err := GetWriteBodyBytes(subURL, subPathFile, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("DONE")
}
