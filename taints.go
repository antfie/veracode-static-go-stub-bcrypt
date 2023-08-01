package bcrypt_stub

import (
	"errors"
	"fmt"
)

func stringTaint() string {
	var taint string
	fmt.Scanln(&taint)
	return taint
}

func errorTaint() error {
	return errors.New(stringTaint())
}
