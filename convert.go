package base62

import "unsafe"

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
