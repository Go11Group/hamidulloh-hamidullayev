// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__u64toa = 80
)

const (
    _stack__u64toa = 8
)

const (
    _size__u64toa = 1264
)

var (
    _pcsp__u64toa = [][2]uint32{
        {1, 0},
        {161, 8},
        {162, 0},
        {457, 8},
        {458, 0},
        {772, 8},
        {773, 0},
        {1249, 8},
        {1251, 0},
    }
)

var _cfunc_u64toa = []loader.CFunc{
    {"_u64toa_entry", 0,  _entry__u64toa, 0, nil},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
}