package filter

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image/jpeg"
	"os"
)

type Filter interface {
	Process(srcPath, dstPath string) error
}

type Grayscale struct {}

func (g Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	dst := imaging.Grayscale(src)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			fmt.Printf("Error while closing, %v\n", err)
		}
	}(dstFile)

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}

type Blur struct {}

func (b Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	dst := imaging.Blur(src, 3.5)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			fmt.Printf("Error while closing, %v\n", err)
		}
	}(dstFile)

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}
