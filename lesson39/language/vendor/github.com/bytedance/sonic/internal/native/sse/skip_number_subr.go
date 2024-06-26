// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__skip_number = 112
)

const (
    _stack__skip_number = 72
)

const (
    _size__skip_number = 1128
)

var (
    _pcsp__skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1046, 72},
        {1050, 48},
        {1051, 40},
        {1053, 32},
        {1055, 24},
        {1057, 16},
        {1059, 8},
        {1060, 0},
        {1128, 72},
    }
)

var _cfunc_skip_number = []loader.CFunc{
    {"_skip_number_entry", 0,  _entry__skip_number, 0, nil},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
}
