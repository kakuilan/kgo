package kgo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"hash"
	"io"
	"strconv"
	"strings"
	"time"
)

// Base64Encode 使用 MIME base64 对数据进行编码.
func (ke *LkkEncrypt) Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// Base64Decode 对使用 MIME base64 编码的数据进行解码.
func (ke *LkkEncrypt) Base64Decode(str string) ([]byte, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Base64UrlSafeEncode url安全的Base64Encode,没有'/'和'+'及结尾的'=' .
func (ke *LkkEncrypt) Base64UrlEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}

// Base64UrlDecode url安全的Base64Decode.
func (ke *LkkEncrypt) Base64UrlDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	return base64.URLEncoding.DecodeString(data)
}

// AuthCode 授权码编码或解码;encode为true时编码,为false解码;expiry为有效期,秒;返回结果为加密/解密的字符串和有效期时间戳.
func (ke *LkkEncrypt) AuthCode(str, key string, encode bool, expiry int64) (string, int64) {
	// DYNAMIC_KEY_LEN 动态密钥长度，相同的明文会生成不同密文就是依靠动态密钥
	// 加入随机密钥，可以令密文无任何规律，即便是原文和密钥完全相同，加密结果也会每次不同，增大破解难度。
	// 取值越大，密文变动规律越大，密文变化 = 16 的 DYNAMIC_KEY_LEN 次方
	// 当此值为 0 时，则不产生随机密钥

	if str == "" {
		return "", 0
	} else if !encode && len(str) < DYNAMIC_KEY_LEN {
		return "", 0
	}

	// 密钥
	keyByte := md5Str([]byte(key), 32)

	// 密钥a会参与加解密
	keya := md5Str(keyByte[:16], 32)

	// 密钥b会用来做数据完整性验证
	keyb := md5Str(keyByte[16:], 32)

	// 密钥c用于变化生成的密文
	var keyc []byte
	if encode == false {
		keyc = []byte(str[:DYNAMIC_KEY_LEN])
	} else {
		cLen := 32 - DYNAMIC_KEY_LEN
		now, _ := time.Now().MarshalBinary()
		timeBytes := md5Str(now, 32)
		keyc = timeBytes[cLen:]
	}

	// 参与运算的密钥
	keyd := md5Str(append(keya, keyc...), 32)
	cryptkey := append(keya, keyd...)
	keyLength := len(cryptkey)
	// 明文，前10位用来保存时间戳，解密时验证数据有效性，10到26位用来保存keyb(密钥b)，解密时会通过这个密钥验证数据完整性
	// 如果是解码的话，会从第 DYNAMIC_KEY_LEN 位开始，因为密文前 DYNAMIC_KEY_LEN 位保存 动态密钥，以保证解密正确
	if encode == false {
		strByte, err := ke.Base64UrlDecode(str[DYNAMIC_KEY_LEN:])
		if err != nil {
			return "", 0
		}
		str = string(strByte)
	} else {
		if expiry != 0 {
			expiry = expiry + time.Now().Unix()
		}
		expMd5 := md5Str(append([]byte(str), keyb...), 16)
		str = fmt.Sprintf("%010d%s%s", expiry, expMd5, str)
	}
	stringLength := len(str)
	resdata := make([]byte, 0, stringLength)
	var rndkey, box [256]int
	// 产生密钥簿
	j := 0
	a := 0
	i := 0
	for i = 0; i < 256; i++ {
		rndkey[i] = int(cryptkey[i%keyLength])
		box[i] = i
	}
	// 用固定的算法，打乱密钥簿，增加随机性，好像很复杂，实际上并不会增加密文的强度
	for i = 0; i < 256; i++ {
		j = (j + box[i] + rndkey[i]) % 256
		box[i], box[j] = box[j], box[i]
	}
	// 核心加解密部分
	a = 0
	j = 0
	for i = 0; i < stringLength; i++ {
		a = ((a + 1) % 256)
		j = ((j + box[a]) % 256)
		box[a], box[j] = box[j], box[a]
		// 从密钥簿得出密钥进行异或，再转成字符
		resdata = append(resdata, byte(int(str[i])^box[(box[a]+box[j])%256]))
	}
	result := string(resdata)
	if encode == false { //解密
		// substr($result, 0, 10) == 0 验证数据有效性
		// substr($result, 0, 10) - time() > 0 验证数据有效性
		// substr($result, 10, 16) == substr(md5(substr($result, 26).$keyb), 0, 16) 验证数据完整性
		// 验证数据有效性，请看未加密明文的格式
		if len(result) <= 26 {
			return "", 0
		}

		expTime, _ := strconv.ParseInt(result[:10], 10, 0)
		if (expTime == 0 || expTime-time.Now().Unix() > 0) && result[10:26] == string(md5Str(append(resdata[26:], keyb...), 16)) {
			return result[26:], expTime
		} else {
			return "", expTime
		}
	} else { //加密
		// 把动态密钥保存在密文里，这也是为什么同样的明文，生产不同密文后能解密的原因
		result = string(keyc) + ke.Base64UrlEncode(resdata)
		return result, expiry
	}
}

// PasswordHash 创建密码的散列值;costs为算法的cost,范围4~31,默认10;注意:值越大越耗时.
func (ke *LkkEncrypt) PasswordHash(password []byte, costs ...int) ([]byte, error) {
	var cost int
	if len(costs) == 0 {
		cost = 10
	} else {
		cost = costs[0]
		if cost < 4 {
			cost = 4
		} else if cost > 31 {
			cost = 15
		}
	}

	bytes, err := bcrypt.GenerateFromPassword(password, cost)
	return bytes, err
}

// PasswordVerify 验证密码是否和散列值匹配.
func (ke *LkkEncrypt) PasswordVerify(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

// EasyEncrypt 简单加密.
// data为要加密的原字符串,key为密钥.
func (ke *LkkEncrypt) EasyEncrypt(data, key string) string {
	dataLen := len(data)
	if dataLen == 0 {
		return ""
	}

	keyByte := md5Str([]byte(key), 32)
	keyLen := len(keyByte)

	var i, x, c int
	var str []byte
	for i = 0; i < dataLen; i++ {
		if x == keyLen {
			x = 0
		}

		c = (int(data[i]) + int(keyByte[x])) % 256
		str = append(str, byte(c))

		x++
	}

	res := string(keyByte[:DYNAMIC_KEY_LEN]) + ke.Base64UrlEncode(str)
	return res
}

// EasyDecrypt 简单解密.
// val为待解密的字符串,key为密钥.
func (ke *LkkEncrypt) EasyDecrypt(val, key string) string {
	if len(val) <= DYNAMIC_KEY_LEN {
		return ""
	}

	data, err := ke.Base64UrlDecode(val[DYNAMIC_KEY_LEN:])
	if err != nil {
		return ""
	}

	keyByte := md5Str([]byte(key), 32)
	if val[:DYNAMIC_KEY_LEN] != string(keyByte[:DYNAMIC_KEY_LEN]) {
		return ""
	}

	dataLen := len(data)
	keyLen := len(keyByte)

	var i, x, c int
	var str []byte
	for i = 0; i < dataLen; i++ {
		if x == keyLen {
			x = 0
		}

		if data[i] < keyByte[x] {
			c = int(data[i]) + 256 - int(keyByte[x])
		} else {
			c = int(data[i]) - int(keyByte[x])
		}
		str = append(str, byte(c))

		x++
	}

	return string(str)
}

// HmacShaX HmacSHA-x加密,x为1/256/512 .
func (ke *LkkEncrypt) HmacShaX(data, secret []byte, x uint16) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	var h hash.Hash
	switch x {
	case 1:
		h = hmac.New(sha1.New, secret)
		break
	case 256:
		h = hmac.New(sha256.New, secret)
		break
	case 512:
		h = hmac.New(sha512.New, secret)
		break
	default:
		panic("[HmacShaX] x must be in [1, 256, 512]")
	}

	// Write Data to it
	h.Write(data)

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

// aesEncrypt AES加密.mode为模式,枚举值(CBC,CFB,CTR).
func (ke *LkkEncrypt) aesEncrypt(clearText, key []byte, mode string, paddingType ...LkkPKCSType) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pt := PKCS7
	blockSize := block.BlockSize()
	if len(paddingType) > 0 {
		pt = paddingType[0]
	}
	switch pt {
	case PKCS0:
		clearText = zeroPadding(clearText, blockSize)
	case PKCS7:
		clearText = pkcs7Padding(clearText, blockSize, false)
	}

	cipherText := make([]byte, blockSize+len(clearText))
	//初始化向量
	iv := cipherText[:blockSize]
	_, _ = io.ReadFull(rand.Reader, iv)
	//if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	//	return nil, err
	//}

	switch mode {
	case "CBC":
		cipher.NewCBCEncrypter(block, iv).CryptBlocks(cipherText[blockSize:], clearText)
	case "CFB":
		cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText[blockSize:], clearText)
	case "CTR":
		cipher.NewCTR(block, iv).XORKeyStream(cipherText[blockSize:], clearText)
	}

	return cipherText, nil
}

// aesDecrypt AES解密.mode为模式,枚举值(CBC,CFB,CTR).
func (ke *LkkEncrypt) aesDecrypt(cipherText, key []byte, mode string, paddingType ...LkkPKCSType) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pt := PKCS7
	if len(paddingType) > 0 {
		pt = paddingType[0]
	}

	blockSize := block.BlockSize()
	clen := len(cipherText)
	if clen < blockSize {
		return nil, errors.New("cipherText too short")
	}

	iv := cipherText[:blockSize]
	cipherText = cipherText[blockSize:]

	switch mode {
	case "CBC":
		cipher.NewCBCDecrypter(block, iv).CryptBlocks(cipherText, cipherText)
	case "CFB":
		cipher.NewCFBDecrypter(block, iv).XORKeyStream(cipherText, cipherText)
	case "CTR":
		cipher.NewCTR(block, iv).XORKeyStream(cipherText, cipherText)
	}

	clen = len(cipherText)
	if clen > 0 && int(cipherText[clen-1]) > clen {
		return nil, errors.New(fmt.Sprintf("aes [%s] decrypt failed", mode))
	}

	var plainText []byte
	switch pt {
	case PKCS0:
		plainText = zeroUnPadding(cipherText)
	case PKCS7:
		plainText = pkcs7UnPadding(cipherText, blockSize)
	}

	return plainText, nil
}

// AesCBCEncrypt AES-CBC密码分组链接(Cipher-block chaining)模式加密.
// clearText为明文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS0,PKCS7),默认PKCS7.
func (ke *LkkEncrypt) AesCBCEncrypt(clearText, key []byte, paddingType ...LkkPKCSType) ([]byte, error) {
	return ke.aesEncrypt(clearText, key, "CBC", paddingType...)
}

// AesCBCDecrypt AES-CBC密码分组链接(Cipher-block chaining)模式解密.
// cipherText为密文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS0,PKCS7),默认PKCS7.
func (ke *LkkEncrypt) AesCBCDecrypt(cipherText, key []byte, paddingType ...LkkPKCSType) ([]byte, error) {
	return ke.aesDecrypt(cipherText, key, "CBC", paddingType...)
}

// AesCFBEncrypt AES-CFB密文反馈(Cipher feedback)模式加密.
// clearText为明文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS0,PKCS7),默认PKCS7.
func (ke *LkkEncrypt) AesCFBEncrypt(clearText, key []byte, paddingType ...LkkPKCSType) ([]byte, error) {
	return ke.aesEncrypt(clearText, key, "CFB", paddingType...)
}

// AesCFBDecrypt AES-CFB密文反馈(Cipher feedback)模式解密.
// cipherText为密文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS0,PKCS7),默认PKCS7.
func (ke *LkkEncrypt) AesCFBDecrypt(cipherText, key []byte, paddingType ...LkkPKCSType) ([]byte, error) {
	return ke.aesDecrypt(cipherText, key, "CFB", paddingType...)
}

// AesECBEncrypt AES-CTR计算器(Counter)模式加密.
// clearText为明文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS0,PKCS7),默认PKCS7.
func (ke *LkkEncrypt) AesCTREncrypt(clearText, key []byte, paddingType ...LkkPKCSType) ([]byte, error) {
	return ke.aesEncrypt(clearText, key, "CTR", paddingType...)
}

// AesECBDecrypt AES-CTR计算器(Counter)模式解密.
// cipherText为密文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS0,PKCS7),默认PKCS7.
func (ke *LkkEncrypt) AesCTRDecrypt(cipherText, key []byte, paddingType ...LkkPKCSType) ([]byte, error) {
	return ke.aesDecrypt(cipherText, key, "CTR", paddingType...)
}
