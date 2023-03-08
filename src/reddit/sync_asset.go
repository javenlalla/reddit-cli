package reddit

import (
	"crypto/md5"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"regexp"
)

func SyncAsset(url, audioUrl, audioFilename string) {
	log.Printf("url to sync: %s", url)

	r, err := GetByUrlWithRetry(url, 3)
	if err != nil {
		log.Fatal(err)
	}

	ext := getExtensionFromContentTypeHeader(r.Header)
	if ext == "" {
		log.Fatalf("unable to determine content type from header: %s", r.Header.Get("content-type"))
	}

	// @TODO: add html edge-case logic here.

	// Generate filename.
	filename := generateFilenameFromUrl(url, ext)

	// Directories.
	dOne := getTopDirectoryFromFilename(filename)
	dTwo := getSubDirectoryFromFilename(filename)

	destPath := getAssetDestinationPath(dOne, dTwo)

	_, err = DownloadAsset(url, fmt.Sprintf("%s/%s", destPath, filename))
	if err != nil {
		log.Fatal(err)
	}

	// @TODO: If applicable, download Audio file and merge. Check for ffmpeg at start of the command before proceeding.
	if audioUrl != "" {

	}
}

func generateFilenameFromUrl(url, ext string) string {
	h := md5.New()
	h.Write([]byte(url))

	return fmt.Sprintf("%x.%s", h.Sum(nil), ext)
}

func getTopDirectoryFromFilename(f string) string {
	return string(f[0])
}

func getSubDirectoryFromFilename(f string) string {
	return f[1:3]
}

func getExtensionFromContentTypeHeader(headers http.Header) string {
	switch headers.Get("content-type") {
	case "image/jpg", "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"

	case "image/webp":
		return "webp"

	case "video/mp4", "image/gif":
		return "mp4"

	case "text/html", "text/html;charset=UTF-8":
		return "html"
	}

	return ""
}

func getAssetDestinationPath(dOne, dTwo string) string {
	rootPath := viper.Get("APP_PUBLIC_PATH")
	destPath := fmt.Sprintf("%v/%s/%s", rootPath, dOne, dTwo)

	// "/var/www/mra/public/r-media/a/a7"

	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to create destination path: %s", err)
	}

	return destPath
}

// getVideoIdFromUrl parses the provided Reddit video URL in order to extract and return the video ID.
// Regex used: v\.redd\.it\/[a-z0-9]+
// With this regex, the ID should be extracted from the following URL structures:
//   - https://v.redd.it/8u3caw3zm6p81
//   - https://v.redd.it/8u3caw3zm6p81/
//   - https://v.redd.it/8u3caw3zm6p81/more-uri?with=false&query=true
//   - https://v.redd.it/8u3caw3zm6p81?queryParam=1234
func getVideoIdFromUrl(url string) (string, error) {
	r, err := regexp.Compile("v\\.redd\\.it/([a-z0-9]+)")
	if err != nil {
		return "", err
	}

	s := r.FindStringSubmatch(url)
	if len(s) < 2 {
		return "", fmt.Errorf("reddit video id not found in url `%s`", url)
	}

	return s[1], nil
}
