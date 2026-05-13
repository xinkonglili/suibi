package sign

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"testing"
)

func TestSign(t *testing.T) {
	stringToSign := "uiuo00ACFDF646A63W"
	t.Log(GenerateSign(stringToSign))
}

func GenerateSign(originalString string) string {
	hash := md5.Sum([]byte(originalString))
	//剔除到字母，输出纯数字的签名
	encodeToString := hex.EncodeToString(hash[:])
	toString := hex.EncodeToString([]byte(encodeToString))
	return strings.ToUpper(toString)
}

func IsDoubleHexEncoded(sign string) (bool, string) {
	// 检查长度是否为64位(二次编码后的长度)
	if len(sign) != 64 {
		return false, ""
	}

	firstDecode, err := hex.DecodeString(sign)
	if err != nil {
		return false, ""
	}

	// 检查解码后是否是32个字符(标准MD5 hex长度)
	if len(firstDecode) != 32 {
		return false, ""
	}

	// 检查解码后的内容是否是合法的hex字符串(0-9, a-f, A-F)
	md5Hex := string(firstDecode)
	for _, char := range md5Hex {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			return false, ""
		}
	}
	return true, md5Hex
}
