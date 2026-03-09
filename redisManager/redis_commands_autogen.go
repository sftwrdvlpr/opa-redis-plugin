// Code generated! DO NOT EDIT
package redisManager

import (
	"github.com/sftwrdvlpr/opa-redis-plugin/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

func registerCLIENTGETNAME(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientgetname",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClientGetName(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerECHO(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.echo",
			Decl:             types.NewFunction(types.Args(types.Any{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 interface{}
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Echo(m.RedisContext, utils.Conv(v0))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPING(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.ping",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Ping(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerQUIT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.quit",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Quit(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDEL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.del",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Del(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerUNLINK(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.unlink",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Unlink(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDUMP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.dump",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Dump(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXISTS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.exists",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Exists(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIRE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expire",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Expire(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIREAT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expireat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Time
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireAt(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIRENX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expirenx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireNX(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIREXX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expirexx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireXX(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIREGT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expiregt",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireGT(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIRELT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expirelt",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireLT(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerKEYS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.keys",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Keys(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp77452 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp77452 = append(tmp77452, term)
				}
				term := ast.ArrayTerm(tmp77452...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerMIGRATE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.migrate",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			var v4 time.Duration
			if err := ast.As(terms[4].Value, &v4); err != nil {
				return nil, err
			}

			val := rdb.Migrate(m.RedisContext, v0, v1, v2, v3, v4)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMOVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.move",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Move(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerOBJECTREFCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.objectrefcount",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ObjectRefCount(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerOBJECTENCODING(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.objectencoding",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ObjectEncoding(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerOBJECTIDLETIME(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.objectidletime",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ObjectIdleTime(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.Seconds()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPERSIST(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.persist",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Persist(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPEXPIRE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pexpire",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PExpire(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPEXPIREAT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pexpireat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Time
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PExpireAt(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPTTL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pttl",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PTTL(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.Milliseconds()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRANDOMKEY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.randomkey",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.RandomKey(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRENAME(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rename",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Rename(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRENAMENX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.renamenx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RenameNX(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRESTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.restore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Restore(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRESTOREREPLACE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.restorereplace",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.RestoreReplace(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSORT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sort",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Order", types.String{}), types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Sort
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Sort(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp39421 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp39421 = append(tmp39421, term)
				}
				term := ast.ArrayTerm(tmp39421...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSORTSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sortstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Order", types.String{}), types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.Sort
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SortStore(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSORTINTERFACES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sortinterfaces",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Order", types.String{}), types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Sort
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SortInterfaces(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp05860 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp05860 = append(tmp05860, term)
				}
				term := ast.ArrayTerm(tmp05860...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerTOUCH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.touch",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Touch(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerTTL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.ttl",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.TTL(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.Seconds()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerTYPE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.type",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Type(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerAPPEND(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.append",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Append(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDECR(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.decr",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Decr(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDECRBY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.decrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.DecrBy(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.get",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Get(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GetRange(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETSET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GetSet(m.RedisContext, v0, utils.Conv(v1))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETEX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GetEx(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETDEL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getdel",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.GetDel(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINCR(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.incr",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Incr(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINCRBY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.incrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.IncrBy(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINCRBYFLOAT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.incrbyfloat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.IncrByFloat(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMGET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.mget",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.MGet(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp67851 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp67851 = append(tmp67851, term)
				}
				term := ast.ArrayTerm(tmp67851...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerMSET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.mset",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Any{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []interface{}
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.MSet(m.RedisContext, utils.Conva(v0)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMSETNX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.msetnx",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Any{})), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []interface{}
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.MSetNX(m.RedisContext, utils.Conva(v0)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.set",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Set(m.RedisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETARGS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setargs",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Mode", types.String{}), types.NewStaticProperty("TTL", types.Number{}), types.NewStaticProperty("ExpireAt", types.Number{}), types.NewStaticProperty("Get", types.Boolean{}), types.NewStaticProperty("KeepTTL", types.Boolean{})}, nil)), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 redis.SetArgs
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetArgs(m.RedisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETEX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetEX(m.RedisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETNX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetNX(m.RedisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETXX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setxx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetXX(m.RedisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetRange(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSTRLEN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.strlen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.StrLen(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCOPY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.copy",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}, types.Boolean{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 bool
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.Copy(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETBIT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getbit",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GetBit(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETBIT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setbit",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetBit(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Start", types.Number{}), types.NewStaticProperty("End", types.Number{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.BitCount
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitCount(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPAND(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopand",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpAnd(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPOR(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopor",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpOr(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPXOR(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopxor",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpXor(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPNOT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopnot",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpNot(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITPOS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitpos",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.NewArray([]types.Type{}, types.Number{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.BitPos(m.RedisContext, v0, v1, v2...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITFIELD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitfield",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.NewArray([]types.Type{}, types.Number{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitField(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp15620 []*ast.Term
				for _, v := range r0 {
					term := ast.IntNumberTerm(int(v))
					tmp15620 = append(tmp15620, term)
				}
				term := ast.ArrayTerm(tmp15620...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSCAN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scan",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 uint64
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Scan(m.RedisContext, v0, v1, v2)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp79123 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp79123 = append(tmp79123, term)
				}
				tr0 := ast.ArrayTerm(tmp79123...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerSCANTYPE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scantype",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.String{}, types.Number{}, types.String{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 uint64
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.ScanType(m.RedisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp63672 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp63672 = append(tmp63672, term)
				}
				tr0 := ast.ArrayTerm(tmp63672...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerSSCAN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sscan",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 uint64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.SScan(m.RedisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp24999 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp24999 = append(tmp24999, term)
				}
				tr0 := ast.ArrayTerm(tmp24999...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerHSCAN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hscan",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 uint64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.HScan(m.RedisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp05489 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp05489 = append(tmp05489, term)
				}
				tr0 := ast.ArrayTerm(tmp05489...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerZSCAN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zscan",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 uint64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.ZScan(m.RedisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp93299 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp93299 = append(tmp93299, term)
				}
				tr0 := ast.ArrayTerm(tmp93299...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerHDEL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hdel",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HDel(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHEXISTS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hexists",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HExists(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHGET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hget",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HGet(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHGETALL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hgetall",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.String{}))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HGetAll(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp97984 [][2]*ast.Term
				for key, value := range r0 {
					k := ast.StringTerm(key)
					v := ast.StringTerm(value)
					tmp97984 = append(tmp97984, [2]*ast.Term{k, v})
				}
				term := ast.ObjectTerm(tmp97984...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHINCRBY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hincrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HIncrBy(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHINCRBYFLOAT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hincrbyfloat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 float64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HIncrByFloat(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHKEYS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hkeys",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HKeys(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp03769 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp03769 = append(tmp03769, term)
				}
				term := ast.ArrayTerm(tmp03769...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHLEN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hlen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HLen(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHMGET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hmget",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HMGet(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp02942 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp02942 = append(tmp02942, term)
				}
				term := ast.ArrayTerm(tmp02942...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHSET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HSet(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHMSET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hmset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HMSet(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHSETNX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hsetnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Any{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HSetNX(m.RedisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHVALS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hvals",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HVals(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp74647 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp74647 = append(tmp74647, term)
				}
				term := ast.ArrayTerm(tmp74647...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHRANDFIELD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hrandfield",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Boolean{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 bool
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HRandField(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp73302 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp73302 = append(tmp73302, term)
				}
				term := ast.ArrayTerm(tmp73302...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBLPOP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.blpop",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BLPop(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp69322 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp69322 = append(tmp69322, term)
				}
				term := ast.ArrayTerm(tmp69322...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBRPOP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.brpop",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BRPop(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp77814 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp77814 = append(tmp77814, term)
				}
				term := ast.ArrayTerm(tmp77814...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBRPOPLPUSH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.brpoplpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.BRPopLPush(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINDEX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lindex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LIndex(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINSERT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.linsert",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Any{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 interface{}
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.LInsert(m.RedisContext, v0, v1, utils.Conv(v2), utils.Conv(v3))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINSERTBEFORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.linsertbefore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LInsertBefore(m.RedisContext, v0, utils.Conv(v1), utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINSERTAFTER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.linsertafter",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LInsertAfter(m.RedisContext, v0, utils.Conv(v1), utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLLEN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.llen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.LLen(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPOP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpop",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.LPop(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPOPCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpopcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LPopCount(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp64255 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp64255 = append(tmp64255, term)
				}
				term := ast.ArrayTerm(tmp64255...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerLPOS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpos",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Rank", types.Number{}), types.NewStaticProperty("MaxLen", types.Number{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 redis.LPosArgs
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LPos(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPOSCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lposcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Rank", types.Number{}), types.NewStaticProperty("MaxLen", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.Number{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 redis.LPosArgs
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.LPosCount(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp88077 []*ast.Term
				for _, v := range r0 {
					term := ast.IntNumberTerm(int(v))
					tmp88077 = append(tmp88077, term)
				}
				term := ast.ArrayTerm(tmp88077...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerLPUSH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LPush(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPUSHX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpushx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LPushX(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LRange(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp60591 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp60591 = append(tmp60591, term)
				}
				term := ast.ArrayTerm(tmp60591...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerLREM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lrem",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LRem(m.RedisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLSET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Any{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LSet(m.RedisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLTRIM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.ltrim",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LTrim(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPOP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpop",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.RPop(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPOPCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpopcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPopCount(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp39966 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp39966 = append(tmp39966, term)
				}
				term := ast.ArrayTerm(tmp39966...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerRPOPLPUSH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpoplpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPopLPush(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPUSH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPush(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPUSHX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpushx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPushX(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLMOVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lmove",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.LMove(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBLMOVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.blmove",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			var v4 time.Duration
			if err := ast.As(terms[4].Value, &v4); err != nil {
				return nil, err
			}

			val := rdb.BLMove(m.RedisContext, v0, v1, v2, v3, v4)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSADD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SAdd(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCARD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scard",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SCard(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSDIFF(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sdiff",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SDiff(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp42781 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp42781 = append(tmp42781, term)
				}
				term := ast.ArrayTerm(tmp42781...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSDIFFSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sdiffstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SDiffStore(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSINTER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sinter",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SInter(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp80636 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp80636 = append(tmp80636, term)
				}
				term := ast.ArrayTerm(tmp80636...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSINTERSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sinterstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SInterStore(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSISMEMBER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sismember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SIsMember(m.RedisContext, v0, utils.Conv(v1))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSMISMEMBER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smismember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.NewArray([]types.Type{}, types.Boolean{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SMIsMember(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp06584 []*ast.Term
				for _, v := range r0 {
					term := ast.BooleanTerm(v)
					tmp06584 = append(tmp06584, term)
				}
				term := ast.ArrayTerm(tmp06584...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSMEMBERS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smembers",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SMembers(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp52934 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp52934 = append(tmp52934, term)
				}
				term := ast.ArrayTerm(tmp52934...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSMEMBERSMAP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smembersmap",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.NewObject([]*types.StaticProperty{}, nil)))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SMembersMap(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp99548 [][2]*ast.Term
				for key, _ := range r0 {
					k := ast.StringTerm(key)

					v := ast.ObjectTerm()

					tmp99548 = append(tmp99548, [2]*ast.Term{k, v})
				}
				term := ast.ObjectTerm(tmp99548...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSMOVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smove",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Any{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SMove(m.RedisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSPOP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.spop",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SPop(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSPOPN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.spopn",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SPopN(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp45712 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp45712 = append(tmp45712, term)
				}
				term := ast.ArrayTerm(tmp45712...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSRANDMEMBER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.srandmember",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SRandMember(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSRANDMEMBERN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.srandmembern",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SRandMemberN(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp72253 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp72253 = append(tmp72253, term)
				}
				term := ast.ArrayTerm(tmp72253...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSREM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.srem",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SRem(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSUNION(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sunion",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SUnion(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp22313 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp22313 = append(tmp22313, term)
				}
				term := ast.ArrayTerm(tmp22313...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSUNIONSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sunionstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SUnionStore(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXADD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xadd",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("NoMkStream", types.Boolean{}), types.NewStaticProperty("MaxLen", types.Number{}), types.NewStaticProperty("MaxLenApprox", types.Number{}), types.NewStaticProperty("MinID", types.String{}), types.NewStaticProperty("Approx", types.Boolean{}), types.NewStaticProperty("Limit", types.Number{}), types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.Any{})}, nil)), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XAddArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XAdd(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXDEL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xdel",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XDel(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXLEN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xlen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XLen(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XRange(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp26919 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp11185 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp11185 = append(tmp11185, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp11185...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp26919 = append(tmp26919, term)
				}
				term := ast.ArrayTerm(tmp26919...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXRANGEN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrangen",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.XRangeN(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp16670 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp47451 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp47451 = append(tmp47451, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp47451...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp16670 = append(tmp16670, term)
				}
				term := ast.ArrayTerm(tmp16670...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREVRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrevrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XRevRange(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp93982 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp16982 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp16982 = append(tmp16982, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp16982...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp93982 = append(tmp93982, term)
				}
				term := ast.ArrayTerm(tmp93982...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREVRANGEN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrevrangen",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.XRevRangeN(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp17148 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp98538 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp98538 = append(tmp98538, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp98538...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp17148 = append(tmp17148, term)
				}
				term := ast.ArrayTerm(tmp17148...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREAD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xread",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Streams", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Block", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XReadArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XRead(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp54385 []*ast.Term
				for _, v := range r0 {

					termStream := ast.StringTerm(v.Stream)

					var tmp22443 []*ast.Term
					for _, v := range v.Messages {

						termID := ast.StringTerm(v.ID)

						var tmp21838 [][2]*ast.Term
						for key, value := range v.Values {
							k := ast.StringTerm(key)

							v := ast.NullTerm()
							if s, ok := value.(string); ok {
								v = ast.StringTerm(s)
							}

							tmp21838 = append(tmp21838, [2]*ast.Term{k, v})
						}
						termValues := ast.ObjectTerm(tmp21838...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

						tmp22443 = append(tmp22443, term)
					}
					termMessages := ast.ArrayTerm(tmp22443...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Stream"), termStream}, [2]*ast.Term{ast.StringTerm("Messages"), termMessages})

					tmp54385 = append(tmp54385, term)
				}
				term := ast.ArrayTerm(tmp54385...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREADSTREAMS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xreadstreams",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XReadStreams(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp93051 []*ast.Term
				for _, v := range r0 {

					termStream := ast.StringTerm(v.Stream)

					var tmp59975 []*ast.Term
					for _, v := range v.Messages {

						termID := ast.StringTerm(v.ID)

						var tmp37336 [][2]*ast.Term
						for key, value := range v.Values {
							k := ast.StringTerm(key)

							v := ast.NullTerm()
							if s, ok := value.(string); ok {
								v = ast.StringTerm(s)
							}

							tmp37336 = append(tmp37336, [2]*ast.Term{k, v})
						}
						termValues := ast.ObjectTerm(tmp37336...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

						tmp59975 = append(tmp59975, term)
					}
					termMessages := ast.ArrayTerm(tmp59975...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Stream"), termStream}, [2]*ast.Term{ast.StringTerm("Messages"), termMessages})

					tmp93051 = append(tmp93051, term)
				}
				term := ast.ArrayTerm(tmp93051...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPCREATE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupcreate",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupCreate(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPCREATEMKSTREAM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupcreatemkstream",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupCreateMkStream(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPSETID(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupsetid",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupSetID(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPDESTROY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupdestroy",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XGroupDestroy(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPCREATECONSUMER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupcreateconsumer",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupCreateConsumer(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPDELCONSUMER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupdelconsumer",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupDelConsumer(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXREADGROUP(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xreadgroup",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("Streams", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Block", types.Number{}), types.NewStaticProperty("NoAck", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XReadGroupArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XReadGroup(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp18257 []*ast.Term
				for _, v := range r0 {

					termStream := ast.StringTerm(v.Stream)

					var tmp67090 []*ast.Term
					for _, v := range v.Messages {

						termID := ast.StringTerm(v.ID)

						var tmp68844 [][2]*ast.Term
						for key, value := range v.Values {
							k := ast.StringTerm(key)

							v := ast.NullTerm()
							if s, ok := value.(string); ok {
								v = ast.StringTerm(s)
							}

							tmp68844 = append(tmp68844, [2]*ast.Term{k, v})
						}
						termValues := ast.ObjectTerm(tmp68844...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

						tmp67090 = append(tmp67090, term)
					}
					termMessages := ast.ArrayTerm(tmp67090...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Stream"), termStream}, [2]*ast.Term{ast.StringTerm("Messages"), termMessages})

					tmp18257 = append(tmp18257, term)
				}
				term := ast.ArrayTerm(tmp18257...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXACK(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xack",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XAck(m.RedisContext, v0, v1, v2...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXPENDING(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xpending",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Lower", types.String{}), types.NewStaticProperty("Higher", types.String{}), types.NewStaticProperty("Consumers", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Number{})))}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XPending(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termCount := ast.IntNumberTerm(int(r0.Count))

				termLower := ast.StringTerm(r0.Lower)

				termHigher := ast.StringTerm(r0.Higher)

				var tmp93598 [][2]*ast.Term
				for key, value := range r0.Consumers {
					k := ast.StringTerm(key)
					v := ast.IntNumberTerm(int(value))
					tmp93598 = append(tmp93598, [2]*ast.Term{k, v})
				}
				termConsumers := ast.ObjectTerm(tmp93598...)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Count"), termCount}, [2]*ast.Term{ast.StringTerm("Lower"), termLower}, [2]*ast.Term{ast.StringTerm("Higher"), termHigher}, [2]*ast.Term{ast.StringTerm("Consumers"), termConsumers})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXPENDINGEXT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xpendingext",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Idle", types.Number{}), types.NewStaticProperty("Start", types.String{}), types.NewStaticProperty("End", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Consumer", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("Idle", types.Number{}), types.NewStaticProperty("RetryCount", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XPendingExtArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XPendingExt(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp45625 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					termConsumer := ast.StringTerm(v.Consumer)

					termIdle := ast.IntNumberTerm(int(v.Idle.Milliseconds()))

					termRetryCount := ast.IntNumberTerm(int(v.RetryCount))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Consumer"), termConsumer}, [2]*ast.Term{ast.StringTerm("Idle"), termIdle}, [2]*ast.Term{ast.StringTerm("RetryCount"), termRetryCount})

					tmp45625 = append(tmp45625, term)
				}
				term := ast.ArrayTerm(tmp45625...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXCLAIM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xclaim",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.String{}))}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XClaim(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp03234 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp55693 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp55693 = append(tmp55693, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp55693...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp03234 = append(tmp03234, term)
				}
				term := ast.ArrayTerm(tmp03234...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXCLAIMJUSTID(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xclaimjustid",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.String{}))}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XClaimJustID(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp34714 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp34714 = append(tmp34714, term)
				}
				term := ast.ArrayTerm(tmp34714...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXAUTOCLAIM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xautoclaim",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Start", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Consumer", types.String{})}, nil)), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)), types.String{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XAutoClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XAutoClaim(m.RedisContext, v0)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp28630 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp63597 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp63597 = append(tmp63597, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp63597...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp28630 = append(tmp28630, term)
				}
				tr0 := ast.ArrayTerm(tmp28630...)
				tr1 := ast.StringTerm(r1)
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerXAUTOCLAIMJUSTID(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xautoclaimjustid",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Start", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Consumer", types.String{})}, nil)), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.String{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XAutoClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XAutoClaimJustID(m.RedisContext, v0)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp17055 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp17055 = append(tmp17055, term)
				}
				tr0 := ast.ArrayTerm(tmp17055...)
				tr1 := ast.StringTerm(r1)
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerXTRIM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrim",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrim(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMAPPROX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimapprox",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrimApprox(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMAXLEN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimmaxlen",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrimMaxLen(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMAXLENAPPROX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimmaxlenapprox",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XTrimMaxLenApprox(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMINID(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimminid",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrimMinID(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMINIDAPPROX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimminidapprox",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XTrimMinIDApprox(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXINFOGROUPS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfogroups",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Consumers", types.Number{}), types.NewStaticProperty("Pending", types.Number{}), types.NewStaticProperty("LastDeliveredID", types.String{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XInfoGroups(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp26843 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termConsumers := ast.IntNumberTerm(int(v.Consumers))

					termPending := ast.IntNumberTerm(int(v.Pending))

					termLastDeliveredID := ast.StringTerm(v.LastDeliveredID)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Consumers"), termConsumers}, [2]*ast.Term{ast.StringTerm("Pending"), termPending}, [2]*ast.Term{ast.StringTerm("LastDeliveredID"), termLastDeliveredID})

					tmp26843 = append(tmp26843, term)
				}
				term := ast.ArrayTerm(tmp26843...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXINFOSTREAM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfostream",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Length", types.Number{}), types.NewStaticProperty("RadixTreeKeys", types.Number{}), types.NewStaticProperty("RadixTreeNodes", types.Number{}), types.NewStaticProperty("Groups", types.Number{}), types.NewStaticProperty("LastGeneratedID", types.String{}), types.NewStaticProperty("FirstEntry", types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)), types.NewStaticProperty("LastEntry", types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XInfoStream(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termLength := ast.IntNumberTerm(int(r0.Length))

				termRadixTreeKeys := ast.IntNumberTerm(int(r0.RadixTreeKeys))

				termRadixTreeNodes := ast.IntNumberTerm(int(r0.RadixTreeNodes))

				termGroups := ast.IntNumberTerm(int(r0.Groups))

				termLastGeneratedID := ast.StringTerm(r0.LastGeneratedID)

				termFirstEntryID := ast.StringTerm(r0.FirstEntry.ID)

				var tmp26677 [][2]*ast.Term
				for key, value := range r0.FirstEntry.Values {
					k := ast.StringTerm(key)

					v := ast.NullTerm()
					if s, ok := value.(string); ok {
						v = ast.StringTerm(s)
					}

					tmp26677 = append(tmp26677, [2]*ast.Term{k, v})
				}
				termFirstEntryValues := ast.ObjectTerm(tmp26677...)

				termFirstEntry := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termFirstEntryID}, [2]*ast.Term{ast.StringTerm("Values"), termFirstEntryValues})

				termLastEntryID := ast.StringTerm(r0.LastEntry.ID)

				var tmp82198 [][2]*ast.Term
				for key, value := range r0.LastEntry.Values {
					k := ast.StringTerm(key)

					v := ast.NullTerm()
					if s, ok := value.(string); ok {
						v = ast.StringTerm(s)
					}

					tmp82198 = append(tmp82198, [2]*ast.Term{k, v})
				}
				termLastEntryValues := ast.ObjectTerm(tmp82198...)

				termLastEntry := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termLastEntryID}, [2]*ast.Term{ast.StringTerm("Values"), termLastEntryValues})

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Length"), termLength}, [2]*ast.Term{ast.StringTerm("RadixTreeKeys"), termRadixTreeKeys}, [2]*ast.Term{ast.StringTerm("RadixTreeNodes"), termRadixTreeNodes}, [2]*ast.Term{ast.StringTerm("Groups"), termGroups}, [2]*ast.Term{ast.StringTerm("LastGeneratedID"), termLastGeneratedID}, [2]*ast.Term{ast.StringTerm("FirstEntry"), termFirstEntry}, [2]*ast.Term{ast.StringTerm("LastEntry"), termLastEntry})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXINFOSTREAMFULL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfostreamfull",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Length", types.Number{}), types.NewStaticProperty("RadixTreeKeys", types.Number{}), types.NewStaticProperty("RadixTreeNodes", types.Number{}), types.NewStaticProperty("LastGeneratedID", types.String{}), types.NewStaticProperty("Entries", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))), types.NewStaticProperty("Groups", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("LastDeliveredID", types.String{}), types.NewStaticProperty("PelCount", types.Number{}), types.NewStaticProperty("Pending", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("DeliveryTime", types.Number{}), types.NewStaticProperty("DeliveryCount", types.Number{})}, nil))), types.NewStaticProperty("Consumers", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("SeenTime", types.Number{}), types.NewStaticProperty("PelCount", types.Number{}), types.NewStaticProperty("Pending", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("DeliveryTime", types.Number{}), types.NewStaticProperty("DeliveryCount", types.Number{})}, nil)))}, nil)))}, nil)))}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XInfoStreamFull(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termLength := ast.IntNumberTerm(int(r0.Length))

				termRadixTreeKeys := ast.IntNumberTerm(int(r0.RadixTreeKeys))

				termRadixTreeNodes := ast.IntNumberTerm(int(r0.RadixTreeNodes))

				termLastGeneratedID := ast.StringTerm(r0.LastGeneratedID)

				var tmp62098 []*ast.Term
				for _, v := range r0.Entries {

					termID := ast.StringTerm(v.ID)

					var tmp84155 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp84155 = append(tmp84155, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp84155...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp62098 = append(tmp62098, term)
				}
				termEntries := ast.ArrayTerm(tmp62098...)

				var tmp82792 []*ast.Term
				for _, v := range r0.Groups {

					termName := ast.StringTerm(v.Name)

					termLastDeliveredID := ast.StringTerm(v.LastDeliveredID)

					termPelCount := ast.IntNumberTerm(int(v.PelCount))

					var tmp60974 []*ast.Term
					for _, v := range v.Pending {

						termID := ast.StringTerm(v.ID)

						termConsumer := ast.StringTerm(v.Consumer)

						termDeliveryTime := ast.IntNumberTerm(int(v.DeliveryTime.UnixMicro()))

						termDeliveryCount := ast.IntNumberTerm(int(v.DeliveryCount))

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Consumer"), termConsumer}, [2]*ast.Term{ast.StringTerm("DeliveryTime"), termDeliveryTime}, [2]*ast.Term{ast.StringTerm("DeliveryCount"), termDeliveryCount})

						tmp60974 = append(tmp60974, term)
					}
					termPending := ast.ArrayTerm(tmp60974...)

					var tmp77241 []*ast.Term
					for _, v := range v.Consumers {

						termName := ast.StringTerm(v.Name)

						termSeenTime := ast.IntNumberTerm(int(v.SeenTime.UnixMicro()))

						termPelCount := ast.IntNumberTerm(int(v.PelCount))

						var tmp59332 []*ast.Term
						for _, v := range v.Pending {

							termID := ast.StringTerm(v.ID)

							termDeliveryTime := ast.IntNumberTerm(int(v.DeliveryTime.UnixMicro()))

							termDeliveryCount := ast.IntNumberTerm(int(v.DeliveryCount))

							term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("DeliveryTime"), termDeliveryTime}, [2]*ast.Term{ast.StringTerm("DeliveryCount"), termDeliveryCount})

							tmp59332 = append(tmp59332, term)
						}
						termPending := ast.ArrayTerm(tmp59332...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("SeenTime"), termSeenTime}, [2]*ast.Term{ast.StringTerm("PelCount"), termPelCount}, [2]*ast.Term{ast.StringTerm("Pending"), termPending})

						tmp77241 = append(tmp77241, term)
					}
					termConsumers := ast.ArrayTerm(tmp77241...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("LastDeliveredID"), termLastDeliveredID}, [2]*ast.Term{ast.StringTerm("PelCount"), termPelCount}, [2]*ast.Term{ast.StringTerm("Pending"), termPending}, [2]*ast.Term{ast.StringTerm("Consumers"), termConsumers})

					tmp82792 = append(tmp82792, term)
				}
				termGroups := ast.ArrayTerm(tmp82792...)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Length"), termLength}, [2]*ast.Term{ast.StringTerm("RadixTreeKeys"), termRadixTreeKeys}, [2]*ast.Term{ast.StringTerm("RadixTreeNodes"), termRadixTreeNodes}, [2]*ast.Term{ast.StringTerm("LastGeneratedID"), termLastGeneratedID}, [2]*ast.Term{ast.StringTerm("Entries"), termEntries}, [2]*ast.Term{ast.StringTerm("Groups"), termGroups})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXINFOCONSUMERS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfoconsumers",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Pending", types.Number{}), types.NewStaticProperty("Idle", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XInfoConsumers(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp47619 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termPending := ast.IntNumberTerm(int(v.Pending))

					termIdle := ast.IntNumberTerm(int(v.Idle))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Pending"), termPending}, [2]*ast.Term{ast.StringTerm("Idle"), termIdle})

					tmp47619 = append(tmp47619, term)
				}
				term := ast.ArrayTerm(tmp47619...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBZPOPMAX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bzpopmax",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{})}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BZPopMax(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termKey := ast.StringTerm(r0.Key)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Key"), termKey})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBZPOPMIN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bzpopmin",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{})}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BZPopMin(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termKey := ast.StringTerm(r0.Key)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Key"), termKey})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZADD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAdd(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDNX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddNX(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDXX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddxx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddXX(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDCH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddCh(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDNXCH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddnxch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddNXCh(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDXXCH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddxxch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddXXCh(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDARGS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddargs",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("NX", types.Boolean{}), types.NewStaticProperty("XX", types.Boolean{}), types.NewStaticProperty("LT", types.Boolean{}), types.NewStaticProperty("GT", types.Boolean{}), types.NewStaticProperty("Ch", types.Boolean{}), types.NewStaticProperty("Members", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)))}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 redis.ZAddArgs
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddArgs(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDARGSINCR(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddargsincr",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("NX", types.Boolean{}), types.NewStaticProperty("XX", types.Boolean{}), types.NewStaticProperty("LT", types.Boolean{}), types.NewStaticProperty("GT", types.Boolean{}), types.NewStaticProperty("Ch", types.Boolean{}), types.NewStaticProperty("Members", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)))}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 redis.ZAddArgs
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddArgsIncr(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCR(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincr",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZIncr(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCRNX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincrnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZIncrNX(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCRXX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincrxx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZIncrXX(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZCARD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zcard",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZCard(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZCount(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZLEXCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zlexcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZLexCount(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCRBY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZIncrBy(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINTER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zinter",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZInter(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp63541 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp63541 = append(tmp63541, term)
				}
				term := ast.ArrayTerm(tmp63541...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZINTERWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zinterwithscores",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZInterWithScores(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp05773 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp05773 = append(tmp05773, term)
				}
				term := ast.ArrayTerm(tmp05773...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZINTERSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zinterstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZStore
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZInterStore(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZMSCORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zmscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Number{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZMScore(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp38877 []*ast.Term
				for _, v := range r0 {
					term := ast.FloatNumberTerm(float64(v))
					tmp38877 = append(tmp38877, term)
				}
				term := ast.ArrayTerm(tmp38877...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZPOPMAX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zpopmax",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Number{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZPopMax(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp96413 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp96413 = append(tmp96413, term)
				}
				term := ast.ArrayTerm(tmp96413...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZPOPMIN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zpopmin",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Number{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZPopMin(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp09279 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp09279 = append(tmp09279, term)
				}
				term := ast.ArrayTerm(tmp09279...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRange(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp35322 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp35322 = append(tmp35322, term)
				}
				term := ast.ArrayTerm(tmp35322...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRangeWithScores(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp73783 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp73783 = append(tmp73783, term)
				}
				term := ast.ArrayTerm(tmp73783...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEBYSCORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangebyscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeByScore(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp87102 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp87102 = append(tmp87102, term)
				}
				term := ast.ArrayTerm(tmp87102...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEBYLEX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangebylex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeByLex(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp47641 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp47641 = append(tmp47641, term)
				}
				term := ast.ArrayTerm(tmp47641...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEBYSCOREWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangebyscorewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeByScoreWithScores(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp98721 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp98721 = append(tmp98721, term)
				}
				term := ast.ArrayTerm(tmp98721...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEARGS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangeargs",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}), types.NewStaticProperty("Start", types.Any{}), types.NewStaticProperty("Stop", types.Any{}), types.NewStaticProperty("ByScore", types.Boolean{}), types.NewStaticProperty("ByLex", types.Boolean{}), types.NewStaticProperty("Rev", types.Boolean{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZRangeArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZRangeArgs(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp73464 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp73464 = append(tmp73464, term)
				}
				term := ast.ArrayTerm(tmp73464...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEARGSWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangeargswithscores",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}), types.NewStaticProperty("Start", types.Any{}), types.NewStaticProperty("Stop", types.Any{}), types.NewStaticProperty("ByScore", types.Boolean{}), types.NewStaticProperty("ByLex", types.Boolean{}), types.NewStaticProperty("Rev", types.Boolean{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZRangeArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZRangeArgsWithScores(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp48747 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp48747 = append(tmp48747, term)
				}
				term := ast.ArrayTerm(tmp48747...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGESTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangestore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}), types.NewStaticProperty("Start", types.Any{}), types.NewStaticProperty("Stop", types.Any{}), types.NewStaticProperty("ByScore", types.Boolean{}), types.NewStaticProperty("ByLex", types.Boolean{}), types.NewStaticProperty("Rev", types.Boolean{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 redis.ZRangeArgs
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeStore(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZRANK(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrank",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRank(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREM(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrem",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRem(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREMRANGEBYRANK(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zremrangebyrank",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRemRangeByRank(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREMRANGEBYSCORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zremrangebyscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRemRangeByScore(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREMRANGEBYLEX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zremrangebylex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRemRangeByLex(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRevRange(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp88679 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp88679 = append(tmp88679, term)
				}
				term := ast.ArrayTerm(tmp88679...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeWithScores(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp77982 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp77982 = append(tmp77982, term)
				}
				term := ast.ArrayTerm(tmp77982...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEBYSCORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangebyscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeByScore(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp05485 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp05485 = append(tmp05485, term)
				}
				term := ast.ArrayTerm(tmp05485...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEBYLEX(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangebylex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeByLex(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp16561 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp16561 = append(tmp16561, term)
				}
				term := ast.ArrayTerm(tmp16561...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEBYSCOREWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangebyscorewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeByScoreWithScores(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp72482 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp72482 = append(tmp72482, term)
				}
				term := ast.ArrayTerm(tmp72482...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANK(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrank",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRank(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZSCORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZScore(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZUNIONSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zunionstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZStore
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZUnionStore(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZUNION(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zunion",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZUnion(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp51014 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp51014 = append(tmp51014, term)
				}
				term := ast.ArrayTerm(tmp51014...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZUNIONWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zunionwithscores",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZUnionWithScores(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp69586 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp69586 = append(tmp69586, term)
				}
				term := ast.ArrayTerm(tmp69586...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANDMEMBER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrandmember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Boolean{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 bool
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRandMember(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp69802 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp69802 = append(tmp69802, term)
				}
				term := ast.ArrayTerm(tmp69802...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZDIFF(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zdiff",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZDiff(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp76712 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp76712 = append(tmp76712, term)
				}
				term := ast.ArrayTerm(tmp76712...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZDIFFWITHSCORES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zdiffwithscores",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZDiffWithScores(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp91312 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp91312 = append(tmp91312, term)
				}
				term := ast.ArrayTerm(tmp91312...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZDIFFSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zdiffstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZDiffStore(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPFADD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pfadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PFAdd(m.RedisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPFCOUNT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pfcount",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PFCount(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPFMERGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pfmerge",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PFMerge(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBGREWRITEAOF(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bgrewriteaof",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.BgRewriteAOF(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBGSAVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bgsave",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.BgSave(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTKILL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientkill",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClientKill(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTKILLBYFILTER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientkillbyfilter",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClientKillByFilter(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTLIST(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientlist",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClientList(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTPAUSE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientpause",
			Decl:             types.NewFunction(types.Args(types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClientPause(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTID(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientid",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClientID(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGGET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configget",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ConfigGet(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp04047 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp04047 = append(tmp04047, term)
				}
				term := ast.ArrayTerm(tmp04047...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGRESETSTAT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configresetstat",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ConfigResetStat(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGSET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ConfigSet(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGREWRITE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configrewrite",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ConfigRewrite(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDBSIZE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.dbsize",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.DBSize(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHALL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushall",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushAll(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHALLASYNC(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushallasync",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushAllAsync(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHDB(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushdb",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushDB(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHDBASYNC(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushdbasync",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushDBAsync(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINFO(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.info",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Info(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLASTSAVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lastsave",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.LastSave(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSAVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.save",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Save(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSHUTDOWN(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.shutdown",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Shutdown(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSHUTDOWNSAVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.shutdownsave",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ShutdownSave(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSHUTDOWNNOSAVE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.shutdownnosave",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ShutdownNoSave(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSLAVEOF(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.slaveof",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SlaveOf(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerTIME(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.time",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Time(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.UnixMicro()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDEBUGOBJECT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.debugobject",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.DebugObject(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerREADONLY(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.readonly",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ReadOnly(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerREADWRITE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.readwrite",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ReadWrite(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMEMORYUSAGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.memoryusage",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Number{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.MemoryUsage(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEVAL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.eval",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{}), types.NewArray([]types.Type{}, types.Any{})), types.Any{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Eval(m.RedisContext, v0, v1, utils.Conva(v2)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.NullTerm()
				if s, ok := r0.(string); ok {
					term = ast.StringTerm(s)
				}

				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEVALSHA(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.evalsha",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{}), types.NewArray([]types.Type{}, types.Any{})), types.Any{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.EvalSha(m.RedisContext, v0, v1, utils.Conva(v2)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.NullTerm()
				if s, ok := r0.(string); ok {
					term = ast.StringTerm(s)
				}

				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTEXISTS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptexists",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Boolean{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ScriptExists(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp28016 []*ast.Term
				for _, v := range r0 {
					term := ast.BooleanTerm(v)
					tmp28016 = append(tmp28016, term)
				}
				term := ast.ArrayTerm(tmp28016...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTFLUSH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptflush",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ScriptFlush(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTKILL(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptkill",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ScriptKill(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTLOAD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptload",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ScriptLoad(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPUBLISH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.publish",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Publish(m.RedisContext, v0, utils.Conv(v1))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPUBSUBCHANNELS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pubsubchannels",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PubSubChannels(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp68232 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp68232 = append(tmp68232, term)
				}
				term := ast.ArrayTerm(tmp68232...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerPUBSUBNUMSUB(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pubsubnumsub",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Number{}))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PubSubNumSub(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp91137 [][2]*ast.Term
				for key, value := range r0 {
					k := ast.StringTerm(key)
					v := ast.IntNumberTerm(int(value))
					tmp91137 = append(tmp91137, [2]*ast.Term{k, v})
				}
				term := ast.ObjectTerm(tmp91137...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerPUBSUBNUMPAT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pubsubnumpat",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.PubSubNumPat(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERSLOTS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterslots",
			Decl:             types.NewFunction(types.Args(), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Start", types.Number{}), types.NewStaticProperty("End", types.Number{}), types.NewStaticProperty("Nodes", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Addr", types.String{})}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterSlots(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp67290 []*ast.Term
				for _, v := range r0 {

					termStart := ast.IntNumberTerm(int(v.Start))

					termEnd := ast.IntNumberTerm(int(v.End))

					var tmp85458 []*ast.Term
					for _, v := range v.Nodes {

						termID := ast.StringTerm(v.ID)

						termAddr := ast.StringTerm(v.Addr)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Addr"), termAddr})

						tmp85458 = append(tmp85458, term)
					}
					termNodes := ast.ArrayTerm(tmp85458...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Start"), termStart}, [2]*ast.Term{ast.StringTerm("End"), termEnd}, [2]*ast.Term{ast.StringTerm("Nodes"), termNodes})

					tmp67290 = append(tmp67290, term)
				}
				term := ast.ArrayTerm(tmp67290...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERNODES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusternodes",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterNodes(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERMEET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustermeet",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterMeet(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERFORGET(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterforget",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterForget(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERREPLICATE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterreplicate",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterReplicate(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERRESETSOFT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterresetsoft",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterResetSoft(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERRESETHARD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterresethard",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterResetHard(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERINFO(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterinfo",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterInfo(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERKEYSLOT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterkeyslot",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterKeySlot(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERGETKEYSINSLOT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustergetkeysinslot",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterGetKeysInSlot(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp25918 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp25918 = append(tmp25918, term)
				}
				term := ast.ArrayTerm(tmp25918...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERCOUNTFAILUREREPORTS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustercountfailurereports",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterCountFailureReports(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERCOUNTKEYSINSLOT(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustercountkeysinslot",
			Decl:             types.NewFunction(types.Args(types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterCountKeysInSlot(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERDELSLOTS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterdelslots",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Number{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterDelSlots(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERDELSLOTSRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterdelslotsrange",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterDelSlotsRange(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERSAVECONFIG(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustersaveconfig",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterSaveConfig(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERSLAVES(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterslaves",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterSlaves(m.RedisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp50236 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp50236 = append(tmp50236, term)
				}
				term := ast.ArrayTerm(tmp50236...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERFAILOVER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterfailover",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterFailover(m.RedisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERADDSLOTS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusteraddslots",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Number{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterAddSlots(m.RedisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERADDSLOTSRANGE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusteraddslotsrange",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterAddSlotsRange(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOADD(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geoadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.GeoLocation
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoAdd(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOPOS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geopos",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoPos(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp96520 []*ast.Term
				for _, v := range r0 {
					if v == nil {
						tmp96520 = append(tmp96520, ast.NullTerm())
						continue
					}

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude})

					tmp96520 = append(tmp96520, term)
				}
				term := ast.ArrayTerm(tmp96520...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUS(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadius",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 float64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 *redis.GeoRadiusQuery
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.GeoRadius(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp90486 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					termDist := ast.FloatNumberTerm(float64(v.Dist))

					termGeoHash := ast.IntNumberTerm(int(v.GeoHash))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude}, [2]*ast.Term{ast.StringTerm("Dist"), termDist}, [2]*ast.Term{ast.StringTerm("GeoHash"), termGeoHash})

					tmp90486 = append(tmp90486, term)
				}
				term := ast.ArrayTerm(tmp90486...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUSSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadiusstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 float64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 *redis.GeoRadiusQuery
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.GeoRadiusStore(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUSBYMEMBER(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadiusbymember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.GeoRadiusQuery
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GeoRadiusByMember(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp23294 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					termDist := ast.FloatNumberTerm(float64(v.Dist))

					termGeoHash := ast.IntNumberTerm(int(v.GeoHash))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude}, [2]*ast.Term{ast.StringTerm("Dist"), termDist}, [2]*ast.Term{ast.StringTerm("GeoHash"), termGeoHash})

					tmp23294 = append(tmp23294, term)
				}
				term := ast.ArrayTerm(tmp23294...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUSBYMEMBERSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadiusbymemberstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.GeoRadiusQuery
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GeoRadiusByMemberStore(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOSEARCH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geosearch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Member", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("RadiusUnit", types.String{}), types.NewStaticProperty("BoxWidth", types.Number{}), types.NewStaticProperty("BoxHeight", types.Number{}), types.NewStaticProperty("BoxUnit", types.String{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("CountAny", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.GeoSearchQuery
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoSearch(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp68977 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp68977 = append(tmp68977, term)
				}
				term := ast.ArrayTerm(tmp68977...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEOSEARCHLOCATION(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geosearchlocation",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithHash", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.GeoSearchLocationQuery
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoSearchLocation(m.RedisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp92884 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					termDist := ast.FloatNumberTerm(float64(v.Dist))

					termGeoHash := ast.IntNumberTerm(int(v.GeoHash))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude}, [2]*ast.Term{ast.StringTerm("Dist"), termDist}, [2]*ast.Term{ast.StringTerm("GeoHash"), termGeoHash})

					tmp92884 = append(tmp92884, term)
				}
				term := ast.ArrayTerm(tmp92884...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEOSEARCHSTORE(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geosearchstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("StoreDist", types.Boolean{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.GeoSearchStoreQuery
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GeoSearchStore(m.RedisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEODIST(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geodist",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.GeoDist(m.RedisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOHASH(m *RedisManager) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geohash",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := m.RedisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoHash(m.RedisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp07073 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp07073 = append(tmp07073, term)
				}
				term := ast.ArrayTerm(tmp07073...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func (m *RedisManager) registerAutogenCommands() {
	registerCLIENTGETNAME(m)
	registerECHO(m)
	registerPING(m)
	registerQUIT(m)
	registerDEL(m)
	registerUNLINK(m)
	registerDUMP(m)
	registerEXISTS(m)
	registerEXPIRE(m)
	registerEXPIREAT(m)
	registerEXPIRENX(m)
	registerEXPIREXX(m)
	registerEXPIREGT(m)
	registerEXPIRELT(m)
	registerKEYS(m)
	registerMIGRATE(m)
	registerMOVE(m)
	registerOBJECTREFCOUNT(m)
	registerOBJECTENCODING(m)
	registerOBJECTIDLETIME(m)
	registerPERSIST(m)
	registerPEXPIRE(m)
	registerPEXPIREAT(m)
	registerPTTL(m)
	registerRANDOMKEY(m)
	registerRENAME(m)
	registerRENAMENX(m)
	registerRESTORE(m)
	registerRESTOREREPLACE(m)
	registerSORT(m)
	registerSORTSTORE(m)
	registerSORTINTERFACES(m)
	registerTOUCH(m)
	registerTTL(m)
	registerTYPE(m)
	registerAPPEND(m)
	registerDECR(m)
	registerDECRBY(m)
	registerGET(m)
	registerGETRANGE(m)
	registerGETSET(m)
	registerGETEX(m)
	registerGETDEL(m)
	registerINCR(m)
	registerINCRBY(m)
	registerINCRBYFLOAT(m)
	registerMGET(m)
	registerMSET(m)
	registerMSETNX(m)
	registerSET(m)
	registerSETARGS(m)
	registerSETEX(m)
	registerSETNX(m)
	registerSETXX(m)
	registerSETRANGE(m)
	registerSTRLEN(m)
	registerCOPY(m)
	registerGETBIT(m)
	registerSETBIT(m)
	registerBITCOUNT(m)
	registerBITOPAND(m)
	registerBITOPOR(m)
	registerBITOPXOR(m)
	registerBITOPNOT(m)
	registerBITPOS(m)
	registerBITFIELD(m)
	registerSCAN(m)
	registerSCANTYPE(m)
	registerSSCAN(m)
	registerHSCAN(m)
	registerZSCAN(m)
	registerHDEL(m)
	registerHEXISTS(m)
	registerHGET(m)
	registerHGETALL(m)
	registerHINCRBY(m)
	registerHINCRBYFLOAT(m)
	registerHKEYS(m)
	registerHLEN(m)
	registerHMGET(m)
	registerHSET(m)
	registerHMSET(m)
	registerHSETNX(m)
	registerHVALS(m)
	registerHRANDFIELD(m)
	registerBLPOP(m)
	registerBRPOP(m)
	registerBRPOPLPUSH(m)
	registerLINDEX(m)
	registerLINSERT(m)
	registerLINSERTBEFORE(m)
	registerLINSERTAFTER(m)
	registerLLEN(m)
	registerLPOP(m)
	registerLPOPCOUNT(m)
	registerLPOS(m)
	registerLPOSCOUNT(m)
	registerLPUSH(m)
	registerLPUSHX(m)
	registerLRANGE(m)
	registerLREM(m)
	registerLSET(m)
	registerLTRIM(m)
	registerRPOP(m)
	registerRPOPCOUNT(m)
	registerRPOPLPUSH(m)
	registerRPUSH(m)
	registerRPUSHX(m)
	registerLMOVE(m)
	registerBLMOVE(m)
	registerSADD(m)
	registerSCARD(m)
	registerSDIFF(m)
	registerSDIFFSTORE(m)
	registerSINTER(m)
	registerSINTERSTORE(m)
	registerSISMEMBER(m)
	registerSMISMEMBER(m)
	registerSMEMBERS(m)
	registerSMEMBERSMAP(m)
	registerSMOVE(m)
	registerSPOP(m)
	registerSPOPN(m)
	registerSRANDMEMBER(m)
	registerSRANDMEMBERN(m)
	registerSREM(m)
	registerSUNION(m)
	registerSUNIONSTORE(m)
	registerXADD(m)
	registerXDEL(m)
	registerXLEN(m)
	registerXRANGE(m)
	registerXRANGEN(m)
	registerXREVRANGE(m)
	registerXREVRANGEN(m)
	registerXREAD(m)
	registerXREADSTREAMS(m)
	registerXGROUPCREATE(m)
	registerXGROUPCREATEMKSTREAM(m)
	registerXGROUPSETID(m)
	registerXGROUPDESTROY(m)
	registerXGROUPCREATECONSUMER(m)
	registerXGROUPDELCONSUMER(m)
	registerXREADGROUP(m)
	registerXACK(m)
	registerXPENDING(m)
	registerXPENDINGEXT(m)
	registerXCLAIM(m)
	registerXCLAIMJUSTID(m)
	registerXAUTOCLAIM(m)
	registerXAUTOCLAIMJUSTID(m)
	registerXTRIM(m)
	registerXTRIMAPPROX(m)
	registerXTRIMMAXLEN(m)
	registerXTRIMMAXLENAPPROX(m)
	registerXTRIMMINID(m)
	registerXTRIMMINIDAPPROX(m)
	registerXINFOGROUPS(m)
	registerXINFOSTREAM(m)
	registerXINFOSTREAMFULL(m)
	registerXINFOCONSUMERS(m)
	registerBZPOPMAX(m)
	registerBZPOPMIN(m)
	registerZADD(m)
	registerZADDNX(m)
	registerZADDXX(m)
	registerZADDCH(m)
	registerZADDNXCH(m)
	registerZADDXXCH(m)
	registerZADDARGS(m)
	registerZADDARGSINCR(m)
	registerZINCR(m)
	registerZINCRNX(m)
	registerZINCRXX(m)
	registerZCARD(m)
	registerZCOUNT(m)
	registerZLEXCOUNT(m)
	registerZINCRBY(m)
	registerZINTER(m)
	registerZINTERWITHSCORES(m)
	registerZINTERSTORE(m)
	registerZMSCORE(m)
	registerZPOPMAX(m)
	registerZPOPMIN(m)
	registerZRANGE(m)
	registerZRANGEWITHSCORES(m)
	registerZRANGEBYSCORE(m)
	registerZRANGEBYLEX(m)
	registerZRANGEBYSCOREWITHSCORES(m)
	registerZRANGEARGS(m)
	registerZRANGEARGSWITHSCORES(m)
	registerZRANGESTORE(m)
	registerZRANK(m)
	registerZREM(m)
	registerZREMRANGEBYRANK(m)
	registerZREMRANGEBYSCORE(m)
	registerZREMRANGEBYLEX(m)
	registerZREVRANGE(m)
	registerZREVRANGEWITHSCORES(m)
	registerZREVRANGEBYSCORE(m)
	registerZREVRANGEBYLEX(m)
	registerZREVRANGEBYSCOREWITHSCORES(m)
	registerZREVRANK(m)
	registerZSCORE(m)
	registerZUNIONSTORE(m)
	registerZUNION(m)
	registerZUNIONWITHSCORES(m)
	registerZRANDMEMBER(m)
	registerZDIFF(m)
	registerZDIFFWITHSCORES(m)
	registerZDIFFSTORE(m)
	registerPFADD(m)
	registerPFCOUNT(m)
	registerPFMERGE(m)
	registerBGREWRITEAOF(m)
	registerBGSAVE(m)
	registerCLIENTKILL(m)
	registerCLIENTKILLBYFILTER(m)
	registerCLIENTLIST(m)
	registerCLIENTPAUSE(m)
	registerCLIENTID(m)
	registerCONFIGGET(m)
	registerCONFIGRESETSTAT(m)
	registerCONFIGSET(m)
	registerCONFIGREWRITE(m)
	registerDBSIZE(m)
	registerFLUSHALL(m)
	registerFLUSHALLASYNC(m)
	registerFLUSHDB(m)
	registerFLUSHDBASYNC(m)
	registerINFO(m)
	registerLASTSAVE(m)
	registerSAVE(m)
	registerSHUTDOWN(m)
	registerSHUTDOWNSAVE(m)
	registerSHUTDOWNNOSAVE(m)
	registerSLAVEOF(m)
	registerTIME(m)
	registerDEBUGOBJECT(m)
	registerREADONLY(m)
	registerREADWRITE(m)
	registerMEMORYUSAGE(m)
	registerEVAL(m)
	registerEVALSHA(m)
	registerSCRIPTEXISTS(m)
	registerSCRIPTFLUSH(m)
	registerSCRIPTKILL(m)
	registerSCRIPTLOAD(m)
	registerPUBLISH(m)
	registerPUBSUBCHANNELS(m)
	registerPUBSUBNUMSUB(m)
	registerPUBSUBNUMPAT(m)
	registerCLUSTERSLOTS(m)
	registerCLUSTERNODES(m)
	registerCLUSTERMEET(m)
	registerCLUSTERFORGET(m)
	registerCLUSTERREPLICATE(m)
	registerCLUSTERRESETSOFT(m)
	registerCLUSTERRESETHARD(m)
	registerCLUSTERINFO(m)
	registerCLUSTERKEYSLOT(m)
	registerCLUSTERGETKEYSINSLOT(m)
	registerCLUSTERCOUNTFAILUREREPORTS(m)
	registerCLUSTERCOUNTKEYSINSLOT(m)
	registerCLUSTERDELSLOTS(m)
	registerCLUSTERDELSLOTSRANGE(m)
	registerCLUSTERSAVECONFIG(m)
	registerCLUSTERSLAVES(m)
	registerCLUSTERFAILOVER(m)
	registerCLUSTERADDSLOTS(m)
	registerCLUSTERADDSLOTSRANGE(m)
	registerGEOADD(m)
	registerGEOPOS(m)
	registerGEORADIUS(m)
	registerGEORADIUSSTORE(m)
	registerGEORADIUSBYMEMBER(m)
	registerGEORADIUSBYMEMBERSTORE(m)
	registerGEOSEARCH(m)
	registerGEOSEARCHLOCATION(m)
	registerGEOSEARCHSTORE(m)
	registerGEODIST(m)
	registerGEOHASH(m)
}
