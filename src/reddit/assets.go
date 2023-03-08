package reddit

import (
	"io"
	"net/http"
	"os"
)

// DownloadAsset executes an HTTP request to download an asset from the provided source URL to the designated
// File Path.
func DownloadAsset(sourceUrl string, targetFilePath string) (int64, error) {
	f, err := initializeLocalFile(targetFilePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	r, err := GetByUrlWithRetry(sourceUrl, 3)
	if err != nil {
		return 0, err
	}

	return writeAssetResponseToFile(r, f)
}

// writeAssetResponseToFile writes the asset from a http.Response to the provided File.
// The number of bytes written is returned.
func writeAssetResponseToFile(r *http.Response, f *os.File) (int64, error) {
	defer r.Body.Close()
	n, err := io.Copy(f, r.Body)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// initializeLocalFile creates a local file at the provided File Path and returns an instance of the file.
func initializeLocalFile(targetFilePath string) (*os.File, error) {
	f, err := os.Create(targetFilePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}
