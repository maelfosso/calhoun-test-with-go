package example

import (
	"fmt"
	"io"

	// Needed for initialize side effect
	_ "image/png"
)

var file string = "this is not used"

func ExampleImage_crop() {
	var r io.Reader
	// f, err := os.Open("img.png")
	// if err != nil {
	// 	panic(err)
	// }

	img, err := Decode(r)
	if err != nil {
		panic(err)
	}

	err = Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}

	// out, err := os.Create("out.png")
	// if err != nil {
	// 	panic(err)
	// }

	var w io.Writer
	err = Encode(img, w)
	if err != nil {
		panic(err)
	}

	fmt.Println("See out.png for the cropped image.")
	// Output:
	// See out.png for the cropped image.
}
