package c2

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

// AssertEqual emits testing error if a and b are not equal. Returns true if
// equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) bool {
	if !reflect.DeepEqual(a, b) {
		err := errors.New(fmt.Sprintf("\nExpected:'%v'\n  Actual:'%v'", a, b))
		t.Error(err)
		return false
	}
	return true
}

// DiffBytes will compare two byte arrays and emit formatted difference
// useful in "Golden File Testing", return true if no differences
func DiffBytes(t *testing.T, a []byte, b []byte) bool {
	f, fErr := ioutil.TempFile(os.TempDir(), "tst")
	if fErr != nil {
		panic(fErr.Error())
	}
	defer os.Remove(f.Name())
	if _, wErr := f.Write(a); wErr != nil {
		panic(wErr)
	}
	return Diff(t, a, f.Name())
}

// Diff will compare byte array to file and emit formatted difference
// useful in "Golden File Testing", return true if no differences
func Diff(t *testing.T, a []byte, b string) bool {
	f, fErr := ioutil.TempFile(os.TempDir(), "tst")
	if fErr != nil {
		panic(fErr.Error())
	}
	defer os.Remove(f.Name())
	if _, wErr := f.Write(a); wErr != nil {
		panic(wErr)
	}
	return DiffFiles(t, f.Name(), b)
}

// DiffFiles will compare two files and emit formatted difference
// useful in "Golden File Testing", return true if no differences
func DiffFiles(t *testing.T, a string, b string) bool {
	cmd := exec.Command("diff", "-U", "3", a, b)
	var outBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Run()
	if !cmd.ProcessState.Success() {
		t.Error(errors.New(outBuff.String()))
		return false
	}
	return true
}
