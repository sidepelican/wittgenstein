package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSample1(t *testing.T) {
	original := "testfiles/sample1.txt"
	expected := "testfiles/sample1_expected.txt"

	testSame(original, expected, t)
}

func TestSample2(t *testing.T) {
	original := "testfiles/sample2.txt"
	expected := "testfiles/sample2_expected.txt"

	testSame(original, expected, t)
}

func testSame(original, expected string, t *testing.T) {
	fw, err := ioutil.TempFile("", filepath.Base(original))
	if err != nil {
		t.Fatal(err)
	}
	defer fw.Close()

	originalFile, err := os.Open(original)
	if err != nil {
		t.Fatal(err)
	}
	defer originalFile.Close()

	io.Copy(fw, originalFile)

	testing := fw.Name()

	err = replace(testing)
	if err != nil {
		t.Fatal(err)
	}

	if !isSameFile(testing, expected) {
		diff := runCommand("diff " + testing + " " + expected)
		t.Fatal("replacing failed: ", diff)
	}
}
