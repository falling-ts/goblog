package error

import "fmt"

// Throw 错误处理
func (err *Error) Throw(args ...interface{}) {
	switch len(args) {
	case 1:
		first := args[0]
		if first != nil {
			fmt.Println(first)
		}
	default:
		fmt.Println(err)
	}
}
