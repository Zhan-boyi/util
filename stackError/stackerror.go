package stackError

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
)

type Error interface {
	fmt.Formatter
	error
	fmt.Stringer
	Stack() string
}

type BaseStackErr struct {
	errorInfo  string
	stack      string
	statusCode StatusCode
}

func (e *BaseStackErr) StatusCode() StatusCode {
	return e.statusCode
}
func (e *BaseStackErr) Code() int {
	return e.statusCode.Code()
}

func (e *BaseStackErr) Error() string {
	return e.String()
}

func (e *BaseStackErr) Stack() string {
	return e.stack
}

func (e *BaseStackErr) String() string {
	return e.errorInfo + "\n" + e.stack
}

func (e *BaseStackErr) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		io.WriteString(s, e.String())
		if s.Flag('+') {
			io.WriteString(s, "\n")
			io.WriteString(s, e.Stack())
			return
		}
	}
}

func NewErr(ctx context.Context, statusCode StatusCode, detailFormat string, args ...interface{}) Error {
	return NewErrSkipN(ctx, 1, statusCode, detailFormat, args...)
}

func WrapErr(ctx context.Context, statusCode StatusCode, e error) Error {
	return WrapErrKipN(ctx, 1, statusCode, e)
}

func WrapErrKipN(ctx context.Context, skip int, statusCode StatusCode, e error) Error {
	stack := Stack(skip + 6)
	fileLine := Location(skip + 2)
	return &BaseStackErr{
		errorInfo:  fileLine + "\t " + e.Error(),
		stack:      string(stack),
		statusCode: statusCode,
	}
}

func Stack(skip int) []byte {
	buf := new(bytes.Buffer)

	for i := 4; i < skip; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fileLine := Location(i)
		fmt.Fprintf(buf, "%s\t(0x%x)\n", fileLine, pc)

	}
	return buf.Bytes()
}

func Location(deep int) string {
	_, file, line, ok := runtime.Caller(deep)
	if !ok {
		return "???:0"
	}
	file = string([]byte(file)[strings.LastIndex(file, "/")+1:])
	return file + ":" + strconv.Itoa(line)
}

func NewErrSkipN(ctx context.Context, skip int, statusCode StatusCode, detailFormat string, args ...interface{}) Error {
	stack := Stack(skip + 6)
	fileLine := Location(skip + 2)

	return &BaseStackErr{
		errorInfo:  fileLine + "\t" + fmt.Sprintf(detailFormat, args...),
		stack:      string(stack),
		statusCode: statusCode,
	}
}
