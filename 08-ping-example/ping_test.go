package main 

import (
	"testing"
)

func TestUrl(t *testing.T) {

	_, err := HttpPing("google.com", "http")

	if err != nil {
		t.Error("Pinging a valid url with http method should not raise an error, got", err)
	}
}

func TestWrongUrl(t *testing.T) {

	_, err := HttpPing("sandals.com", "smt")

	if err != nil {
		t.Error(err)
	}
}