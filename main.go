package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/font/gofont/goregular"
)

var defaultFontSize = 28.0

func main() {

	image := flag.String("img", "", "image path")
	text := flag.String("text", "", "text to write on it")
	fontSize := flag.Float64("size", defaultFontSize, "font size to write on it")

	flag.Usage = func() {
		fmt.Println("\ne.g.: go run -img=myimage.jpeg -text=1234567890 -size=28")
	}

	flag.Parse()

	if *image == "" {
		log.Fatal("img flag with image path is required")
		flag.Usage()
	}

	if *text == "" {
		flag.Usage()
		log.Fatal("text flag is required")
	}

	if err := run(*image, *text, *fontSize); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(image, text string, fontSize float64) error {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})

	backgroundImage, err := gg.LoadImage(image)
	if err != nil {
		return errors.Wrap(err, "load background image")
	}

	dc := gg.NewContextForImage(backgroundImage)
	dc.SetFontFace(face)

	dc.SetColor(color.RGBA{255, 0, 0, 255}) // set red color

	//text position
	x := float64(dc.Width() / 2)
	y := float64(10)

	dc.DrawStringAnchored(text, x, y, 0.5, 0.5)

	if err := dc.SavePNG(outputName(image)); err != nil {
		return errors.Wrap(err, "save png")
	}

	return nil
}

func outputName(oldName string) string {
	splitName := strings.Split(oldName, ".")
	return fmt.Sprintf("%s_code.%s", splitName[0], splitName[1])
}
