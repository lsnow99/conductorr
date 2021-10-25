package csl

import (
	"github.com/lsnow99/conductorr/constant"
	"github.com/lsnow99/conductorr/csllib"
)

var CSL = csllib.NewCSL(true)

func init() {
	CSL.RegisterFunction("r-title", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(0)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-indexer", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(1)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-downloadtype", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(2)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-contenttype", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(3)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-riptype", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(4)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-resolution", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(5)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-encoding", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(6)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-seeders", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(7)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-age", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(8)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-size", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(9)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-runtime", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(10)}, args...)...)
	}, nil)
	CSL.RegisterFunction("r-riptype-order", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		retVal, err := GetNthFromRelease(env, append([]interface{}{int64(4)}, args...)...)
		if err != nil {
			return nil, err
		}
		retStr, ok := retVal.(string)
		if !ok {
			return nil, csllib.ErrBadType
		}
		for val, opts := range constant.RipTypes {
			if val == retStr {
				return int64(opts.Priority), nil
			}
		}
		return int64(0), nil
	}, nil)
	CSL.RegisterFunction("r-resolution-order", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		retVal, err := GetNthFromRelease(env, append([]interface{}{int64(5)}, args...)...)
		if err != nil {
			return nil, err
		}
		retStr, ok := retVal.(string)
		if !ok {
			return nil, csllib.ErrBadType
		}
		for val, opts := range constant.ResolutionTypes {
			if val == retStr {
				return int64(opts.Priority), nil
			}
		}
		return int64(0), nil
	}, nil)
	CSL.RegisterFunction("r-encoding-order", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		retVal, err := GetNthFromRelease(env, append([]interface{}{int64(6)}, args...)...)
		if err != nil {
			return nil, err
		}
		retStr, ok := retVal.(string)
		if !ok {
			return nil, csllib.ErrBadType
		}
		for val, opts := range constant.EncodingTypes {
			if val == retStr {
				return int64(opts.Priority), nil
			}
		}
		return int64(0), nil
	}, nil)
	CSL.RegisterFunction("r-bitrate", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		size, err := GetNthFromRelease(env, append([]interface{}{int64(9)}, args...)...)
		if err != nil {
			return nil, err
		}
		runtime, err := GetNthFromRelease(env, append([]interface{}{int64(10)}, args...)...)
		if err != nil {
			return nil, err
		}
		sizeNum, ok := size.(int64)
		if !ok {
			return nil, csllib.ErrBadType
		}
		runtimeNum, ok := runtime.(int64)
		if !ok {
			return nil, csllib.ErrBadType
		}
		if runtimeNum == 0 {
			return int64(-1), nil
		}
		return int64(sizeNum / (runtimeNum * 60)), nil
	}, nil)
}

func GetNthFromRelease(env map[string]interface{}, args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, csllib.ErrNumOperands
	}
	i, ok := args[0].(int64)
	if !ok {
		return nil, csllib.ErrMismatchOperandTypes
	}
	l, ok := args[1].(csllib.List)
	if !ok {
		if s, ok := args[1].(string); ok {
			if i < 0 || i >= int64(len(s)) {
				return nil, csllib.ErrOutOfBounds
			}
			return string(s[i]), nil
		}
		l = csllib.List{
			Elems: []interface{}{args[1]},
		}
	}
	if i < 0 || i >= int64(len(l.Elems)) {
		return nil, csllib.ErrOutOfBounds
	}
	return l.Elems[i], nil
}
