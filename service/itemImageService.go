package service

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"golang.org/x/image/draw"

	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type ItemImageService struct {}

const imageFilePath = "./public/images/"


func RenameFile(fileName string) string{
  //get extension
	fileExt := filepath.Ext(fileName)
  newFileName := uuid.New().String() + fileExt
	return newFileName
}


func ResizeFile(filePath string) {
  fmt.Println(filePath)
	imageData, err := os.Open(filePath)
	if err != nil {
			fmt.Fprintln(os.Stderr, err)
      return 
	}
	defer imageData.Close()

	img, t, err := image.Decode(imageData)
	if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
	}
	fmt.Println("Type of image:", t)

	//rectange info of image
	rct := img.Bounds()
	fmt.Println("Width:", rct.Dx())
	fmt.Println("Height:", rct.Dy())

	//scale down by 1/2
	imgDst := image.NewRGBA(image.Rect(0, 0, rct.Dx()/2, rct.Dy()/2))
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, rct, draw.Over, nil)

	//create resized image file
	dst, err := os.Create(filePath) //maybe dst file path
	if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
	}
	defer dst.Close()
	return
	
  // return filePath
	// switch t {
	// case "jpeg":
	// 		if err := jpeg.Encode(dst, imgDst, &jpeg.Options{Quality: 100}); err != nil {
	// 				fmt.Fprintln(os.Stderr, err)
	// 				return
	// 		}
	// case "gif":
	// 		if err := gif.Encode(dst, imgDst, nil); err != nil {
	// 				fmt.Fprintln(os.Stderr, err)
	// 				return
	// 		}
	// case "png":
	// 		if err := png.Encode(dst, imgDst); err != nil {
	// 				fmt.Fprintln(os.Stderr, err)
	// 				return
	// 		}
	// default:
	// 		fmt.Fprintln(os.Stderr, "format error")
	// }
}