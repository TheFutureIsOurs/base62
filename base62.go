/*
 * @Author: Daiming Liu (xingrufeng)
 * Copyright (C) Daiming Liu (xingrufeng)
 */

package base62

import (
	"errors"
	"strings"
)

const (
	strs       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base int64 = 62
)

func Encode(num int64) string {
	b := make([]byte, 0, 12)
	for num > 0 {
		i := num % base
		num /= base
		b = append(b, strs[i])
	}
	return string(b)
}

func Decode(s string) (int64, error) {
	var num int64
	for i := len(s) - 1; i >= 0; i-- {
		pos := strings.IndexByte(strs, s[i])
		if pos == -1 {
			return int64(pos), errors.New("char not in base:" + string(s[i]))
		}
		num = num*base + int64(pos)
	}
	return num, nil
}
