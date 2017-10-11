package util

import "github.com/willf/bitset"

func ByteReverse(buff []byte) []byte {
	length := len(buff)
	if length <= 1 {
		return buff
	}
	for i := 0; i < length / 2; i++  {
		buff[i], buff[length - i] = buff[length - i], buff[i]
	}
	return buff
}

func NumOfBitSet(set bitset.BitSet) int {
	var result int = 0
	for i, _ := set.NextSet(0); i >= 0; i, _ = set.NextSet(i + 1)  {
		result++
	}
	return result
}