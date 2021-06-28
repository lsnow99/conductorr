package main

import (
	"github.com/lsnow99/conductorr/constant"
	"github.com/lsnow99/conductorr/csl"
)

func init() {
	csl.RegisterFunction("r-title", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(0)}, args...)...)
	})
	csl.RegisterFunction("r-indexer", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(1)}, args...)...)
	})
	csl.RegisterFunction("r-downloadtype", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(2)}, args...)...)
	})
	csl.RegisterFunction("r-contenttype", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(3)}, args...)...)
	})
	csl.RegisterFunction("r-riptype", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(4)}, args...)...)
	})
	csl.RegisterFunction("r-resolution", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(5)}, args...)...)
	})
	csl.RegisterFunction("r-encoding", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(6)}, args...)...)
	})
	csl.RegisterFunction("r-seeders", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(7)}, args...)...)
	})
	csl.RegisterFunction("r-age", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(8)}, args...)...)
	})
	csl.RegisterFunction("r-size", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(9)}, args...)...)
	})
	csl.RegisterFunction("r-runtime", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(10)}, args...)...)
	})
	csl.RegisterFunction("r-riptype-order", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		retVal, err := GetNthFromRelease(env, append([]interface{}{int64(4)}, args...)...)
		if err != nil {
			return nil, err
		}
		retStr, ok := retVal.(string)
		if !ok {
			return nil, csl.ErrBadType
		}
		for val, opts := range constant.RipTypes {
			if val == retStr {
				return int64(opts.Priority), nil
			}
		}
		return int64(0), nil
	})
	csl.RegisterFunction("r-resolution-order", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		retVal, err := GetNthFromRelease(env, append([]interface{}{int64(5)}, args...)...)
		if err != nil {
			return nil, err
		}
		retStr, ok := retVal.(string)
		if !ok {
			return nil, csl.ErrBadType
		}
		for val, opts := range constant.ResolutionTypes {
			if val == retStr {
				return int64(opts.Priority), nil
			}
		}
		return int64(0), nil
	})
	csl.RegisterFunction("r-encoding-order", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		retVal, err := GetNthFromRelease(env, append([]interface{}{int64(6)}, args...)...)
		if err != nil {
			return nil, err
		}
		retStr, ok := retVal.(string)
		if !ok {
			return nil, csl.ErrBadType
		}
		for val, opts := range constant.EncodingTypes {
			if val == retStr {
				return int64(opts.Priority), nil
			}
		}
		return int64(0), nil
	})
	csl.RegisterFunction("r-bitrate", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
			return nil, csl.ErrBadType
		}
		runtimeNum, ok := runtime.(int64)
		if !ok {
			return nil, csl.ErrBadType
		}
		return int64(sizeNum / (runtimeNum * 60)), nil
	})
}

func GetNthFromRelease(env map[string]interface{}, args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, csl.ErrNumOperands
	}
	i, ok := args[0].(int64)
	if !ok {
		return nil, csl.ErrMismatchOperandTypes
	}
	l, ok := args[1].(csl.List)
	if !ok {
		if s, ok := args[1].(string); ok {
			if i < 0 || i >= int64(len(s)) {
				return nil, csl.ErrOutOfBounds
			}
			return string(s[i]), nil
		}
		l = csl.List{
			Elems: []interface{}{args[1]},
		}
	}
	if i < 0 || i >= int64(len(l.Elems)) {
		return nil, csl.ErrOutOfBounds
	}
	return l.Elems[i], nil
}
