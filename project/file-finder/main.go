package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	FILEDIST = "texts"
	NAME     = "name.csv"
	JOB      = "job.csv"
)

func main() {
	// read
	fsArr := readDir(FILEDIST)
	fileNameBytes := readFileDesc(fsArr)
	fileDesc := getByteToString(fileNameBytes)

	// handling
	fmt.Println(fileDesc)
}

func readDir(dirname string) []string {
	files, err := ioutil.ReadDir(dirname)
	fileArr := make([]string, len(files))
	handleErr(err)

	for i, file := range files {
		fileSrc := fmt.Sprintf("%s/%s", FILEDIST, file.Name())
		fileArr[i] = fileSrc
	}

	return fileArr
}

func readFileDesc(fileArr []string) [][]byte {
	var (
		names []byte
		jobs  []byte
	)

	for _, fileName := range fileArr {

		if strings.Contains(fileName, NAME) {
			b, err := ioutil.ReadFile(fileName)
			handleErr(err)
			names = make([]byte, len(b))
			names = b
		} else {
			b, err := ioutil.ReadFile(fileName)
			handleErr(err)
			jobs = make([]byte, len(b))
			jobs = b
		}
	}

	b := [][]byte{names, jobs}

	return b
}

func getByteToString(b [][]byte) [][]string {
	names := make([]string, 10)
	jobs := make([]string, 10)

	for i, desc := range b {

		if i == 0 {
			names = strings.Fields(string(desc))
		} else {
			jobs = strings.Fields(string(desc))
		}

	}

	s := [][]string{names, jobs}
	return s
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
