package ris

import (
	"io"

	"github.com/scrouthtv/hayagriva-converter/lib/common"
)

type Reader struct {
	tok *Tokenizer
}

func NewReader(in io.Reader) *Reader {
	return &Reader{
		tok: NewTokenizer(in),
	}
}

func (r *Reader) ReadEntry() (common.Entry, error) {
	var entry common.Entry
	typeset := false
	loop := true

	for loop {
		tok, err := r.tok.Next()
		if err != nil {
			return entry, err
		}

		member, err := ParseToken(tok)
		if err != nil {
			return entry, err
		}

		ty, ok := member.(common.Type)
		if ok {
			if typeset {
				return entry, NewParserError("type set twice", tok.Value)
			}

			entry.Type = ty
			typeset = true
			continue
		}

		_, ok = member.(EndTag)
		if ok {
			break
		}

		entry.Information = append(entry.Information, member)
	}

	if !typeset {
		return entry, NewParserError("type not set", "")
	}

	return entry, nil
}
