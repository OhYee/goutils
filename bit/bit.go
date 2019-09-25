package bit

import (
	"bytes"
	"fmt"
)

type Bit struct {
	data       []byte
	bitLength  int
	byteLength int
}

// NewBit declare a Bit struct
func NewBit(bitLength int, bytes ...byte) (b Bit, err error) {
	byteLength := bitLength / 8
	if bitLength%8 != 0 {
		byteLength++
	}

	if byteLength < len(bytes) {
		err = fmt.Errorf("Bit error: The bit is %d bytes (%d bits) length, but got %d bytes data", byteLength, bitLength, len(bytes))
		return
	}
	b = Bit{
		data:       make([]byte, byteLength),
		bitLength:  bitLength,
		byteLength: byteLength,
	}
	for idx, bb := range bytes {
		b.data[idx] = bb
	}
	return
}

// SetValue the pos-th bit to 1 or 0
func (bit *Bit) SetValue(pos int, value byte) (err error) {
	if pos > bit.bitLength {
		err = fmt.Errorf("Bit error: Index out of range [%d] with %d", pos, bit.bitLength)
		return
	}
	if value == 0 {
		bit.Clear(pos)
	} else {
		bit.Set(pos)
	}
	return
}

// Set the pos-th bit to 1
func (bit *Bit) Set(pos int) (err error) {
	if pos > bit.bitLength {
		err = fmt.Errorf("Bit error: Index out of range [%d] with %d", pos, bit.bitLength)
		return
	}
	bit.data[pos>>3] |= 1 << (pos & 7)
	return
}

// SetByte the pos-th byte
func (bit *Bit) SetByte(pos int, value byte) (err error) {
	if pos > bit.byteLength {
		err = fmt.Errorf("Bit error: Index out of range [%d] with %d", pos, bit.byteLength)
		return
	}
	bit.data[pos] = value
	return
}

// Clear the pos-th bit (set to 0)
func (bit *Bit) Clear(pos int) (err error) {
	if pos > bit.bitLength {
		err = fmt.Errorf("Bit error: Index out of range [%d] with %d", pos, bit.bitLength)
		return
	}
	bit.data[pos>>3] &= ^(1 << (pos & 7))
	return
}

// Reset the Bit (Set all is 0)
func (bit *Bit) Reset() {
	for i := range bit.data {
		bit.data[i] = 0
	}
}

// Get the pos-th bit value
func (bit Bit) Get(pos int) (b byte, err error) {
	if pos > bit.bitLength {
		err = fmt.Errorf("Bit error: Index out of range [%d] with %d", pos, bit.bitLength)
		return
	}
	b = bit.data[pos>>3] & (1 << (pos & 7))
	return
}

// GetByte get the pos-tg byte value
func (bit Bit) GetByte(pos int) (b byte, err error) {
	if pos > bit.byteLength {
		err = fmt.Errorf("Bit error: Index out of range [%d] with %d", pos, bit.byteLength)
		return
	}
	b = bit.data[pos]
	return
}

// Compare with another Bit, return 0,1,-1
func (bit Bit) Compare(bb Bit) int {
	if bit.bitLength == bb.bitLength {
		for i := 0; i < bit.byteLength; i++ {
			res := bytes.Compare(bit.data, bb.data)
			if res != 0 {
				return res
			}
		}
		return 0
	}
	return bit.bitLength - bb.bitLength
}

// Sub get the part of the Bit from start to end
func (bit Bit) Sub(start int, end int) (res Bit) {
	if end > start {
		res, _ = NewBit(end - start)
		for i := start; i < end; i++ {
			if i%8 == 0 {
				b, _ := bit.GetByte((i / 8) % bit.byteLength)
				res.SetByte((i-start)/8, b)
			} else {
				b, _ := res.Get(i % bit.bitLength)
				res.SetValue(i-start, b)
			}
		}
	} else {
		res, _ = NewBit(bit.bitLength - start + end)
		for i := start; i < end+bit.bitLength; i++ {
			if i%8 == 0 {
				b, _ := bit.GetByte((i / 8) % bit.byteLength)
				res.SetByte((i-start)/8, b)
			} else {
				b, _ := res.Get(i % bit.bitLength)
				res.SetValue(i-start, b)
			}
		}
	}
	return
}

// // Xor two Bit struct
// func (bit Bit) Xor(bb Bit) (res Bit, err error) {
// 	if bit.bitLength != bb.bitLength {
// 		err = fmt.Errorf("Bit error: Can not xor Bit (%d bits) with Bit (%d bits)", bit.bitLength, bb.bitLength)
// 		return
// 	}
// 	var t1, t2 *Bit
// 	if bit.bitLength > bb.bitLength {
// 		t1 = &bb
// 		t2 = &bit
// 	} else {
// 		t1 = &bit
// 		t2 = &bb
// 	}
// 	res, _ = NewBit(t2.bitLength)
// 	for i := range t1.data {
// 		res.SetByte(i, t1.GetByte(i)^t2.GetByte(i))
// 	}
// 	for i := t1.bitLength; i < t2.bitLength; i++ {
// 		res.SetByte(i, t2.GetByte(i))
// 	}
// 	return res
// }
