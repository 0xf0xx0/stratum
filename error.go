package stratum

import (
	"errors"
	"strconv"

	"github.com/bytedance/sonic"
)

type ErrorCode uint32

// the Stratum protocol does not define any error codes. Each pool
// has its own set of errors, apparently. You can define your own.
const (
	None ErrorCode = iota
)

// Error is a 2(G: or more??? see public-pool)-element json array.
type Error struct {
	Code    ErrorCode
	Message string
}

// error interface
func (e *Error) Error() string {
	return strconv.Itoa(int(e.Code)) + ": " + e.Message
}

func (e *Error) UnmarshalJSON(b []byte) error {
	res := [2]string{}
	err := sonic.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	code, err := strconv.ParseUint(res[0], 10, 32)
	if err != nil {
		return err
	}
	if len(res) < 2 {
		return errors.New("invalid error array len (less than 2)")
	}
	e.Code = ErrorCode(code)
	e.Message = res[1]
	return nil
}
func (e *Error) MarshalJSON() ([]byte, error) {
	res := [2]string{strconv.Itoa(int(e.Code)), e.Message}
	return sonic.Marshal(res)
}
