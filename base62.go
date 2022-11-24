package base62

import (
	"errors"
	"fmt"
)

func init() {
	if len(indexes) != len(Charset) {
		panic(fmt.Sprintf("charset: %d. charindex: %d\n", len(Charset), len(indexes)))
	}
	if len(indexes) != len(values) {
		panic(fmt.Sprintf("charvalue: %d. charindex: %d\n", len(values), len(indexes)))
	}
}

var (
	ErrOverflow         = errors.New("value is overflow")
	ErrInvalidCharacter = errors.New("invalid character")
)

func Encode(value uint64) string {
	if value < uint64(len(Charset)) {
		return string(values[int(value)])
	}
	buf := make([]byte, 16)
	currentIndex := len(buf) - 1
	for value > 0 {
		current := value % uint64(len(Charset))
		value = value / uint64(len(Charset))
		buf[currentIndex] = values[int(current)]
		currentIndex -= 1
	}
	return bytes2string(buf[currentIndex+1:])
}

func Decode(value string) (uint64, error) {
	var prevValue, result uint64
	for i := 0; i < len(value); i++ {
		c, exists := indexes[value[i]]
		if !exists {
			return 0, fmt.Errorf("%w: %q", ErrInvalidCharacter, value[i])
		}
		result = result*uint64(len(Charset)) + uint64(c)
		if result < uint64(prevValue) {
			return 0, fmt.Errorf("%w: %q", ErrOverflow, value)
		}
		prevValue = result
	}
	return result, nil
}

func IsOverflow(err error) bool {
	return errors.Is(err, ErrOverflow)
}

func BatchEncode(batch []uint64) []string {
	results := make([]string, len(batch))
	for i, input := range batch {
		results[i] = Encode(input)
	}
	return results
}

func BatchDecode(batch []string) ([]uint64, error) {
	results := make([]uint64, len(batch))
	for i, input := range batch {
		output, err := Decode(input)
		if err != nil {
			return nil, err
		}
		results[i] = output
	}
	return results, nil
}
