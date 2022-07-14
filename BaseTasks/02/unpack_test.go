package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestUnpack(t *testing.T) {
	got, err := Unpack("a4bc2d5e")
	want := "aaaabccddddde"
	fmt.Printf("got: %s\nwant: %s\n", got, want)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !strings.EqualFold(got, want) {
		t.Errorf("got: %q\nwant: %q", got, want)
	}
}

func TestUnpack2(t *testing.T) {
	got, err := Unpack("abcd")
	want := "abcd"
	fmt.Printf("got: %s\nwant: %s\n", got, want)
	if err != nil {
		t.Errorf(err.Error())
	} else if !strings.EqualFold(got, want) {
		t.Errorf("got: %q\nwant: %q", got, want)
	}
}

func TestUnpack3(t *testing.T) {
	got, err := Unpack("45")
	want := ""
	wantErr := "unpack: Incorrect Input"
	fmt.Printf("got: %s\nwant: %s\n", got, want)
	fmt.Printf("err: %s\nwantErr: %s\n", err.Error(), wantErr)
	if !strings.EqualFold(got, want) || !strings.EqualFold(err.Error(), wantErr) {
		t.Errorf("got: %q\nwant: %q", got, want)
	}
}

func TestUnpack4(t *testing.T) {
	got, err := Unpack("")
	want := ""
	fmt.Printf("got: %s\nwant: %s\n", got, want)
	if err != nil {
		t.Errorf(err.Error())
	} else if !strings.EqualFold(got, want) {
		t.Errorf("got: %q\nwant: %q", got, want)
	}
}
