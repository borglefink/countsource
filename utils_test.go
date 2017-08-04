package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	notexistdirectory = "thisshouldnotbeasubdirectoryinthecurrentdirectory"
)

// determineNewline tests
func TestDetermineWindowsNewline(t *testing.T) {
	var actualResult = determineNewline("Hello\r\n")
	var expectedResult = windowsNewline

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestDetermineUnixNewline(t *testing.T) {
	var actualResult = determineNewline("Hello\n")
	var expectedResult = unixNewline

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestDetermineOldMacNewline(t *testing.T) {
	var actualResult = determineNewline("Hello\r")
	var expectedResult = oldMacNewline

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestDetermineDefaultWindowsNewline(t *testing.T) {
	var actualResult = determineNewline("Hello")
	var expectedResult = windowsNewline

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

// isInString tests
func TestIsInStringFound(t *testing.T) {
	var actualResult = isInString("Hello", []string{"xx", "He"})
	var expectedResult = true

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsInStringNotFound(t *testing.T) {
	var actualResult = isInString("Hello", []string{"xx"})
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

// isInSlice tests
func TestIsInSliceFound(t *testing.T) {
	var actualResult = isInSlice([]string{"Hello", "Hallo", "Hullu"}, "Hallo")
	var expectedResult = true

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsInSliceNotFound(t *testing.T) {
	var actualResult = isInSlice([]string{"Hello", "Hallo", "Hullu"}, "Hilly")
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

// int64ToString tests
func TestInt64ToString(t *testing.T) {
	var feed = int64(100000)
	var actualResult = int64ToString(feed, ' ')
	var expectedResult = "100 000"

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

// round tests
func TestRound(t *testing.T) {
	var roundTests = []struct {
		number    float64
		precision int
		expected  float64
	}{
		{1.45, 1, 1.5},
		{1.44, 1, 1.4},
		{2.45454, 0, 2.0},
		{2.45454, 1, 2.5},
		{2.45454, 2, 2.45},
		{2.45454, 3, 2.455},
		{2.45454, 4, 2.4545},
	}

	for _, tt := range roundTests {
		var actual = round(tt.number, tt.precision)
		if actual != tt.expected {
			t.Fatalf("Rounding %v with precision %v. Expected %v but got %v", tt.number, tt.precision, tt.expected, actual)
		}
	}
}

// isBinaryFormat tests
func TestIsBinaryFormatFalse(t *testing.T) {
	var actualResult = isBinaryFormat([]byte("<html><body><br/></body></html>"))
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsBinaryFormatTrue(t *testing.T) {
	var actualResult = isBinaryFormat([]byte("‰PNG IHDR  h     ‰"))
	var expectedResult = true

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

//func getDirectory(pathFromFlag, defaultPath string) string {
func TestGetDirectoryNotExist(t *testing.T) {
	if os.Getenv("BE_GETDIRECTORY") == "1" {
		getDirectory(notexistdirectory, "")
		return
	}

	var cmd = exec.Command(os.Args[0], "-test.run=TestGetDirectoryNotExist")
	cmd.Env = append(os.Environ(), "BE_GETDIRECTORY=1")

	var err = cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("Expected error but got %v", err)
}

func TestIsWindows(t *testing.T) {
	var actualResult = isWindows()
	var expectedResult = strings.Index(os.Getenv("OS"), "Windows") >= 0

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
