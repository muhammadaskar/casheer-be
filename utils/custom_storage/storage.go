package customstorage

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Upload(filePath string, fileName string, content string) (string, error) {
	fileName = generateUUID() + ".png"
	decode, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return fileName, err
	}
	fullPath := filepath.Join(filePath, fileName)
	if err != nil {
		return fileName, err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fileName, err
	}
	defer file.Close()

	_, err = file.Write(decode)
	if err != nil {
		return fileName, err
	}
	return fileName, nil
}

func Delete(path string, fileName string) error {
	filePath := path + fileName
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func GetFileImage(path string, fileName string) (string, error) {
	// Read the entire file into a byte slice
	filePath := path + fileName
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	return base64Encoding, nil

}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func generateUUID() string {
	rand.Seed(time.Now().UnixNano())
	length := 12

	const availableChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(availableChars))
		code[i] = availableChars[randomIndex]
	}

	return string(code)
}
