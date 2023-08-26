package util

import (
	"fmt"
	"os"

	"github.com/google/uuid"
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

func WriteTempFile(bytes []byte, suffix string) (string, error) {
	id2, UUIDErr := uuid.NewRandom()
	if UUIDErr != nil {
		return "", UUIDErr
	}
	file, TempFileErr := os.CreateTemp(tempdir, id2.String()+"*."+suffix)
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
