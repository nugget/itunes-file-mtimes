package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	plist "github.com/DHowett/go-plist"
)

type ITunes struct {
	Folder string           `plist:"Music Folder"`
	Tracks map[string]Track `plist:"Tracks"`
}

type Track struct {
	ID       int       `plist:"Track ID"`
	Name     string    `plist:"Name"`
	Added    time.Time `plist:"Date Added"`
	Location string    `plist:"Location"`
}

func massage(t Track) (bool, error) {
	var err error

	t.Location = strings.TrimPrefix(t.Location, "file://")

	t.Location, err = url.QueryUnescape(t.Location)
	if err != nil {
		return false, err
	}

	fi, err := os.Stat(t.Location)
	if err != nil {
		return false, err
	}

	if fi.ModTime().After(t.Added) {
		fmt.Printf("%v -> %v :: %s\n", fi.ModTime(), t.Added, t.Location)
		// File timestamp is later than liberary added timestamp
		err := os.Chtimes(t.Location, t.Added, t.Added)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func main() {
	filename := "Library.xml"

	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var library ITunes

	dec := plist.NewDecoder(fh)
	err = dec.Decode(&library)
	if err != nil {
		log.Fatal(err)
	}

	fh.Close()

	fCount := 0
	mCount := 0

	for _, t := range library.Tracks {
		fCount++
		changed, err := massage(t)
		if changed {
			mCount++
		}
		if err != nil {
			fmt.Printf("\n%s ERROR %v\n", t.Location, err)
		} else {
			fmt.Printf("%5d files scanned.  %5d changed.\r", fCount, mCount)
		}
	}

}
