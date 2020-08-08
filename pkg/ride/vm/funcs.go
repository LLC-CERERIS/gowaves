package vm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/errs"
	"github.com/wavesplatform/gowaves/pkg/proto"
	"github.com/wavesplatform/gowaves/pkg/ride/evaluator/ast"
)

//func withLong(expr ast.Expr, f func(l int64) error) error {

//}

//type With struct {
//}
//
//func (a *With) Long(f func(long int64, w *With) error) error {
//
//}

// expects f looks like `func(x int64) error`
func with(s Context, f interface{}) error {
	v := reflect.ValueOf(f)
	x := v.Type()

	if x.NumOut() != 1 {
		return errors.Errorf("expected passed function returns exactly 1 arguument, passed %d, %s", x.NumOut(), x.Name())
	}
	args := make([]reflect.Value, x.NumIn())
	for i := x.NumIn() - 1; i >= 0; i-- {
		inV := x.In(i)
		value := s.Pop()
		if value == nil {
			return errors.Errorf("with: empty stack")
		}
		// TODO check outher types
		switch inV.Kind() {
		case reflect.Int64:
			args[i] = reflect.ValueOf(value.(*ast.LongExpr).Value)
		case reflect.Bool:
			args[i] = reflect.ValueOf(value.(*ast.BooleanExpr).Value)
		case reflect.String:
			v, ok := value.(*ast.StringExpr)
			if !ok {
				return errors.Errorf("with: expected '%d' argument to be *ast.StringExpr, found %T", i, value)
			}
			args[i] = reflect.ValueOf(v.Value)
		case reflect.Interface:
			args[i] = reflect.ValueOf(value)
		case reflect.Slice: // []byte
			args[i] = reflect.ValueOf(value.(*ast.BytesExpr).Value)
		}
	}
	resp := v.Call(args)[0]
	if err, ok := resp.Interface().(error); ok {
		return err
	}
	return nil
}

//func GteLong(s Context) error {
//	second := s.Pop().(*ast.LongExpr).Value
//	first := s.Pop().(*ast.LongExpr).Value
//	s.Push(ast.NewBoolean(first >= second))
//	return nil
//}

func GteLong(s Context) error {
	//return with(s).Long(func(second int64, w *With) error {
	//	return w.Long(func(first int64, w *With) error {
	//		s.Push(ast.NewBoolean(first >= second))
	//		return nil
	//	})
	//})
	//
	//w := with(s)
	//w.Second().AsLong()

	//first := s.Pop().L
	//s.Push(ast.NewBoolean(first >= second))
	//return nil
	return with(s, func(first int64, second int64) error {
		s.Push(ast.NewBoolean(first >= second))
		return nil
	})
}

func Eq(s Context) error {
	second := s.Pop()
	first := s.Pop()

	//if first.IsObject() {
	//	s.Push(B(first.O.Eq(second.O)))
	//	return nil
	//}
	s.Push(ast.NewBoolean(first.Eq(second)))
	return nil
}

func GetterFn(s Context) error {
	second := s.Pop().(*ast.StringExpr)
	first := s.Pop()
	g, ok := first.(ast.Getable)
	if !ok {
		return errors.Errorf("GetterFn: expected first argument to be ast.Getable, found %T", first)
	}
	expr, err := g.Get(second.Value)
	if err != nil {
		return err
	}
	s.Push(expr)
	return nil
}

func Neq(s Context) error {
	err := Eq(s)
	if err != nil {
		return err
	}
	first := s.Pop().(*ast.BooleanExpr)
	s.Push(ast.NewBoolean(!first.Value))
	return nil
}

func IsInstanceOf(s Context) error {
	return with(s, func(first ast.Expr, second string) error {
		s.Push(ast.NewBoolean(first.InstanceOf() == second))
		return nil
	})
}

// Size of list
func NativeSizeList(s Context) error {
	const funcName = "NativeSizeList"
	//if l := len(e); l != 1 {
	//	return nil, errors.Errorf("%s: invalid params, expected 1, passed %d", funcName, l)
	//}
	e := s.Pop()
	// optimize not evaluate inner list
	if v, ok := e.(ast.Exprs); ok {
		s.Push(ast.NewLong(int64(len(v))))
		return nil
	}
	return errors.Errorf("%s: expected first argument to be ast.Expr, got %T", funcName, e)
	//rs, err := e[0].Evaluate(s)
	//if err != nil {
	//	return nil, errors.Wrap(err, funcName)
	//}
	//lst, ok := rs.(Exprs)
	//if !ok {
	//	return nil, errors.Errorf("%s: expected first argument Exprs, got %T", funcName, rs)
	//}
	//return NewLong(int64(len(lst))), nil
}

// type constructor
func UserAddress(s Context) error {
	//const funcName = "UserAddress"
	return with(s, func(value []byte) error {
		addr, err := proto.NewAddressFromBytes(value)
		if err != nil {
			s.Push(&ast.InvalidAddressExpr{Value: value})
			return nil
		}
		s.Push(ast.NewAddressFromProtoAddress(addr))
		return nil
	})
	//if l := len(e); l != 1 {
	//	return nil, errors.Errorf("%s: invalid params, expected 1, passed %d", funcName, l)
	//}
	//first = s.Pop()
	////if err != nil {
	////	return nil, errors.Wrap(err, funcName)
	////}
	//bts, ok := first.(*BytesExpr)
	//if !ok {
	//	return nil, errors.Errorf("%s: first argument expected to be *BytesExpr, found %T", funcName, first)
	//}
	//addr, err := proto.NewAddressFromBytes(bts.Value)
	//if err != nil {
	//	return &ast.InvalidAddressExpr{Value: bts.Value}, nil
	//}
	//return NewAddressFromProtoAddress(addr), nil
}

//// Decode account address
//func UserAddressFromString(s Context) error {
//	str := s.Pop().(*ast.StringExpr).Value
//	//proto.NewAddressFromString(str)
//
//	//rs, err := e[0].Evaluate(s)
//	//if err != nil {
//	//	return nil, errors.Wrap(err, "UserAddressFromString")
//	//}
//	//str, ok := rs.(*StringExpr)
//	//if !ok {
//	//	return nil, errors.Errorf("UserAddressFromString: expected first argument to be *StringExpr, found %T", rs)
//	//}
//	addr, err := ast.NewAddressFromString(str)
//	if err != nil {
//		s.Push(ast.NewUnit())
//		return nil
//	}
//	// TODO return this back
//	//if addr[1] != s.Scheme() {
//	//	return NewUnit(), nil
//	//}
//	s.Push(addr)
//	return nil
//}

func NativeCreateList(s Context) error {
	const funcName = "NativeCreateList"
	second := s.Pop()
	head := s.Pop()
	//if l := len(e); l != 2 {
	//	return nil, errors.Errorf("%s: invalid parameters, expected 2, received %d", funcName, l)
	//}
	//head, err := e[0].Evaluate(s)
	//if err != nil {
	//	return nil, errors.Wrap(err, funcName)
	//}
	//t, err := e[1].Evaluate(s)
	//if err != nil {
	//	return nil, errors.Wrap(err, funcName)
	//}
	tail, ok := second.(ast.Exprs)
	if !ok {
		return errors.Errorf("%s: invalid second parameter, expected Exprs, received %T", funcName, second)
	}
	if len(tail) == 0 {
		s.Push(ast.NewExprs(head))
		//return NewExprs(head), nil
		return nil
	}
	//return append(ast.NewExprs(head), tail...), nil
	s.Push(append(ast.NewExprs(head), tail...))
	return nil
}

// Get list element by position
func NativeGetList(s Context) error {
	const funcName = "NativeGetList"
	//if l := len(e); l != 2 {
	//	return nil, errors.Errorf("%s: invalid params, expected 2, passed %d", funcName, l)
	//}
	second := s.Pop()
	first := s.Pop()

	lst, ok := first.(ast.Exprs)
	if !ok {
		return errors.Errorf("%s: expected first argument Exprs, got %T", funcName, first)
	}
	lng, ok := second.(*ast.LongExpr)
	if !ok {
		return errors.Errorf("%s: expected second argument *LongExpr, got %T", funcName, second)
	}
	if lng.Value < 0 || lng.Value >= int64(len(lst)) {
		return errors.Errorf("%s: invalid index %d, len %d", funcName, lng.Value, len(lst))
	}
	s.Push(lst[lng.Value])
	return nil
}

// Fail script
func NativeThrow(s Context) error {
	const funcName = "NativeThrow"
	first := s.Pop()
	//if l := len(e); l != 1 {
	//	return nil, errors.Errorf("%s: invalid params, expected 1, passed %d", funcName, l)
	//}
	//first, err := e[0].Evaluate(s)
	//if err != nil {
	//	return nil, errors.Wrap(err, funcName)
	//}
	str, ok := first.(*ast.StringExpr)
	if !ok {
		return errors.Errorf("%s: expected first argument to be *StringExpr, found %T", funcName, first)
	}
	return &ast.Throw{
		Message: str.Value,
	}
}

func SigVerifyV2(s Context) error {
	return with(s, func(data []byte, sigBytes []byte, publicKey []byte) error {
		//if l := len(data); !s.validMessageLength(l) || limit > 0 && l > limit*1024 {
		//	return errors.Errorf("%s: invalid message size %d", fn, l)
		//}

		signature, err := crypto.NewSignatureFromBytes(sigBytes)
		if err != nil {
			return errs.Extendf(err, "bytes len: %d", len(sigBytes))
		}
		pk, err := crypto.NewPublicKeyFromBytes(publicKey)
		if err != nil {
			return err
		}
		out := crypto.Verify(pk, signature, data)
		s.Push(ast.NewBoolean(out))
		return nil
	})
}

//func limitedSigVerify(limit int) Func {
//	fn := "SigVerify"
//	if limit > 0 {
//		fn = fmt.Sprintf("%s_%dKb", fn, limit)
//	}
//	return func(s Context) error {
//		return with(s, func(data []byte, sigBytes []byte, publicKey []byte) error {
//			if l := len(data); !s.validMessageLength(l) || limit > 0 && l > limit*1024 {
//				return errors.Errorf("%s: invalid message size %d", fn, l)
//			}
//
//			signature, err := crypto.NewSignatureFromBytes(sigBytes)
//			if err != nil {
//				return err
//			}
//			pk, err := crypto.NewPublicKeyFromBytes(publicKey)
//			if err != nil {
//				return err
//			}
//			out := crypto.Verify(pk, signature, data)
//			s.Push(ast.NewBoolean(out))
//			return nil
//		})
//		//if l := len(e); l != 3 {
//		//	return nil, errors.Errorf("%s: invalid number of parameters %d, expected 3", fn, l)
//		//}
//		//rs, err := e.EvaluateAll(s)
//		//if err != nil {
//		//	return nil, errors.Wrap(err, fn)
//		//}
//		/*
//			messageExpr, ok := rs[0].(*BytesExpr)
//			if !ok {
//				return nil, errors.Errorf("%s: first argument expects to be *BytesExpr, found %T", fn, rs[0])
//			}
//			if l := len(messageExpr.Value); !s.validMessageLength(l) || limit > 0 && l > limit*1024 {
//				return nil, errors.Errorf("%s: invalid message size %d", fn, l)
//			}
//			signatureExpr, ok := rs[1].(*BytesExpr)
//			if !ok {
//				return nil, errors.Errorf("%s: second argument expects to be *BytesExpr, found %T", fn, rs[1])
//			}
//			pkExpr, ok := rs[2].(*BytesExpr)
//			if !ok {
//				return nil, errors.Errorf("%s: third argument expects to be *BytesExpr, found %T", fn, rs[2])
//			}
//			pk, err := crypto.NewPublicKeyFromBytes(pkExpr.Value)
//			if err != nil {
//				return NewBoolean(false), nil
//			}
//			signature, err := crypto.NewSignatureFromBytes(signatureExpr.Value)
//			if err != nil {
//				return NewBoolean(false), nil
//			}
//			out := crypto.Verify(pk, signature, messageExpr.Value)
//			return NewBoolean(out), nil
//		*/
//	}
//}

func extractRecipient(e ast.Expr) (proto.Recipient, error) {
	var r proto.Recipient
	switch a := e.(type) {
	case *ast.AddressExpr:
		r = proto.NewRecipientFromAddress(proto.Address(*a))
	case *ast.AliasExpr:
		r = proto.NewRecipientFromAlias(proto.Alias(*a))
	case *ast.RecipientExpr:
		r = proto.Recipient(*a)
	default:
		return proto.Recipient{}, errors.Errorf("expected to be AddressExpr or AliasExpr, found %T", e)
	}
	return r, nil
}

func extractRecipientAndKey(s Context) (proto.Recipient, string, error) {
	//if l := len(e); l != 2 {
	//	return proto.Recipient{}, "", errors.Errorf("invalid params, expected 2, passed %d", l)
	//}

	second := s.Pop()
	first := s.Pop()
	//if err != nil {
	//	return proto.Recipient{}, "", err
	//}
	r, err := extractRecipient(first)
	if err != nil {
		return proto.Recipient{}, "", errors.Errorf("first argument %v", err)
	}

	//if err != nil {
	//	return proto.Recipient{}, "", err
	//}
	key, ok := second.(*ast.StringExpr)
	if !ok {
		return proto.Recipient{}, "", errors.Errorf("second argument expected to be *StringExpr, found %T", second)
	}
	return r, key.Value, nil
}

// Get integer from account state
func NativeDataIntegerFromState(s Context) error {
	r, k, err := extractRecipientAndKey(s)
	if err != nil {
		return s.Push(ast.NewUnit())
	}
	entry, err := s.State().RetrieveNewestIntegerEntry(r, k)
	if err != nil {
		return s.Push(ast.NewUnit())
	}
	return s.Push(ast.NewLong(entry.Value))
}

// Decode account address
func UserAddressFromString(s Context) error {
	return with(s, func(value string) error {
		addr, err := ast.NewAddressFromString(value)
		if err != nil {
			return s.Push(ast.NewUnit())
		}
		if addr[1] != s.Scheme() {
			return s.Push(ast.NewUnit())
		}
		return s.Push(addr)
	})
}

// Integer sum
func NativeSumLong(s Context) error {
	return with(s, func(i int64, i2 int64) error {
		return s.Push(ast.NewLong(i + i2))
	})
}
