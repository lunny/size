// Copyright 2018 Lunny Xiao. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package size

import (
	"fmt"
	"strconv"
	"strings"
)

// Size represents a size type
type Size float64

// all possible size
const (
	B = Size(1)
	K = 1024 * B
	M = 1024 * K
	G = 1024 * M
	T = 1024 * G
	P = 1024 * T
)

// String shows the size's format layout
func (s Size) String() string {
	switch {
	case s < K:
		return fmt.Sprintf("%dB", int64(s))
	case s < M:
		return formatSize(float64(s)/float64(K), "K")
	case s < G:
		return formatSize(float64(s)/float64(M), "M")
	case s < T:
		return formatSize(float64(s)/float64(G), "G")
	case s < P:
		return formatSize(float64(s)/float64(T), "T")
	default:
		return formatSize(float64(s)/float64(P), "P")
	}
}

func formatSize(f float64, suffix string) string {
	s := fmt.Sprintf("%.3f", f)
	return trimZero(s) + suffix
}

func trimZero(s string) string {
	if !strings.Contains(s, ".") {
		return s
	}
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func calcSize(layout string, suffix Size) (Size, error) {
	size, err := strconv.ParseFloat(strings.TrimSpace(layout), 64)
	if err != nil {
		return 0, err
	}
	return Size(size * float64(suffix)), nil
}

// ParseSize parses the layout to size
func ParseSize(layout string) (Size, error) {
	if len(layout) <= 0 {
		return 0, nil
	}

	switch layout[len(layout)-1] {
	case 'B':
		return calcSize(layout[:len(layout)-1], B)
	case 'K':
		return calcSize(layout[:len(layout)-1], K)
	case 'M':
		return calcSize(layout[:len(layout)-1], M)
	case 'G':
		return calcSize(layout[:len(layout)-1], G)
	case 'T':
		return calcSize(layout[:len(layout)-1], T)
	case 'P':
		return calcSize(layout[:len(layout)-1], P)
	default:
		return calcSize(layout, B)
	}
}
