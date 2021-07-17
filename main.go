package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	for i := 2014; i <= 2040; i++ {

		path := fmt.Sprintf("%s%d%s", "pics/IMG-", i, ".jpg")
		image := gocv.IMRead(path, gocv.IMReadColor)
		fmt.Println(image.Cols())
		//fmt.Println(analysis.AHash(gocv.IMRead(path, gocv.IMReadColor)))

	}

}
