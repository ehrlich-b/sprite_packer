package main

import (
	"fmt"
	"github.com/adotout/pack_2d"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: sprite_pack [directory]")
		return
	}

	dirName := args[0]

	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	packer := pack_2d.Packer2d{}
	id := 0
	images := map[int]image.Image{}
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !strings.HasSuffix(name, ".jpg") && !strings.HasSuffix(name, ".png") && !strings.HasSuffix(name, ".gif") {
			continue
		}
		imgReader, err := os.Open(filepath.Join(dirName, file.Name()))
		if err != nil {
			panic(err)
		}
		defer imgReader.Close()
		imgDecoded, _, err := image.Decode(imgReader)
		packer.AddNewBlock(imgDecoded.Bounds().Max.X, imgDecoded.Bounds().Max.Y, id)
		images[id] = imgDecoded
		id++
	}

	packedImages, maxWidth, maxHeight := packer.Pack()
	outImage := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
	for _, img := range packedImages {
		currentImage := images[img.Id]
		mX := currentImage.Bounds().Max.X
		mY := currentImage.Bounds().Max.Y
		draw.Draw(outImage, image.Rect(img.X, img.Y, img.X+mX, img.Y+mY), currentImage, image.ZP, draw.Src)
	}

	toimg, _ := os.Create("sprite.png")
	defer toimg.Close()
	png.Encode(toimg, outImage)
}
