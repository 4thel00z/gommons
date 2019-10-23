package utils

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

func ValidateHtml(payload string) error {
	tokenizer := html.NewTokenizer(strings.NewReader(payload))
	for {
		nextToken := tokenizer.Next()
		if nextToken == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				// Not an error, we're done and it's valid!
				return nil
			}
			return err
		}
	}
}
