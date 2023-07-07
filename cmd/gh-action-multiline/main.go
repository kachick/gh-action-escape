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
	const usage = `Usage: gh-action-multiline <name>

$ your_command | gh-action-multiline output_name >> "$GITHUB_OUTPUT"
$ your_command | gh-action-multiline ENV_NAME -bytesize=42 >> "$GITHUB_ENV"
$ gh-action-multiline -version`

	versionFlag := flag.Bool("version", false, "print the version of this program")
	byteSizeFlag := flag.Int("bytesize", gham.ByteSizeFromGitHubDoc, "specify delimiter byte size")

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
	if byteSize < 2 {
		log.Fatalf("given byte size is too small: %d", byteSize)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	name := os.Args[1]

	nr := gham.DefaultNormalizer()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf(err.Error())
	}

	normalized, err := nr.Normalize(name, string(input), byteSize)

	if err != nil {
		log.Fatalf(err.Error())
	}

	println(normalized)
}
