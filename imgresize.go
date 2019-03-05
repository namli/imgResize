package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {
	dirname := "."

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var name = file.Name()
		splits := strings.Split(name, ".")
		var fileType = splits[1]
		if fileType == "jpg" || fileType == "JPG" {
			src, err := imaging.Open(name)
			if err != nil {
				log.Fatalf("failed to open image: %v", err)
			}
			img, err := os.Open(name)
			if err != nil {
				log.Fatalf("failed to open image: %v", err)
			}
			im, _, err := image.DecodeConfig(img)
			fmt.Println(im.Width)
			if im.Width > im.Height {
				src = imaging.Resize(src, 4000, 0, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.36)
				src = imaging.Resize(src, 2800, 0, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.12)
				src = imaging.Resize(src, 2200, 0, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.15)
				src = imaging.Resize(src, 1920, 0, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.15)
				dst := imaging.New(1920, 1280, color.NRGBA{0, 0, 0, 0})
				dst = imaging.Paste(dst, src, image.Pt(0, 0))
				err = imaging.Save(dst, "s_1920_"+name, imaging.JPEGQuality(90))
				if err != nil {
					log.Fatalf("failed to save image: %v", err)
				}
			}
			if im.Height > im.Width {
				src = imaging.Resize(src, 0, 4000, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.36)
				src = imaging.Resize(src, 0, 2800, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.25)
				src = imaging.Resize(src, 0, 2200, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.15)
				src = imaging.Resize(src, 0, 1920, imaging.Lanczos)
				src = imaging.Sharpen(src, 0.15)
				dst := imaging.New(1280, 1920, color.NRGBA{0, 0, 0, 0})
				dst = imaging.Paste(dst, src, image.Pt(0, 0))
				err = imaging.Save(dst, "s_1920_"+name, imaging.JPEGQuality(90))
				if err != nil {
					log.Fatalf("failed to save image: %v", err)
				}
			}

			img.Close()
		}

	}
}
