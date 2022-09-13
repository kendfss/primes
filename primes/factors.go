package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/otiai10/jsonindent"
	"github.com/otiai10/primes"
	"github.com/urfave/cli"
)

var numericExp = regexp.MustCompile("([0-9]+)")

var factorize = func(ctx *cli.Context) error {
	if len(ctx.Args()) == 0 {
		return fmt.Errorf("`factorize` needs second arg like `primes f 12`")
	}
	for _, arg := range ctx.Args() {
		if err := factorPrint(arg, ctx); err != nil {
			return err
		}
	}
	return nil
}

var factorPrint = func(arg string, ctx *cli.Context) error {
	// num, err := strconv.ParseInt(ctx.Args().First(), 10, 64)
	num, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return err
	}
	factors := primes.Factorize(num)
	if ctx.Bool("json") {
		dict := factors.Powers()
		return jsonindent.NewEncoder(ctx.App.Writer).Encode(dict)
	}
	if ctx.Bool("bash") {

		for _, factor := range factors.All() {
			println(factor)
		}
		return nil
	}
	fmt.Println(factors.All())
	return nil
}
