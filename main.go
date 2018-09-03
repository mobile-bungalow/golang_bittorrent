package main

import (
	"./src/torrent"
	"flag"
	"fmt"
)

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var url = flag.String("url", "", "The magnet link for a torrent file.")
	var fpath = flag.String("file", "", "The relative path link for a torrent file.")
	flag.Parse()

	if *url != "" {
		t := torrent.NewTorrent(*url)
		fmt.Print("\n\naccepted uri : ", t.Name, " beggining torrent.\n")
		//fmt.Print(t)
		//function call for url shits
	}

	if *fpath != "" {
		t := torrent.NewTorrentFile(*fpath)
		fmt.Print("\n\naccepted file : ", *fpath, " beggining torrent.\n")
		fmt.Print(t)
		//function call for file path shits
	}
}
