package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

// test data as dataProvider
var testOk = `1
2
3
4
5`

// result data
var testOkResult = `1
2
3
4
5
`

// starts with Test.. and argument from testing module
func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err != nil {
		t.Errorf("assertion failed, data not sorted")
	}

	result := out.String()
	if result != testOkResult {
		t.Errorf("assertion failed, uniqalization failed\n %v %v", result, testOkResult)
	}
}

var invalidInput = `1
3
2
4
5`
func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(invalidInput))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err == nil {
		t.Errorf("assertion failed, error not return: %v", err)
	}
}