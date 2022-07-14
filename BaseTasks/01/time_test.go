package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	got := GetCurrentTime()
	want := time.Now()
	fmt.Printf("got: %q\nwant: %q\n", got, want)
	if got.Sub(want) > 30*time.Second {
		t.Errorf("got: %q\nwant: %q", got, want)
	}

}
