package ris

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Tokenizer struct {
	in *bufio.Scanner
}

func NewTokenizer(in io.Reader) *Tokenizer {
	return &Tokenizer{
		in: bufio.NewScanner(in),
	}
}

type Token struct {
	Key   string
	Value string
}

func (t *Tokenizer) Next() (Token, error) {
	t.in.Scan()

	if t.in.Err() != nil {
		return Token{}, t.in.Err()
	}

	if t.in.Text() == "" {
		return Token{}, io.EOF
	}

	line := strings.TrimSuffix(t.in.Text(), "\r")
	split := strings.IndexRune(line, '-')

	if split == -1 {
		return Token{}, NewTokenizerError("key-value pair", "at least one dash")
	}

	key := strings.TrimRight(line[:split], " ")
	value := strings.TrimLeft(line[split+1:], " ")

	return Token{Key: key, Value: value}, nil
}

type TokenizerError struct {
	Expected string
	Actual   string
	Err      error
}

func NewTokenizerError(expected, actual string) TokenizerError {
	return TokenizerError{
		Expected: expected,
		Actual:   actual,
		Err:      nil,
	}
}

func WrapTokenizerError(expected, actual string, err error) TokenizerError {
	return TokenizerError{
		Expected: expected,
		Actual:   actual,
		Err:      err,
	}
}

func (e TokenizerError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("tokenizer error: expected %s, got %s: %s", e.Expected, e.Actual, e.Err.Error())
	} else {
		return fmt.Sprintf("tokenizer error: expected %s, got %s", e.Expected, e.Actual)
	}
}

func (e TokenizerError) Unwrap() error {
	return e.Err
}
