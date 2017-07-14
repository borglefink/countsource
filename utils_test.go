package main

import "testing"

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
func TestRoundUp(t *testing.T) {
	var feed = float64(11.45)
	var actualResult = round(feed, 1)
	var expectedResult = float64(11.5)

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestRoundDown(t *testing.T) {
	var feed = float64(11.44)
	var actualResult = round(feed, 1)
	var expectedResult = float64(11.4)

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
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
