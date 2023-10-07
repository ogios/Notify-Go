package util

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var tempdir string

func CreateTempDir() error {
	dir, err := os.MkdirTemp("", YMLConfig.Tempfile.Name)
	if err != nil {
		return err
	}
	tempdir = dir
	fmt.Println("Tempfile Dir: ", tempdir)
	return nil
}

func RemoveTempDir() {
	fmt.Println("Removing Tempfile dir: ", tempdir)
	err := os.RemoveAll(tempdir)
	if err != nil {
		panic(err)
	}
}

func WriteTempFile(bytes []byte, prefix string, suffix string) (string, error) {
	file, TempFileErr := os.CreateTemp(tempdir, prefix+getRandom()+"*."+suffix)
	if TempFileErr != nil {
		return "", TempFileErr
	}
	defer file.Close()

	c, FileWriteErr := file.Write(bytes)
	if FileWriteErr != nil {
		return "", FileWriteErr
	}
	fmt.Println("Write file: ", c)
	return file.Name(), nil
}

func GetTempFile(prefix string, suffix string) (*os.File, error) {
	file, TempFileErr := os.CreateTemp(tempdir, prefix+getRandom()+"*."+suffix)
	if TempFileErr != nil {
		return nil, TempFileErr
	}
	return file, nil
}

func getRandom() string {
	str := ""
	for i := 0; i < 5; i++ {
		str += strconv.Itoa(rand.Intn(100))
	}
	return str
}
