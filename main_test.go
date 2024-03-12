package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPopulateStaticPages(t *testing.T) {
	expected := "templates"
	result := populateStaticPages().Name()

	if result != expected {
		t.Errorf("populateStaticPages() returned unexpected result: got %v want %v",
			result, expected)
	}
}

func TestGetThemeName(t *testing.T) {
	expected := "bs4"
	result := getThemeName()

	if result != expected {
		t.Errorf("getThemeName() returned unexpected result: got %v, want %v", result, expected)
	}
}

func TestKeepAlive(t *testing.T) {
	// Redirect standard output to a buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function being tested
	keepAlive()

	// Restore standard output
	w.Close()
	os.Stdout = old

	// Read the output from the buffer
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Verify the output
	expected := "Application is Up and Running... Press Ctrl+C to stop.\n"
	result := buf.String()
	if result != expected {
		t.Errorf("keepAlive() printed unexpected output: got %q, want %q", result, expected)
	}
}
