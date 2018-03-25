package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/terryh/google_photo_img"
)

var (
	tmpl = template.Must(template.New("gphoto2img").Parse(`<a href="{{ .Url }}"><img src="{{ .Image }}" /></a>`))
)

func usage() {
	fmt.Println("The command is used to turn a google photo share link to html embed code")
	fmt.Println("Usage:  gphoto2img your_google_photo_share_url")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	} else {
		gh, err := google_photo_image.GPhoto2Img(args[0])
		if err == nil {
			tmpl.Execute(os.Stdout, gh)
			fmt.Println()
		} else {
			fmt.Println(err)
		}

	}
}
