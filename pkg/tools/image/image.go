package images

import (
    "encoding/base64"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

var (
    imageData []byte
    maxDocumentSize int64 = 10 * 1024 * 1024
)

func LoadImageFromPath(path string) {
	file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error opening image file:", err)
        return
    }
    defer file.Close()

	imageData, err = ioutil.ReadAll(io.LimitReader(file, maxDocumentSize))
    if err != nil {
        fmt.Println("Error reading image file:", err)
        return
    }
}

func EncodeImageToBase64() (string) {
	return  base64.StdEncoding.EncodeToString(imageData)
}