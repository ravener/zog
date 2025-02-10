package zog

import (
	"github.com/Oudwins/zog/conf"
	p "github.com/Oudwins/zog/internals"
)

// Options that can be passed to a test
type TestOption = func(test *p.Test)

// Message is a function that allows you to set a custom message for the test.
func Message(msg string) TestOption {
	return func(test *p.Test) {
		test.ErrFmt = func(e p.ZogError, p ParseCtx) {
			e.SetMessage(msg)
		}
	}
}

// MessageFunc is a function that allows you to set a custom message formatter for the test.
func MessageFunc(fn p.ErrFmtFunc) TestOption {
	return func(test *p.Test) {
		test.ErrFmt = fn
	}
}

// Options that can be passed to a `schema.New()` call
type SchemaOption = func(s ZogSchema)

func WithCoercer(c conf.CoercerFunc) SchemaOption {
	return func(s ZogSchema) {
		s.setCoercer(c)
	}
}

// Options that can be passed to a `schema.Parse()` call
type ExecOption = func(p *p.ExecCtx)

// Deprecated: use ExecOption instead
type ParsingOption = ExecOption

func WithErrFormatter(fmter p.ErrFmtFunc) ExecOption {
	return func(p *p.ExecCtx) {
		p.SetErrFormatter(fmter)
	}
}

func WithCtxValue(key string, val any) ExecOption {
	return func(p *p.ExecCtx) {
		p.Set(key, val)
	}
}
