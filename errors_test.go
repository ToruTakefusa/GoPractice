package GoPractice

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var foo = errors.New("error")
	if foo == nil {
		t.Fail()
	}
}

func TestError(t *testing.T) {
	var foo = errors.New("error")
	if "error" != foo.Error() {
		t.Fail()
	}
}

func TestUnWrap(t *testing.T) {
	var foo = fmt.Errorf("second error %w", errors.New("error"))
	if "error" != errors.Unwrap(foo).Error() {
		t.Fail()
	}
}
