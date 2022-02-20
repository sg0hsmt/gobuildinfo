package main

import (
	"debug/buildinfo"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func main() {
	var path string

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.StringVar(&path, "path", "", "path of target binary")
	flags.Parse(os.Args[1:])

	if path == "" {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			log.Fatalln("can not read build info from runtime")
		}

		fmt.Println(info)

		return
	}

	info, err := buildinfo.ReadFile(path)
	if err != nil {
		log.Fatalln("can not read build info from file:", err)
	}

	fmt.Println(info)
}
