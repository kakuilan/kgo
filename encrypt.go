package kgo

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

// Base64Encode 使用 MIME base64 对数据进行编码.
func (ke *LkkEncrypt) Base64Encode(str []byte) []byte {
	l := len(str)
	if l > 0 {
		buf := make([]byte, base64.StdEncoding.EncodedLen(l))
		base64.StdEncoding.Encode(buf, str)
		return buf
	}

	return nil
}

// Base64Decode 对使用 MIME base64 编码的数据进行解码.
func (ke *LkkEncrypt) Base64Decode(str []byte) ([]byte, error) {
	l := len(str)
	if l > 0 {
		dbuf := make([]byte, base64.StdEncoding.DecodedLen(l))
		n, err := base64.StdEncoding.Decode(dbuf, str)
		return dbuf[:n], err
	}

	return nil, nil
}

// Base64UrlSafeEncode url安全的Base64Encode,没有'/'和'+'及结尾的'=' .
func (ke *LkkEncrypt) Base64UrlEncode(str []byte) []byte {
	l := len(str)
	if l > 0 {
		buf := make([]byte, base64.StdEncoding.EncodedLen(l))
		base64.StdEncoding.Encode(buf, str)

		// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
		buf = bytes.Replace(buf, []byte("/"), []byte("_"), -1)
		buf = bytes.Replace(buf, []byte("+"), []byte("-"), -1)
		buf = bytes.Replace(buf, []byte("="), []byte(""), -1)

		return buf
	}

	return nil
}

// Base64UrlDecode url安全的Base64Decode.
func (ke *LkkEncrypt) Base64UrlDecode(str []byte) ([]byte, error) {
	l := len(str)
	if l > 0 {
		var missing = (4 - len(str)%4) % 4
		str = append(str, bytes.Repeat([]byte("="), missing)...)

		dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(str)))
		n, err := base64.URLEncoding.Decode(dbuf, str)
		return dbuf[:n], err
	}

	return nil, nil
}

// AuthCode 授权码编码或解码;encode为true时编码,为false解码;expiry为有效期,秒;返回结果为加密/解密的字符串和有效期时间戳.
func (ke *LkkEncrypt) AuthCode(str, key []byte, encode bool, expiry int64) ([]byte, int64) {
	// DYNAMIC_KEY_LEN 动态密钥长度，相同的明文会生成不同密文就是依靠动态密钥
	// 加入随机密钥，可以令密文无任何规律，即便是原文和密钥完全相同，加密结果也会每次不同，增大破解难度。
	// 取值越大，密文变动规律越大，密文变化 = 16 的 DYNAMIC_KEY_LEN 次方
	// 当此值为 0 时，则不产生随机密钥

	strLen := len(str)
	if str == nil || strLen == 0 {
		return nil, 0
	} else if !encode && strLen < DYNAMIC_KEY_LEN {
		return nil, 0
	}

	// 密钥
	keyByte := md5Byte(key, 32)

	// 密钥a会参与加解密
	keya := keyByte[:16]

	// 密钥b会用来做数据完整性验证
	keyb := keyByte[16:]

	// 密钥c用于变化生成的密文
	var keyc []byte
	if encode == false {
		keyc = str[:DYNAMIC_KEY_LEN]
	} else {
		now, _ := time.Now().MarshalBinary()
		keycLen := 32 - DYNAMIC_KEY_LEN
		timeBytes := md5Byte(now, 32)
		keyc = timeBytes[keycLen:]
	}

	// 参与运算的密钥
	keyd := md5Byte(append(keya, keyc...), 32)
	cryptkey := append(keya, keyd...)
	cryptkeyLen := len(cryptkey)
	// 明文，前10位用来保存时间戳，解密时验证数据有效性，10到26位用来保存keyb(密钥b)，解密时会通过这个密钥验证数据完整性
	// 如果是解码的话，会从第 DYNAMIC_KEY_LEN 位开始，因为密文前 DYNAMIC_KEY_LEN 位保存 动态密钥，以保证解密正确
	if encode == false { //解密
		var err error
		str, err = ke.Base64UrlDecode(str[DYNAMIC_KEY_LEN:])
		if err != nil {
			return nil, 0
		}
	} else {
		if expiry != 0 {
			expiry = expiry + time.Now().Unix()
		}
		expMd5 := md5Byte(append(str, keyb...), 16)
		str = []byte(fmt.Sprintf("%010d%s%s", expiry, expMd5, str))
		//str = append([]byte(fmt.Sprintf("%010d", expiry)), append(expMd5, str...)...)
	}

	strLen = len(str)
	resdata := make([]byte, 0, strLen)
	var rndkey, box [256]int
	// 产生密钥簿
	h := 0
	i := 0
	j := 0

	for i = 0; i < 256; i++ {
		rndkey[i] = int(cryptkey[i%cryptkeyLen])
		box[i] = i
	}
	// 用固定的算法，打乱密钥簿，增加随机性，好像很复杂，实际上并不会增加密文的强度
	for i = 0; i < 256; i++ {
		j = (j + box[i] + rndkey[i]) % 256
		box[i], box[j] = box[j], box[i]
	}
	// 核心加解密部分
	h = 0
	j = 0
	for i = 0; i < strLen; i++ {
		h = ((h + 1) % 256)
		j = ((j + box[h]) % 256)
		box[h], box[j] = box[j], box[h]
		// 从密钥簿得出密钥进行异或，再转成字符
		resdata = append(resdata, byte(int(str[i])^box[(box[h]+box[j])%256]))
	}
	if encode == false { //解密
		// substr($result, 0, 10) == 0 验证数据有效性
		// substr($result, 0, 10) - time() > 0 验证数据有效性
		// substr($result, 10, 16) == substr(md5(substr($result, 26).$keyb), 0, 16) 验证数据完整性
		// 验证数据有效性，请看未加密明文的格式
		if len(resdata) <= 26 {
			return nil, 0
		}

		expTime, _ := strconv.ParseInt(string(resdata[:10]), 10, 0)
		if (expTime == 0 || expTime-time.Now().Unix() > 0) && string(resdata[10:26]) == string(md5Byte(append(resdata[26:], keyb...), 16)) {
			return resdata[26:], expTime
		} else {
			return nil, expTime
		}
	} else { //加密
		// 把动态密钥保存在密文里，这也是为什么同样的明文，生产不同密文后能解密的原因
		resdata = append(keyc, ke.Base64UrlEncode(resdata)...)
		return resdata, expiry
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

	res, err := bcrypt.GenerateFromPassword(password, cost)
	return res, err
}
