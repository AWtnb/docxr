package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AWtnb/docxr/reader"
)

func checkPath(s string) bool {
	if !strings.HasSuffix(s, ".docx") {
		return false
	}
	fs, err := os.Stat(s)
	return err == nil && !fs.IsDir()
}

func run(src string) int {
	if !checkPath(src) {
		fmt.Println("invalid path:", src)
		return 1
	}

	r, err := reader.NewReader(src)
	if err != nil {
		fmt.Print(err.Error())
		return 1
	}
	defer r.Close()
	ps, err := r.ReadAll()
	if err != nil {
		fmt.Print(err.Error())
		return 1
	}

	for _, p := range ps {
		fmt.Println(p)
	}
	return 0
}

func main() {
	var (
		src string
	)
	flag.StringVar(&src, "src", "", "source file")
	flag.Parse()
	os.Exit(run(src))
}
