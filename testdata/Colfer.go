// Package testdata covers the mapping for all supported lanugages.
package testdata

// This file was generated by colf(1); DO NOT EDIT
// The compiler used schema file test.colf.

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"time"
)

var intconv = binary.BigEndian

// Colfer configuration attributes
var (
	// ColferSizeMax is the upper limit for serial byte sizes.
	ColferSizeMax = 16 * 1024 * 1024
	// ColferListMax is the upper limit for the number of elements in a list.
	ColferListMax = 64 * 1024
)

// ColferMax signals an upper limit breach.
type ColferMax string

// Error honors the error interface.
func (m ColferMax) Error() string { return string(m) }

// ColferError signals a data mismatch as as a byte index.
type ColferError int

// Error honors the error interface.
func (i ColferError) Error() string {
	return fmt.Sprintf("colfer: unknown header at byte %d", i)
}

// ColferTail signals data continuation as a byte index.
type ColferTail int

// Error honors the error interface.
func (i ColferTail) Error() string {
	return fmt.Sprintf("colfer: data continuation at byte %d", i)
}

// O contains all supported data types.
type O struct {
	// B tests booleans.
	B bool
	// U32 tests unsigned 32-bit integers.
	U32 uint32
	// U64 tests unsigned 64-bit integers.
	U64 uint64
	// I32 tests signed 32-bit integers.
	I32 int32
	// I64 tests signed 64-bit integers.
	I64 int64
	// F32 tests 32-bit floating points.
	F32 float32
	// F64 tests 64-bit floating points.
	F64 float64
	// T tests timestamps.
	T time.Time
	// S tests text.
	S string
	// A tests binaries.
	A []byte
	// O tests nested data structures.
	O *O
	// Os tests data structure lists.
	Os []*O
	// Ss tests text lists.
	Ss []string
	// As tests binary lists.
	As [][]byte
	// U8 tests unsigned 8-bit integers.
	U8 uint8
	// U16 tests unsigned 16-bit integers.
	U16 uint16
	// F32s tests 32-bit floating point lists.
	F32s []float32
	// F64s tests 64-bit floating point lists.
	F64s []float64
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
// All nil entries in o.Os will be replaced with a new value.
func (o *O) MarshalTo(buf []byte) int {
	var i int

	if o.B {
		buf[i] = 0
		i++
	}

	if x := o.U32; x >= 1<<21 {
		buf[i] = 1 | 0x80
		intconv.PutUint32(buf[i+1:], x)
		i += 5
	} else if x != 0 {
		buf[i] = 1
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if x := o.U64; x >= 1<<49 {
		buf[i] = 2 | 0x80
		intconv.PutUint64(buf[i+1:], x)
		i += 9
	} else if x != 0 {
		buf[i] = 2
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.I32; v != 0 {
		x := uint32(v)
		if v >= 0 {
			buf[i] = 3
		} else {
			x = ^x + 1
			buf[i] = 3 | 0x80
		}
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.I64; v != 0 {
		x := uint64(v)
		if v >= 0 {
			buf[i] = 4
		} else {
			x = ^x + 1
			buf[i] = 4 | 0x80
		}
		i++
		for n := 0; x >= 0x80 && n < 8; n++ {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.F32; v != 0 {
		buf[i] = 5
		intconv.PutUint32(buf[i+1:], math.Float32bits(v))
		i += 5
	}

	if v := o.F64; v != 0 {
		buf[i] = 6
		intconv.PutUint64(buf[i+1:], math.Float64bits(v))
		i += 9
	}

	if v := o.T; !v.IsZero() {
		s, ns := uint64(v.Unix()), uint32(v.Nanosecond())
		if s < 1<<32 {
			buf[i] = 7
			intconv.PutUint32(buf[i+1:], uint32(s))
			i += 5
		} else {
			buf[i] = 7 | 0x80
			intconv.PutUint64(buf[i+1:], s)
			i += 9
		}
		intconv.PutUint32(buf[i:], ns)
		i += 4
	}

	if l := len(o.S); l != 0 {
		buf[i] = 8
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		i += copy(buf[i:], o.S)
	}

	if l := len(o.A); l != 0 {
		buf[i] = 9
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		i += copy(buf[i:], o.A)
	}

	if v := o.O; v != nil {
		buf[i] = 10
		i++
		i += v.MarshalTo(buf[i:])
	}

	if l := len(o.Os); l != 0 {
		buf[i] = 11
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		for vi, v := range o.Os {
			if v == nil {
				v = new(O)
				o.Os[vi] = v
			}
			i += v.MarshalTo(buf[i:])
		}
	}

	if l := len(o.Ss); l != 0 {
		buf[i] = 12
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		for _, a := range o.Ss {
			x = uint(len(a))
			for x >= 0x80 {
				buf[i] = byte(x | 0x80)
				x >>= 7
				i++
			}
			buf[i] = byte(x)
			i++
			i += copy(buf[i:], a)
		}
	}

	if l := len(o.As); l != 0 {
		buf[i] = 13
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		for _, a := range o.As {
			x = uint(len(a))
			for x >= 0x80 {
				buf[i] = byte(x | 0x80)
				x >>= 7
				i++
			}
			buf[i] = byte(x)
			i++
			i += copy(buf[i:], a)
		}
	}

	if x := o.U8; x != 0 {
		buf[i] = 14
		i++
		buf[i] = x
		i++
	}

	if x := o.U16; x >= 1<<8 {
		buf[i] = 15
		i++
		buf[i] = byte(x >> 8)
		i++
		buf[i] = byte(x)
		i++
	} else if x != 0 {
		buf[i] = 15 | 0x80
		i++
		buf[i] = byte(x)
		i++
	}

	if l := len(o.F32s); l != 0 {
		buf[i] = 16
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		for _, v := range o.F32s {
			intconv.PutUint32(buf[i:], math.Float32bits(v))
			i += 4
		}
	}

	if l := len(o.F64s); l != 0 {
		buf[i] = 17
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		for _, v := range o.F64s {
			intconv.PutUint64(buf[i:], math.Float64bits(v))
			i += 8
		}
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is testdata.ColferMax.
func (o *O) MarshalLen() (int, error) {
	l := 1

	if o.B {
		l++
	}

	if x := o.U32; x >= 1<<21 {
		l += 5
	} else if x != 0 {
		l += 2
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if x := o.U64; x >= 1<<49 {
		l += 9
	} else if x != 0 {
		l += 2
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if v := o.I32; v != 0 {
		l += 2
		x := uint32(v)
		if v < 0 {
			x = ^x + 1
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if v := o.I64; v != 0 {
		l += 2
		x := uint64(v)
		if v < 0 {
			x = ^x + 1
		}
		for n := 0; x >= 0x80 && n < 8; n++ {
			x >>= 7
			l++
		}
	}

	if o.F32 != 0 {
		l += 5
	}

	if o.F64 != 0 {
		l += 9
	}

	if v := o.T; !v.IsZero() {
		if s := uint64(v.Unix()); s < 1<<32 {
			l += 9
		} else {
			l += 13
		}
	}

	if x := len(o.S); x != 0 {
		l += x
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if x := len(o.A); x != 0 {
		l += x
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if v := o.O; v != nil {
		vl, err := v.MarshalLen()
		if err != nil {
			return -1, err
		}
		l += vl + 1
	}

	if x := len(o.Os); x != 0 {
		if x > ColferListMax {
			return -1, ColferMax(fmt.Sprintf("colfer: field testdata.o.os exceeds %d elements", ColferListMax))
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
		for _, v := range o.Os {
			if v == nil {
				l++
				continue
			}
			vl, err := v.MarshalLen()
			if err != nil {
				return -1, err
			}
			l += vl
		}
	}

	if x := len(o.Ss); x != 0 {
		if x > ColferListMax {
			return -1, ColferMax(fmt.Sprintf("colfer: field testdata.o.ss exceeds %d elements", ColferListMax))
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
		for _, a := range o.Ss {
			x = len(a)
			l += x
			for x >= 0x80 {
				x >>= 7
				l++
			}
			l++
		}
	}

	if x := len(o.As); x != 0 {
		if x > ColferListMax {
			return -1, ColferMax(fmt.Sprintf("colfer: field testdata.o.as exceeds %d elements", ColferListMax))
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
		for _, a := range o.As {
			x = len(a)
			l += x
			for x >= 0x80 {
				x >>= 7
				l++
			}
			l++
		}
	}

	if x := o.U8; x != 0 {
		l += 2
	}

	if x := o.U16; x >= 1<<8 {
		l += 3
	} else if x != 0 {
		l += 2
	}

	if x := len(o.F32s); x != 0 {
		l += 2 + x*4
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if x := len(o.F64s); x != 0 {
		l += 2 + x*8
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct testdata.o exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// All nil entries in o.Os will be replaced with a new value.
// The error return option is testdata.ColferMax.
func (o *O) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, testdata.ColferError and testdata.ColferMax.
func (o *O) Unmarshal(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		if i >= len(data) {
			goto eof
		}
		o.B = true
		header = data[i]
		i++
	}

	if header == 1 {
		start := i
		i++
		if i >= len(data) {
			goto eof
		}
		x := uint32(data[start])

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint32(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.U32 = x

		header = data[i]
		i++
	} else if header == 1|0x80 {
		start := i
		i += 4
		if i >= len(data) {
			goto eof
		}
		o.U32 = intconv.Uint32(data[start:])
		header = data[i]
		i++
	}

	if header == 2 {
		start := i
		i++
		if i >= len(data) {
			goto eof
		}
		x := uint64(data[start])

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint64(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 || shift == 56 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.U64 = x

		header = data[i]
		i++
	} else if header == 2|0x80 {
		start := i
		i += 8
		if i >= len(data) {
			goto eof
		}
		o.U64 = intconv.Uint64(data[start:])
		header = data[i]
		i++
	}

	if header == 3 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint32(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint32(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.I32 = int32(x)

		header = data[i]
		i++
	} else if header == 3|0x80 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint32(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint32(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.I32 = int32(^x + 1)

		header = data[i]
		i++
	}

	if header == 4 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint64(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint64(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 || shift == 56 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.I64 = int64(x)

		header = data[i]
		i++
	} else if header == 4|0x80 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint64(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint64(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 || shift == 56 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.I64 = int64(^x + 1)

		header = data[i]
		i++
	}

	if header == 5 {
		start := i
		i += 4
		if i >= len(data) {
			goto eof
		}
		o.F32 = math.Float32frombits(intconv.Uint32(data[start:]))
		header = data[i]
		i++
	}

	if header == 6 {
		start := i
		i += 8
		if i >= len(data) {
			goto eof
		}
		o.F64 = math.Float64frombits(intconv.Uint64(data[start:]))
		header = data[i]
		i++
	}

	if header == 7 {
		start := i
		i += 8
		if i >= len(data) {
			goto eof
		}
		o.T = time.Unix(int64(intconv.Uint32(data[start:])), int64(intconv.Uint32(data[start+4:]))).In(time.UTC)
		header = data[i]
		i++
	} else if header == 7|0x80 {
		start := i
		i += 12
		if i >= len(data) {
			goto eof
		}
		o.T = time.Unix(int64(intconv.Uint64(data[start:])), int64(intconv.Uint32(data[start+8:]))).In(time.UTC)
		header = data[i]
		i++
	}

	if header == 8 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.s size %d exceeds %d bytes", x, ColferSizeMax))
		}

		start := i
		i += int(x)
		if i >= len(data) {
			goto eof
		}
		o.S = string(data[start:i])

		header = data[i]
		i++
	}

	if header == 9 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.a size %d exceeds %d bytes", x, ColferSizeMax))
		}
		v := make([]byte, int(x))

		start := i
		i += len(v)
		if i >= len(data) {
			goto eof
		}
		copy(v, data[start:i])
		o.A = v

		header = data[i]
		i++
	}

	if header == 10 {
		o.O = new(O)
		n, err := o.O.Unmarshal(data[i:])
		if err != nil {
			if err == io.EOF && len(data) >= ColferSizeMax {
				return 0, ColferMax(fmt.Sprintf("colfer: testdata.o size exceeds %d bytes", ColferSizeMax))
			}
			return 0, err
		}
		i += n

		if i >= len(data) {
			goto eof
		}
		header = data[i]
		i++
	}

	if header == 11 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferListMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.os length %d exceeds %d elements", x, ColferListMax))
		}

		l := int(x)
		a := make([]*O, l)
		malloc := make([]O, l)
		for ai, _ := range a {
			v := &malloc[ai]
			a[ai] = v

			n, err := v.Unmarshal(data[i:])
			if err != nil {
				if err == io.EOF && len(data) >= ColferSizeMax {
					return 0, ColferMax(fmt.Sprintf("colfer: testdata.o size exceeds %d bytes", ColferSizeMax))
				}
				return 0, err
			}
			i += n
		}
		o.Os = a

		if i >= len(data) {
			goto eof
		}
		header = data[i]
		i++
	}

	if header == 12 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferListMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.ss length %d exceeds %d elements", x, ColferListMax))
		}
		a := make([]string, int(x))
		o.Ss = a

		for ai := range a {
			if i >= len(data) {
				goto eof
			}
			x := uint(data[i])
			i++

			if x >= 0x80 {
				x &= 0x7f
				for shift := uint(7); ; shift += 7 {
					if i >= len(data) {
						goto eof
					}
					b := uint(data[i])
					i++

					if b < 0x80 {
						x |= b << shift
						break
					}
					x |= (b & 0x7f) << shift
				}
			}

			if x > uint(ColferSizeMax) {
				return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.ss element %d size %d exceeds %d bytes", ai, x, ColferSizeMax))
			}

			start := i
			i += int(x)
			if i >= len(data) {
				goto eof
			}
			a[ai] = string(data[start:i])
		}

		if i >= len(data) {
			goto eof
		}
		header = data[i]
		i++
	}

	if header == 13 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferListMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.as length %d exceeds %d elements", x, ColferListMax))
		}
		a := make([][]byte, int(x))
		o.As = a
		for ai := range a {
			if i >= len(data) {
				goto eof
			}
			x := uint(data[i])
			i++

			if x >= 0x80 {
				x &= 0x7f
				for shift := uint(7); ; shift += 7 {
					if i >= len(data) {
						goto eof
					}
					b := uint(data[i])
					i++

					if b < 0x80 {
						x |= b << shift
						break
					}
					x |= (b & 0x7f) << shift
				}
			}

			if x > uint(ColferSizeMax) {
				return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.as element %d size %d exceeds %d bytes", ai, x, ColferSizeMax))
			}
			v := make([]byte, int(x))

			start := i
			i += len(v)
			if i >= len(data) {
				goto eof
			}

			copy(v, data[start:i])
			a[ai] = v
		}

		if i >= len(data) {
			goto eof
		}
		header = data[i]
		i++
	}

	if header == 14 {
		start := i
		i++
		if i >= len(data) {
			goto eof
		}
		o.U8 = data[start]
		header = data[i]
		i++
	}

	if header == 15 {
		start := i
		i += 2
		if i >= len(data) {
			goto eof
		}
		o.U16 = intconv.Uint16(data[start:])
		header = data[i]
		i++
	} else if header == 15|0x80 {
		start := i
		i++
		if i >= len(data) {
			goto eof
		}
		o.U16 = uint16(data[start])
		header = data[i]
		i++
	}

	if header == 16 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferListMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.f32s length %d exceeds %d elements", x, ColferListMax))
		}

		l := int(x)

		if end := i + l*4; end >= len(data) {
			i = end
			goto eof
		}
		a := make([]float32, l)
		for ai := range a {
			a[ai] = math.Float32frombits(intconv.Uint32(data[i:]))
			i += 4
		}
		o.F32s = a

		header = data[i]
		i++
	}

	if header == 17 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferListMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: testdata.o.f64s length %d exceeds %d elements", x, ColferListMax))
		}
		l := int(x)

		if end := i + l*8; end >= len(data) {
			i = end
			goto eof
		}
		a := make([]float64, l)
		for ai := range a {
			a[ai] = math.Float64frombits(intconv.Uint64(data[i:]))
			i += 8
		}
		o.F64s = a

		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	if i < ColferSizeMax {
		return i, nil
	}
eof:
	if i >= ColferSizeMax {
		return 0, ColferMax(fmt.Sprintf("colfer: testdata.o size exceeds %d bytes", ColferSizeMax))
	}
	return 0, io.EOF
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, testdata.ColferError, testdata.ColferTail and testdata.ColferMax.
func (o *O) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if i < len(data) && err == nil {
		return ColferTail(i)
	}
	return err
}
