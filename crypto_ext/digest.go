package crypto_ext

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Digest(msg []byte, algo string) []byte {
	switch algo {
	case "md5":
		hash := md5.Sum(msg)
		return hash[:]
	case "sha1":
		hash := sha1.Sum(msg)
		return hash[:]
	}
	panic(fmt.Sprintf("algorithm %s not supported", algo))
}

func DigestString(msg []byte, algo string) string {
	return ToHexString(Digest(msg, algo))
}

func ToHexString(data []byte) string {
	return fmt.Sprintf("%x", data)
}
