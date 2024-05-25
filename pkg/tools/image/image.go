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

type ImageTool struct {
    image []byte
}

func NewImageTool() *ImageTool {
    return &ImageTool{}
}

func (imageTool *ImageTool) LoadImageFromPath(path string) {
	file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error opening image file:", err)
        return
    }
    defer file.Close()

	imageTool.image, err = ioutil.ReadAll(io.LimitReader(file, maxDocumentSize))
    if err != nil {
        fmt.Println("Error reading image file:", err)
        return
    }
}

func (imageTool *ImageTool) EncodeImageToBase64() string {
    encodedImage := base64.StdEncoding.EncodeToString(imageTool.image)
    return encodedImage
}