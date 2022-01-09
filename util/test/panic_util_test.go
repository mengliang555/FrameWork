package test

import (
	"FrameWork/util/method_util"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var err = errors.New("hello world")

func TestErrToPanic(t *testing.T) {
	assert.Panics(t, func() {
		method_util.ErrToPanic(err)
	})
}

func TestNotErrToPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		method_util.ErrToPanic(nil)
	})
}

func TestRecoverPanic2(t *testing.T) {
	assert.NotPanics(t, func() {
		err := method_util.RecoverPanic(func() {
			panic(errors.New("hello_world"))
		})
		assert.NotNil(t, err)
		assert.Equal(t, "hello_world", err.Error())
	})
}

func TestRecoverPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		err := method_util.RecoverPanic(func() {
			panic("hello_world")
		})
		assert.NotNil(t, err)
		assert.Equal(t, "err hello_world", err.Error())
	})
}

func TestRecoverPanicWithError(t *testing.T) {
	assert.NotPanics(t, func() {
		err := method_util.RecoverPanicWithError(func() error {
			return errors.New("hello_world")
		})
		assert.NotNil(t, err)
		assert.Equal(t, "hello_world", err.Error())
	})
}

func TestRecoverPanicWithError2(t *testing.T) {
	assert.NotPanics(t, func() {
		err := method_util.RecoverPanicWithError(func() error {
			panic("hello_world")
		})
		assert.NotNil(t, err)
		assert.Equal(t, "err hello_world", err.Error())
	})
}

func TestRecoverPanicWithError3(t *testing.T) {
	assert.NotPanics(t, func() {
		err := method_util.RecoverPanicWithError(func() error {
			panic(errors.New("hello_world"))
		})
		assert.NotNil(t, err)
		assert.Equal(t, "hello_world", err.Error())
	})
}
