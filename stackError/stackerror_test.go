package stackError

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestNewErr(t *testing.T) {
	PrintErr()
}

func PrintErr() {
	ctx := context.Background()
	fmt.Println(NewErr(ctx, stackError.StatusDBGeneral, "错误：%+v", 5))
	fmt.Println("---")
	fmt.Println(WrapErr(ctx, stackError.StatusDBGeneral, errors.New("dddd")))
}

func PrintErr1() {
	ctx := context.Background()
	fmt.Println(NewErr(ctx, stackError.StatusDBGeneral, "错误：%+v", 6))
	PrintErr2()
}

func PrintErr2() {
	ctx := context.Background()
	fmt.Println(NewErr(ctx, stackError.StatusDBGeneral, "错误：%+v", 7))
}
