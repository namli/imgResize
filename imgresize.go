package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {

	fmt.Print("Enter folder name to save photo (default:current): ")
	in := bufio.NewReader(os.Stdin)
	folderToSave, err := in.ReadString('\n')
	if err == nil {
		folderToSave = strings.Trim(folderToSave, "\n\r")
		folderToSave = strings.Trim(folderToSave, "\r\n")
		folderToSave += "/"

		if folderToSave == "/" {
			fmt.Print("You enter nothing will use current folder\n")
			folderToSave = ""
		}
	}

	fmt.Print("Enter new size of photo in px (default:1920): ")
	in = bufio.NewReader(os.Stdin)
	toSize, err := in.ReadString('\n')
	if err == nil {
		toSize = strings.Trim(toSize, "\n\r")
		toSize = strings.Trim(toSize, "\r\n")
		if toSize == "" {
			fmt.Print("You enter nothing will use 1920px\n")
			toSize = "1920"
		}
	}

	if exists(folderToSave) == false {
		os.MkdirAll(folderToSave, os.ModePerm)
		fmt.Println("Created folder: ", folderToSave)
	}
	// get list of files
	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// filter arr just for JPG
	mytest := func(s os.FileInfo) bool {
		splits := strings.Split(s.Name(), ".")
		if len(splits) > 1 {
			fileType := splits[1]
			if fileType == "jpg" || fileType == "JPG" {
				return true
			}
		}
		return false
	}
	jpgFiles := choose(files, mytest)

	// for each file let's do
	var i = 0
	for _, file := range jpgFiles {
		var name = file.Name()
		i++
		fmt.Println(i, " of ", len(jpgFiles))

		src, err := imaging.Open(name)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}

		img, err := os.Open(name)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		im, _, err := image.DecodeConfig(img)

		var horizontal bool
		if im.Width > im.Height {
			horizontal = true
		} else {
			horizontal = false
		}

		toSize, err := strconv.Atoi(toSize)

		if horizontal && toSize < im.Width {
			fmt.Println("уменьшаем ширину")
			var steps = ((im.Width - toSize) / 500)
			var proportion = (float64(im.Width) / float64(im.Height))
			fmt.Println(proportion)
			if steps > 0 {
				curSize := im.Width - 500
				for i := 0; i <= steps; i++ {
					fmt.Println(toSize, curSize)
					if curSize > toSize {
						src = imaging.Resize(src, curSize, 0, imaging.Lanczos)
						src = imaging.Sharpen(src, 0.2)
						curSize = curSize - 500
					} else {
						src = imaging.Resize(src, toSize, 0, imaging.Lanczos)
						src = imaging.Sharpen(src, 0.15)
					}
				}
				dst := imaging.New(toSize, int(float64(toSize)/proportion), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Paste(dst, src, image.Pt(0, 0))
				err = imaging.Save(dst, folderToSave+"s_"+strconv.Itoa(toSize)+"_"+name, imaging.JPEGQuality(90))
				if err != nil {
					log.Fatalf("failed to save image: %v", err)
				}
			}
		}

		if !horizontal && toSize < im.Height {
			fmt.Println("уменьшаем высоту")
			var steps = ((im.Height - toSize) / 500)
			var proportion = (float64(im.Height) / float64(im.Width))
			fmt.Println(proportion)
			if steps > 0 {
				curSize := im.Height - 500
				for i := 0; i <= steps; i++ {
					fmt.Println(toSize, curSize)
					if curSize > toSize {
						src = imaging.Resize(src, 0, curSize, imaging.Lanczos)
						src = imaging.Sharpen(src, 0.2)
						curSize = curSize - 500
					} else {
						src = imaging.Resize(src, 0, toSize, imaging.Lanczos)
						src = imaging.Sharpen(src, 0.15)
					}
				}
				dst := imaging.New(int(float64(toSize)/proportion), toSize, color.NRGBA{0, 0, 0, 0})
				dst = imaging.Paste(dst, src, image.Pt(0, 0))
				err = imaging.Save(dst, folderToSave+"s_"+strconv.Itoa(toSize)+"_"+name, imaging.JPEGQuality(90))
				if err != nil {
					log.Fatalf("failed to save image: %v", err)
				}
			}
		}
		img.Close()

	}

	fmt.Println("Press ENTER to exit ... ")
	exit, err := in.ReadString('\n')
	fmt.Println(exit)

}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func resize(imgname string, newSize int, i int) *image.NRGBA {
	src, err := imaging.Open(imgname)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
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
	return dst
}

func saveImage(img *image.NRGBA, name string, folderToSave string) {
	err := imaging.Save(img, "s_1920_"+name, imaging.JPEGQuality(90))
	fmt.Println(folderToSave)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

// this func get every element in arr and check it for given func
// return arr of compared element
func choose(ss []os.FileInfo, test func(os.FileInfo) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

// TODO 1 put resize and save to function
//
//
