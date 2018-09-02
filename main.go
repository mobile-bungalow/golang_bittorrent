package main

import (
	"flag"
	"fmt"
)

func main() {
	var url = flag.String("url", "", "The magnet link for a torrent file.")
	var fpath = flag.String("file", "", "The relative path link for a torrent file.")
	flag.Parse()
	fmt.Print(*url, *fpath)

	if *url != "" {
		//function call for url shits
	}

	if *fpath != "" {
		//function call for file path shits
	}
}
