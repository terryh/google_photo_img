package google_photo_image

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	fakeAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
)

type GooglePhoto struct {
	Url   string
	Image string
}

var (
	imageErr = errors.New("Image Error")
)

func GPhoto2Img(shareUrl string) (gp GooglePhoto, err error) {
	var exist bool

	client := &http.Client{}
	req, err := http.NewRequest("GET", shareUrl, nil)

	if err != nil {
		return gp, err
	}

	req.Header.Set("User-Agent", fakeAgent)
	resp, err := client.Do(req)

	if err != nil {
		return gp, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return gp, err
	}

	gp.Url, exist = doc.Find("meta[property='og:url']").First().Attr("content")

	if !exist {
		return gp, imageErr
	}
	gp.Image, exist = doc.Find("meta[property='og:image']").First().Attr("content")
	if !exist {
		return gp, imageErr
	}

	return gp, nil
}
