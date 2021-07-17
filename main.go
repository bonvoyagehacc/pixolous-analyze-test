package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

var resizeWidth int = 8
var resizeHeight int = 8

func main() {
	for i := 2014; i <= 2040; i++ {

		path := fmt.Sprintf("%s%d%s", "pics/IMG-", i, ".jpg")
		prepImage(path)

	}

}
func average2D(table2D gocv.Mat){
	var avg float32;
	var total :=float32(resizeWidth*resizeHeight)
	for i := 0; i<resizeWidth; i++{
		for j:=0; j<resizeHeight; j++{
			avg+=table2D.GetFloatAt(i,j)
		}
	}
	return avg/total
}
func prepImage(path string) gocv.Mat{
	img := gocv.IMRead(path, gocv.IMReadColor)
	defer img.Close()
	resizedImg := gocv.NewMat()
	defer resizedImg.Close()
	gocv.Resize(img, &resizedImg, image.Pt(resizeWidth, resizeHeight), 0, 0, gocv.InterpolationDefault)
	gocv.CvtColor(resizedImg,&resizedImg, gocv.ColorBGRToGray)
	return resizedImg


}
func aHash() {

}
