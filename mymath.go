
package mymath

import "errors"

type Args struct {
	A, B int
}

type Result struct {
	Quo, Rem int
}

type MyMath int

func (t *MyMath) Divide(args *Args, quo *Result) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
