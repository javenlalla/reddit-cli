package reddit

import (
	"net/http"
	"testing"
)

func Test_generateFilenameFromUrl(t *testing.T) {
	type args struct {
		url string
		ext string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "jpg image",
			args: args{
				url: "https://i.imgur.com/ThRMZx5.jpg",
				ext: "jpg",
			},
			want: "46f25e262b7481f62265d0d879a2729d.jpg",
		},
		{
			name: "reddit-hosted image",
			args: args{
				url: "https://i.redd.it/cnfk33iv9sh91.jpg",
				ext: "jpg",
			},
			want: "2af6ecd6022400dc4c4fecc4714b8ab2.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateFilenameFromUrl(tt.args.url, tt.args.ext); got != tt.want {
				t.Errorf("generateFilenameFromUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directoriesFromFilename(t *testing.T) {
	tests := []struct {
		filename string
		dOne     string
		dTwo     string
	}{
		{
			filename: "46f25e262b7481f62265d0d879a2729d.jpg",
			dOne:     "4",
			dTwo:     "6f",
		},
		{
			filename: "2af6ecd6022400dc4c4fecc4714b8ab2.jpg",
			dOne:     "2",
			dTwo:     "af",
		},
	}
	for _, tt := range tests {
		if got := getTopDirectoryFromFilename(tt.filename); got != tt.dOne {
			t.Errorf("getTopDirectoryFromFilename() = %v, want %v", got, tt.dOne)
		}

		if got := getSubDirectoryFromFilename(tt.filename); got != tt.dTwo {
			t.Errorf("getSubDirectoryFromFilename() = %v, want %v", got, tt.dTwo)
		}
	}
}

func Test_getExtensionFromContentTypeHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "sample-url", nil)
	req.Header.Set("content-type", "image/jpeg")

	tests := []struct {
		contentType string
		extension   string
	}{
		{contentType: "image/jpeg", extension: "jpg"},
		{contentType: "image/jpg", extension: "jpg"},
		{contentType: "image/png", extension: "png"},
		{contentType: "image/webp", extension: "webp"},
		{contentType: "video/mp4", extension: "mp4"},
		{contentType: "image/gif", extension: "mp4"},
		{contentType: "text/html", extension: "html"},
		{contentType: "text/html;charset=UTF-8", extension: "html"},
		{contentType: "unexpected-content-type", extension: ""},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "sample-url", nil)
		req.Header.Set("content-type", tt.contentType)

		if got := getExtensionFromContentTypeHeader(req.Header); got != tt.extension {
			t.Errorf("getExtensionFromContentTypeHeader(%s) = %v, want %v", tt.contentType, got, tt.extension)
		}
	}
}

func Test_getVideoIdFromUrl(t *testing.T) {
	tests := []string{
		"https://v.redd.it/8u3caw3zm6p81",
		"https://v.redd.it/8u3caw3zm6p81/",
		"https://v.redd.it/8u3caw3zm6p81/more-uri?with=false&query=true",
		"https://v.redd.it/8u3caw3zm6p81?queryParam=1234",
	}
	want := "8u3caw3zm6p81"
	for _, test := range tests {
		if got, _ := getVideoIdFromUrl(test); got != want {
			t.Errorf("getVideoIdFromUrl() = %v, want %v", got, want)
		}
	}
}
