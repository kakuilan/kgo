package kgo

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Base64Encode 使用 MIME base64 对数据进行编码
func (ke *LkkEncrypt) Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// Base64Decode 对使用 MIME base64 编码的数据进行解码
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

// Base64UrlSafeEncode url安全的Base64Encode,没有'/'和'+'及结尾的'='
func (ke *LkkEncrypt) Base64UrlEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}

// Base64UrlDecode url安全的Base64Decode
func (ke *LkkEncrypt) Base64UrlDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	return base64.URLEncoding.DecodeString(data)
}

// AuthCode 授权码编码或解码;encode为true时编码,为false解码;expiry为有效期,秒;
func (ke *LkkEncrypt) AuthCode(str, key string, encode bool, expiry int64) string {
	if str == "" {
		return ""
	}

	// 动态密钥长度，相同的明文会生成不同密文就是依靠动态密钥
	// 加入随机密钥，可以令密文无任何规律，即便是原文和密钥完全相同，加密结果也会每次不同，增大破解难度。
	// 取值越大，密文变动规律越大，密文变化 = 16 的 ckeyLength 次方
	// 当此值为 0 时，则不产生随机密钥
	ckeyLength := 4

	// 密钥
	keyByte := md5Str([]byte(key), 32)

	// 密钥a会参与加解密
	keya := md5Str(keyByte[:16], 32)

	// 密钥b会用来做数据完整性验证
	keyb := md5Str(keyByte[16:], 32)

	// 密钥c用于变化生成的密文
	var keyc []byte
	if encode == false {
		keyc = []byte(str[:ckeyLength])
	} else {
		cLen := 32 - ckeyLength
		now, _ := time.Now().MarshalBinary()
		timeBytes := md5Str(now, 32)
		keyc = timeBytes[cLen:]
	}

	// 参与运算的密钥
	keyd := md5Str(append(keya, keyc...), 32)
	cryptkey := append(keya, keyd...)
	keyLength := len(cryptkey)
	// 明文，前10位用来保存时间戳，解密时验证数据有效性，10到26位用来保存keyb(密钥b)，解密时会通过这个密钥验证数据完整性
	// 如果是解码的话，会从第ckeyLength位开始，因为密文前ckeyLength位保存 动态密钥，以保证解密正确
	if encode == false {
		strByte, err := base64.StdEncoding.DecodeString(str[ckeyLength:])
		if err != nil {
			println(err.Error())
			return ""
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
	tmp := 0
	for i = 0; i < 256; i++ {
		rndkey[i] = int(cryptkey[i%keyLength])
		box[i] = i
	}
	// 用固定的算法，打乱密钥簿，增加随机性，好像很复杂，实际上并不会增加密文的强度
	for i = 0; i < 256; i++ {
		j = (j + box[i] + rndkey[i]) % 256
		tmp = box[i]
		box[i] = box[j]
		box[j] = tmp
	}
	// 核心加解密部分
	a = 0
	j = 0
	tmp = 0
	for i = 0; i < stringLength; i++ {
		a = ((a + 1) % 256)
		j = ((j + box[a]) % 256)
		tmp = box[a]
		box[a] = box[j]
		box[j] = tmp
		// 从密钥簿得出密钥进行异或，再转成字符
		resdata = append(resdata, byte(int(str[i])^box[(box[a]+box[j])%256]))
	}
	result := string(resdata)
	if encode == false {
		// substr($result, 0, 10) == 0 验证数据有效性
		// substr($result, 0, 10) - time() > 0 验证数据有效性
		// substr($result, 10, 16) == substr(md5(substr($result, 26).$keyb), 0, 16) 验证数据完整性
		// 验证数据有效性，请看未加密明文的格式
		frontTen, _ := strconv.ParseInt(result[:10], 10, 0)
		if (frontTen == 0 || frontTen-time.Now().Unix() > 0) && result[10:26] == string(md5Str(append(resdata[26:], keyb...), 16)) {
			return result[26:]
		} else {
			return ""
		}
	} else {
		// 把动态密钥保存在密文里，这也是为什么同样的明文，生产不同密文后能解密的原因
		result = string(keyc) + base64.StdEncoding.EncodeToString([]byte(result))
		return result
	}
}
