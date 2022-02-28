package csl

import (
	"fmt"

	"github.com/lsnow99/conductorr/pkg/constant"
	"github.com/lsnow99/conductorr/pkg/csllib"
)

func NewCSL() *csllib.CSL {
	csl := csllib.NewCSL(true)
	csl.RegisterFunction("r-title", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(0)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-indexer", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(1)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-downloadtype", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(2)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-contenttype", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(3)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-riptype", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(4)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-resolution", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(5)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-encoding", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(6)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-seeders", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(7)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-age", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(8)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-size", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(9)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-runtime", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		return GetNthFromRelease(env, append([]interface{}{int64(10)}, args...)...)
	}, nil)
	csl.RegisterFunction("r-riptype-order", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	csl.RegisterFunction("r-resolution-order", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	csl.RegisterFunction("r-encoding-order", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	csl.RegisterFunction("r-bitrate", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	return csl
}

func GetNthFromRelease(env map[string]interface{}, args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, csllib.ErrNumOperands
	}
	i, ok := args[0].(int64)
	if !ok {
		return nil, fmt.Errorf("non list passed as a release: %w", csllib.ErrMismatchOperandTypes)
	}
	l, ok := args[1].(csllib.List)
	if !ok {
		if s, ok := args[1].(string); ok {
			if i < 0 || i >= int64(len(s)) {
				return nil, csllib.ErrOutOfBounds
			}
			return string(s[i]), nil
		}
		l = csllib.List{args[1]}
	}
	if i < 0 || i >= int64(len(l)) {
		return nil, csllib.ErrOutOfBounds
	}
	return l[i], nil
}
