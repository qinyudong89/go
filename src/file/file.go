package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
)

const (
	path = "C:\\go\\src"
)

type FileInfo interface {
	Name() string // base name of the file
	Size() int64  // length in bytes
	// Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}

func main() {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("%s", errors.Wrap(err, "read failed"))
	}
	for _, file := range files {
		fmt.Printf("Name: %v, Size: %v , IsDir: %v \n", file.Name(), file.Size(), file.IsDir())
	}
}
