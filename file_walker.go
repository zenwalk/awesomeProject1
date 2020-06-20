package main

import (
	"bytes"
	"crypto/md5"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fromPath, _ := filepath.Abs(os.Args[1])
	toPath, _ := filepath.Abs(os.Args[2])
	filepath.Walk(fromPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		newPath := strings.Replace(path, fromPath, toPath, 1)
		_, err = os.Stat(newPath)
		if os.IsNotExist(err) {
			println("file not exist.")
			return nil
		}

		hash, _ := checkMd5(path)
		hash2, _ := checkMd5(newPath)

		if bytes.Compare(hash, hash2) == 0 {
			//log.Println("identical file.")
		} else {
			log.Fatal("file conflict.")
		}
		return nil
	})
}

func checkMd5(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	hash := h.Sum(nil)
	return hash, nil
}
