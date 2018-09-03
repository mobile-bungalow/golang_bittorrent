package torrent

import (
	"encoding/hex"
	//"fmt"
	"net/url"
	"strings"
)

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}

type Torrent struct {
	Url         string
	Hash        []byte
	TrackerList []string
	Name        string
}

func NewTorrent(uri string) *Torrent {
	t := Torrent{}
	u, err := url.Parse(uri)
	check_error(err)
	t.Url = uri
	//fmt.Print("U : ", u, "\n\n")

	query, err := url.ParseQuery(u.RawQuery)
	check_error(err)
	//fmt.Print("query : ", query, "\n\n")

	Name := query["dn"][0]
	t.Name = Name
	//fmt.Print("Name : ", Name, "\n\n")

	xt := strings.Split(query["xt"][0], ":")
	hash, err := hex.DecodeString(xt[len(xt)-1])
	check_error(err)
	t.Hash = hash
	//fmt.Print("Hash : ", hash, "\n\n")

	tr := query["tr"]
	t.TrackerList = tr
	//fmt.Print("TR : ", tr, "\n\n")
	return &t
}

func NewTorrentFile(filepath string) *Torrent {
	tf := Torrent{}
	return &tf
}
