// Copyright 2018 Lunny Xiao. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package size

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSize(t *testing.T) {
	var kases = []struct {
		Layout string
		Size   Size
	}{
		{"10B", 10 * B},
		{"10", 10 * B},
		{"2.0K", 2 * K},
		{"3M", 3 * M},
		{"4G", 4 * G},
		{"1.67M", 1.67 * M},
		{"1.011T", 1.011 * T},
		{"0.1P", 0.1 * P},
		{"10P", 10 * P},
	}

	for _, k := range kases {
		size, err := ParseSize(k.Layout)
		assert.NoError(t, err)
		assert.EqualValues(t, k.Size, size)
	}
}

func TestFormatSize(t *testing.T) {
	var kases = []struct {
		Size   Size
		Format string
	}{
		{0, "0B"},
		{10 * B, "10B"},
		{2 * K, "2K"},
		{3 * M, "3M"},
		{4 * G, "4G"},
		{1.67 * M, "1.67M"},
		{1.011 * T, "1.011T"},
		{15 * P, "15P"},
	}

	for _, k := range kases {
		assert.EqualValues(t, k.Format, k.Size.String())
	}
}

func TestErrorSizeParse(t *testing.T) {
	var kases = []string{
		"10BB",
		"",
		"B",
		"K",
		"M",
		"G",
		"T",
		"P",
		"2-1K",
		"1*2",
		"-1K",
	}
	for _, k := range kases {
		_, err := ParseSize(k)
		assert.Error(t, err)
		assert.EqualValues(t, fmt.Sprintf("size format %s is not corrected", k), err.Error())
	}
}
