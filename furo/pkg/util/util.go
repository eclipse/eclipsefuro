package util

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func MkdirRelative(subdir string) {
	p := "./"
	for _, folder := range strings.Split(subdir, "/") {
		os.Mkdir(path.Clean(p+folder), 0755)
		p = p + folder + "/"
	}
}

func ScanForStringPosition(substr string, path string) string {
	f, err := os.Open(path)
	if err != nil {
		return path
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, substr) {
			p := strings.Index(t, substr) + 1
			return strings.Join([]string{path, strconv.Itoa(line), strconv.Itoa(p)}, ":")
		}

		line++
	}

	return path + ":1:1"

}
