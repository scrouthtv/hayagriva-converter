package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/scrouthtv/hayagriva-converter/lib/common"
	"github.com/scrouthtv/hayagriva-converter/lib/hayagriva"
	"github.com/scrouthtv/hayagriva-converter/lib/ris"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatalf("usage: %s <input-file>", os.Args[0])
		log.Fatalf("usage: %s <input-file> <output-file>", os.Args[0])
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("error opening file: ", err)
	}
	defer f.Close()

	r := ris.NewReader(f)
	var entries []common.Entry
	for {
		entry, err := r.ReadEntry()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal("error reading entry: ", err)
		}

		log.Println("succesfully read")
		entries = append(entries, entry)
	}

	var outname string
	if len(os.Args) > 2 {
		outname = os.Args[2]
	} else {
		outname = strings.TrimSuffix(os.Args[1], filepath.Ext(os.Args[1])) + ".yaml"
	}

	outf, err := os.OpenFile(outname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("error creating file: ", err)
	}
	defer outf.Close()

	w := hayagriva.NewWriter(outf)

	for i, e := range entries {
		err = w.Write("ris-"+strconv.Itoa(i), e)

		if err != nil {
			log.Fatal("error writing entry: ", err)
		}
	}
}
