package hayagriva

import (
	"io"

	"github.com/goccy/go-yaml"
	"github.com/scrouthtv/hayagriva-converter/lib/common"
)

type Writer struct {
	Enc *yaml.Encoder
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Enc: yaml.NewEncoder(w),
	}
}

func (w *Writer) Write(key string, e common.Entry) error {
	mp := make(map[string]HcEntry)
	var err error

	mp[key], err = ToHcEntry(e)
	if err != nil {
		return err
	}

	return w.Enc.Encode(mp)
}
