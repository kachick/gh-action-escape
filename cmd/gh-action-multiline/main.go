package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	gham "gh-action-multiline"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	const usage = `Usage: gh-action-multiline [options]

$ your_command | gh-action-multiline -name=output_name >> "$GITHUB_OUTPUT"
$ your_command | gh-action-multiline -name=ENV_NAME -bytesize=42 >> "$GITHUB_ENV"
$ gh-action-multiline -version`

	nameFlag := flag.String("name", "", "specify OUTPUT property or ENV name")
	byteSizeFlag := flag.Int("bytesize", gham.ByteSizeFromGitHubDoc, "specify delimiter byte size")
	versionFlag := flag.Bool("version", false, "print the version of this program")

	flag.Usage = func() {
		// https://github.com/golang/go/issues/57059#issuecomment-1336036866
		fmt.Printf("%s", usage+"\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *versionFlag {
		revision := commit[:7]
		fmt.Printf("%s\n", "gh-action-multiline"+" "+version+" "+"("+revision+") # "+date)
		return
	}

	byteSize := *byteSizeFlag
	name := *nameFlag

	if byteSize < 2 {
		log.Fatalf("specified byte size is too small: %d", byteSize)
	}

	if flag.NArg() != 0 || name == "" {
		flag.Usage()
		os.Exit(1)
	}

	nr := gham.DefaultNormalizer()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf(err.Error())
	}

	normalized, err := nr.Normalize(name, string(input), byteSize)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(normalized)
}
