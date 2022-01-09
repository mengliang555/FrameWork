package method_util

import (
	"FrameWork/util/common"
	"errors"
	"fmt"
)

func ErrToPanic(err error) {
	if err != nil {
		common.Error(err.Error())
		panic(err.Error())
	}
}

func RecoverPanic(f func()) (err error) {
	defer func() {
		if er := recover(); er != nil {
			if v, ok := er.(error); ok {
				err = v
			} else {
				err = errors.New(fmt.Sprintf("err %v", er))
			}
		}
	}()
	f()
	return nil
}

func RecoverPanicWithError(f func() error) (err error) {
	defer func() {
		if er := recover(); er != nil {
			if v, ok := er.(error); ok {
				err = v
			} else {
				err = errors.New(fmt.Sprintf("err %v", er))
			}
		}
	}()
	err = f()
	return err
}
