package editor

import (
	"errors"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

func ExtractImageURLs(content string) []string {
	re := regexp.MustCompile(`<img[^>]+src="([^"]+)"`)
	matches := re.FindAllStringSubmatch(content, -1)
	srcs := make([]string, len(matches))
	for i, match := range matches {
		srcs[i] = match[1]
	}
	return srcs
}

func TempToImg(content string, oldPath, newPath string) (string, error) {
	//oldPath:"/public/upload/temp/article/", newPath:"/public/upload/article/"
	imgSlice := ExtractImageURLs(content)
	for _, v := range imgSlice {
		if strings.Contains(v, oldPath) {
			u, _ := url.Parse(v)
			dir, fileName := path.Split(u.Path)
			newDir := strings.Replace(dir, oldPath, newPath, 1)
			err := os.MkdirAll("."+newDir, 0755)
			if err != nil {
				return "", errors.New("错误1")
			}
			err = os.Rename("."+dir+fileName, "."+newDir+fileName)
			if err != nil {
				return "", err
			}
			content = strings.Replace(content, dir+fileName, newDir+fileName, 1)
		}
	}
	return content, nil
}
