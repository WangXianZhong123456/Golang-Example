package main

import ( 
    "crypto/md5" 
    "encoding/hex" 
    "fmt" 
)

func main() { 
    h := md5.New() 
    h.Write([]byte("123456")) // 需要加密的字符串为 123456 
	cipherStr := h.Sum(nil) 
	cipherStr1:= []byte("12 32")
	fmt.Println(cipherStr) 
	fmt.Println(cipherStr1) 
	fmt.Printf("%x\n",cipherStr1) 
	// %x 以16进制打印字符串
	fmt.Printf("%x\n",cipherStr) 
    fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果 
}

//--------
// [225 10 220 57 73 186 89 171 190 86 224 87 242 15 136 62]
// e10adc3949ba59abbe56e057f20f883e
// e10adc3949ba59abbe56e057f20f883e

/*
EncodeToString returns the hexadecimal encoding of src.
func EncodeToString(src []byte) string {
	dst := make([]byte, EncodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

const hextable = "0123456789abcdef"

// EncodedLen returns the length of an encoding of n source bytes.
// Specifically, it returns n * 2.
func EncodedLen(n int) int { return n * 2 }

// Encode encodes src into EncodedLen(len(src))
// bytes of dst. As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// Encode implements hexadecimal encoding.
func Encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}

	return len(src) * 2
}
*/