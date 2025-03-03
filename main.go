package main

import (
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	fieldsToCheck    = []string{"type", "external_name"}
)

type Result struct {
	fieldName string
	value1    string
	value2    string
	reason string
}

type FeatureDef map[string]interface{}

func init() {
	var (
		dir1, dir2       string
		displayNonExists,bothWay bool
	)	
	flag.StringVar(&dir1, "d1", "", "Dir 1")
	flag.StringVar(&dir2, "d2", "", "Dir 2")
	flag.BoolVar(&displayNonExists, "ne", false, "display not existing files")
	flag.BoolVar(&bothWay, "bw", false, "Compare both way")
	flag.Parse()
}

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	dir1 := flag.CommandLine.Lookup("d1").Value.String()
	dir2 := flag.CommandLine.Lookup("d2").Value.String()
	displayNonExists := flag.CommandLine.Lookup("ne").Value.String() == "true"
	bothWay := flag.CommandLine.Lookup("bw").Value.String() == "true"
	compare(dir1, dir2, displayNonExists)
	if bothWay {
		compare(dir2, dir1, displayNonExists)
	}
}

func compare(dir1, dir2 string, displayNonExists bool) {
	var (
		entries1 []fs.DirEntry
		entry1   fs.DirEntry
		err      error
		bytes    []byte
		feature1 FeatureDef
		feature2 FeatureDef
	)
	log.Printf("Comparing %s ==> %s", dir1, dir2)
	entries1, err = os.ReadDir(dir1)
	handleErr(err)
	for _, entry1 = range entries1 {
		if entry1.IsDir() {
			continue
		}
		results := []Result{}
		name := entry1.Name()
		filepath1 := filepath.Join(dir1, name)
		filepath2 := filepath.Join(dir2, name)
		bytes, err = os.ReadFile(filepath1)
		handleErr(err)
		handleErr(json.Unmarshal(bytes, &feature1))
		bytes, err = os.ReadFile(filepath2)
		if os.IsNotExist(err) {
			if displayNonExists {
				log.Printf("File %s not exists", filepath2)
			}
			continue
		}
		handleErr(json.Unmarshal(bytes, &feature2))
		fields1 := feature1["fields"]
		fields2 := feature2["fields"]
		res := compareFields(fields1, fields2, dir2)
		results = append(results, res...)
		logExport(results, filepath1, filepath2)
	}
}


func compareFields(fields1 interface{}, fields2 interface{}, secondDir string) (results []Result) {
	mFields1, _ := fields1.([]interface{})
		mFields2, _ := fields2.([]interface{})
		for _, intFields1 := range mFields1 {
			field1, _ := intFields1.(map[string]interface{})
				found := false
				fieldName1 := field1["name"].(string)
				for _, intFields2 := range mFields2 {
					field2, _ := intFields2.(map[string]interface{})
					fieldName2 := field2["name"].(string)
					if fieldName1 == fieldName2 {
						found = true
						for _, fieldName := range fieldsToCheck {
							value1 := field1[fieldName].(string)
							value2 := field2[fieldName].(string)
							if (value1 != value2) {
								results = append(results, Result{fieldName: fieldName1, 
									value1: value1, 
									value2: value2,
									reason: "Different values",
								})
							}
						}
						break
					}
				}
				if !found {
					results = append(results, Result{fieldName: fieldName1, reason: "Field does not exists in " + secondDir})
				}
		}
		return
}

func logExport(results []Result, filepath1, filepath2 string){
	if len(results) > 0 {
		log.Printf("Compared files: %s ==> %s", filepath1, filepath2)
		log.Printf("  %-30s|%-30s|%-30s|%-30s", "Field Name", "Value 1", "Value 2", "Description")	
		log.Printf("%s", strings.Repeat("-", 120))	
		for _, s := range results {
			log.Printf("  %-30s|%-30s|%-30s|%-30s", s.fieldName, s.value1, s.value2, s.reason)	
		}
	}
}