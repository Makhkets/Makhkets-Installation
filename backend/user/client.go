package user

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
}

func (u *User) GreetinForUser(name string) string {
	return fmt.Sprintf("Hello GOOOOOO!!!!, %s", name)

}

func DownloadFile(ctx context.Context, url, dest string, fileIndex, totalFiles int, overallProgressChan chan int) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	totalSize := resp.ContentLength
	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	var downloadedSize int64
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			if _, writeErr := file.Write(buffer[:n]); writeErr != nil {
				return writeErr
			}
			downloadedSize += int64(n)
			fileProgress := float64(downloadedSize) / float64(totalSize) * 100
			overallProgress := int(((float64(fileIndex-1) + fileProgress/100) / float64(totalFiles)) * 100)
			overallProgressChan <- overallProgress
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}
