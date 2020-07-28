package hello

import (
	"rsc.io/quote/v3"
)

// Hello function
func Hello() string {
	return quote.HelloV3()
}

// Proverb function
func Proverb() string {
	return quote.Concurrency()
}
