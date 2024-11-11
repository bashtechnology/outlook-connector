package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"

	"github.com/nfnt/resize"
)

func ResizeImage(bodyReader io.Reader) (string, error) {
	const maxSize = 1200
	data, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return "", fmt.Errorf("unable to read image data: %w", err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("unable to decode image: %w", err)
	}
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	if width > maxSize || height > maxSize {
		var newWidth, newHeight uint
		if width > height {
			newWidth = uint(maxSize)
			newHeight = uint(height) * uint(maxSize) / uint(width)
		} else {
			newWidth = uint(width) * uint(maxSize) / uint(height)
			newHeight = uint(maxSize)
		}
		resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
		var resizedImageBuf bytes.Buffer
		err = jpeg.Encode(&resizedImageBuf, resizedImg, nil)
		if err != nil {
			return "", fmt.Errorf("unable to encode image: %w", err)
		}
		return base64.StdEncoding.EncodeToString(resizedImageBuf.Bytes()), nil
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
