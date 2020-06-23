package kgo

import (
	"bytes"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	xhtml "golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
	"hash/crc32"
	"html"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Nl2br 将换行符转换为br标签.
func (ks *LkkString) Nl2br(html string) string {
	if html == "" {
		return ""
	}
	return strings.Replace(html, "\n", "<br />", -1)
}

// Br2nl 将br标签转换为换行符.
func (ks *LkkString) Br2nl(str string) string {
	// <br> , <br /> , <br/>
	// <BR> , <BR /> , <BR/>
	nlchar := []byte("\n")

	l := len(str)
	buf := make([]byte, 0, l) //prealloca

	for i := 0; i < l; i++ {
		switch str[i] {
		case 60: //<
			if l >= i+3 {
				/*
					b = 98
					B = 66
					r = 82
					R = 114
					SPACE = 32
					/ = 47
					> = 62
				*/

				if l >= i+3 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 62) { // <br> || <BR>
					buf = append(buf, nlchar...)
					i += 3
					continue
				}

				if l >= i+4 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 47 && str[i+4] == 62) { // <br/> || <BR/>
					buf = append(buf, nlchar...)
					i += 4
					continue
				}

				if l >= i+5 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 32 && str[i+4] == 47 && str[i+5] == 62) { // <br /> || <BR />
					buf = append(buf, nlchar...)
					i += 5
					continue
				}
			}
		default:
			buf = append(buf, str[i])
		}
	}

	return string(buf)
}

// RemoveSpace 移除字符串中的空白字符.
// all为true时移除全部空白,为false时只替换连续的空白字符为一个空格.
func (ks *LkkString) RemoveSpace(str string, all bool) string {
	if all && str != "" {
		return strings.Join(strings.Fields(str), "")
	} else if str != "" {
		//先将2个以上的连续空白符转为空格
		str = RegWhitespaceDuplicate.ReplaceAllString(str, " ")
		//再将[\t\n\f\r]等转为空格
		str = RegWhitespace.ReplaceAllString(str, " ")
	}

	return strings.TrimSpace(str)
}

// StripTags 过滤html和php标签.
func (ks *LkkString) StripTags(str string) string {
	if str == "" {
		return ""
	}

	return RegHtmlTag.ReplaceAllString(str, "")
}

// Html2Text 将html转换为纯文本.
func (ks *LkkString) Html2Text(str string) string {
	if str == "" {
		return ""
	}

	domDoc := xhtml.NewTokenizer(strings.NewReader(str))
	previousStartToken := domDoc.Token()
	var text string
loopDom:
	for {
		nx := domDoc.Next()
		switch {
		case nx == xhtml.ErrorToken:
			break loopDom // End of the document
		case nx == xhtml.StartTagToken:
			previousStartToken = domDoc.Token()
		case nx == xhtml.TextToken:
			if chk, _ := ks.Dstrpos(previousStartToken.Data, TextHtmlExcludeTags, false); chk {
				continue
			}

			text += " " + strings.TrimSpace(xhtml.UnescapeString(string(domDoc.Text())))
		}
	}

	return ks.RemoveSpace(text, false)
}

// Md5 获取字符串md5值,length指定结果长度32/16 .
func (ks *LkkString) Md5(str string, length uint8) string {
	return string(md5Str([]byte(str), length))
}

// ShaX 计算字符串的 shaX 散列值,x为1/256/512 .
func (ks *LkkString) ShaX(str string, x uint16) string {
	return string(shaXStr([]byte(str), x))
}

// Random 生成随机字符串.
// length为长度,rtype为枚举:
// RAND_STRING_ALPHA 字母;
// RAND_STRING_NUMERIC 数值;
// RAND_STRING_ALPHANUM 字母+数值;
// RAND_STRING_SPECIAL 字母+数值+特殊字符;
// RAND_STRING_CHINESE 仅中文.
func (ks *LkkString) Random(length uint8, rtype LkkRandString) string {
	if length == 0 {
		return ""
	}

	var letter []rune
	alphas := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specials := "~!@#$%^&*()_+{}:|<>?`-=;,."

	rand.Seed(time.Now().UTC().UnixNano())

	switch rtype {
	case RAND_STRING_ALPHA:
		letter = []rune(alphas)
	case RAND_STRING_NUMERIC:
		letter = []rune(numbers)
	case RAND_STRING_ALPHANUM:
		letter = []rune(alphas + numbers)
	case RAND_STRING_SPECIAL:
		letter = []rune(alphas + numbers + specials)
	case RAND_STRING_CHINESE:
		chineses := "们以我到他会作时要动国产的一是工就年阶义发成部民可出能方进在了不和有大这主中人上为来分生对于学下级地个用同行面说种过命度革而多子后自社加小机也经力线本电高量长党得实家定深法表着水理化争现所二起政三好十战无农使性前等反体合斗路图把结第里正新开论之物从当两些还天资事队批点育重其思与间内去因件日利相由压员气业代全组数果期导平各基或月毛然如应形想制心样干都向变关问比展那它最及外没看治提五解系林者米群头意只明四道马认次文通但条较克又公孔领军流入接席位情运器并飞原油放立题质指建区验活众很教决特此常石强极土少已根共直团统式转别造切九你取西持总料连任志观调七么山程百报更见必真保热委手改管处己将修支识病象几先老光专什六型具示复安带每东增则完风回南广劳轮科北打积车计给节做务被整联步类集号列温装即毫知轴研单色坚据速防史拉世设达尔场织历花受求传口断况采精金界品判参层止边清至万确究书术状厂须离再目海交权且儿青才证低越际八试规斯近注办布门铁需走议县兵固除般引齿千胜细影济白格效置推空配刀叶率述今选养德话查差半敌始片施响收华觉备名红续均药标记难存测士身紧液派准斤角降维板许破述技消底床田势端感往神便贺村构照容非搞亚磨族火段算适讲按值美态黄易彪服早班麦削信排台声该击素张密害侯草何树肥继右属市严径螺检左页抗苏显苦英快称坏移约巴材省黑武培著河帝仅针怎植京助升王眼她抓含苗副杂普谈围食射源例致酸旧却充足短划剂宣环落首尺波承粉践府鱼随考刻靠够满夫失包住促枝局菌杆周护岩师举曲春元超负砂封换太模贫减阳扬江析亩木言球朝医校古呢稻宋听唯输滑站另卫字鼓刚写刘微略范供阿块某功套友限项余倒卷创律雨让骨远帮初皮播优占死毒圈伟季训控激找叫云互跟裂粮粒母练塞钢顶策双留误础吸阻故寸盾晚丝女散焊功株亲院冷彻弹错散商视艺灭版烈零室轻血倍缺厘泵察绝富城冲喷壤简否柱李望盘磁雄似困巩益洲脱投送奴侧润盖挥距触星松送获兴独官混纪依未突架宽冬章湿偏纹吃执阀矿寨责熟稳夺硬价努翻奇甲预职评读背协损棉侵灰虽矛厚罗泥辟告卵箱掌氧恩爱停曾溶营终纲孟钱待尽俄缩沙退陈讨奋械载胞幼哪剥迫旋征槽倒握担仍呀鲜吧卡粗介钻逐弱脚怕盐末阴丰雾冠丙街莱贝辐肠付吉渗瑞惊顿挤秒悬姆烂森糖圣凹陶词迟蚕亿矩康遵牧遭幅园腔订香肉弟屋敏恢忘编印蜂急拿扩伤飞露核缘游振操央伍域甚迅辉异序免纸夜乡久隶缸夹念兰映沟乙吗儒杀汽磷艰晶插埃燃欢铁补咱芽永瓦倾阵碳演威附牙芽永瓦斜灌欧献顺猪洋腐请透司危括脉宜笑若尾束壮暴企菜穗楚汉愈绿拖牛份染既秋遍锻玉夏疗尖殖井费州访吹荣铜沿替滚客召旱悟刺脑措贯藏敢令隙炉壳硫煤迎铸粘探临薄旬善福纵择礼愿伏残雷延烟句纯渐耕跑泽慢栽鲁赤繁境潮横掉锥希池败船假亮谓托伙哲怀割摆贡呈劲财仪沉炼麻罪祖息车穿货销齐鼠抽画饲龙库守筑房歌寒喜哥洗蚀废纳腹乎录镜妇恶脂庄擦险赞钟摇典柄辩竹谷卖乱虚桥奥伯赶垂途额壁网截野遗静谋弄挂课镇妄盛耐援扎虑键归符庆聚绕摩忙舞遇索顾胶羊湖钉仁音迹碎伸灯避泛亡答勇频皇柳哈揭甘诺概宪浓岛袭谁洪谢炮浇斑讯懂灵蛋闭孩释乳巨徒私银伊景坦累匀霉杜乐勒隔弯绩招绍胡呼痛峰零柴簧午跳居尚丁秦稍追梁折耗碱殊岗挖氏刃剧堆赫荷胸衡勤膜篇登驻案刊秧缓凸役剪川雪链渔啦脸户洛孢勃盟买杨宗焦赛旗滤硅炭股坐蒸凝竟陷枪黎救冒暗洞犯筒您宋弧爆谬涂味津臂障褐陆啊健尊豆拔莫抵桑坡缝警挑污冰柬嘴啥饭塑寄赵喊垫丹渡耳刨虎笔稀昆浪萨茶滴浅拥穴覆伦娘吨浸袖珠雌妈紫戏塔锤震岁貌洁剖牢锋疑霸闪埔猛诉刷狠忽灾闹乔唐漏闻沈熔氯荒茎男凡抢像浆旁玻亦忠唱蒙予纷捕锁尤乘乌智淡允叛畜俘摸锈扫毕璃宝芯爷鉴秘净蒋钙肩腾枯抛轨堂拌爸循诱祝励肯酒绳穷塘燥泡袋朗喂铝软渠颗惯贸粪综墙趋彼届墨碍启逆卸航衣孙龄岭骗休借"
		letter = []rune(chineses)
	default:
		letter = []rune(alphas)
	}

	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

// Index 查找子串sub在字符串str中第一次出现的位置,不存在则返回-1;
// ignoreCase为是否忽略大小写.
func (ks *LkkString) Index(str, sub string, ignoreCase bool) int {
	if str == "" || sub == "" {
		return -1
	}

	if ignoreCase {
		str = strings.ToLower(str)
		sub = strings.ToLower(sub)
	}

	return strings.Index(str, sub)
}

// LastIndex 查找子串sub在字符串str中最后一次出现的位置,不存在则返回-1;
// ignoreCase为是否忽略大小写.
func (ks *LkkString) LastIndex(str, sub string, ignoreCase bool) int {
	if str == "" || sub == "" {
		return -1
	}

	if ignoreCase {
		str = strings.ToLower(str)
		sub = strings.ToLower(sub)
	}

	return strings.LastIndex(str, sub)
}

// Strpos 查找字符串首次出现的位置,找不到时返回-1.
// haystack在该字符串中进行查找,needle要查找的字符串;
// offset起始位置,为负数时时,搜索会从字符串结尾指定字符数开始.
func (ks *LkkString) Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Stripos  查找字符串首次出现的位置（不区分大小写）,找不到时返回-1.
// haystack在该字符串中进行查找,needle要查找的字符串;
// offset起始位置,为负数时时,搜索会从字符串结尾指定字符数开始.
func (ks *LkkString) Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := ks.Index(haystack[offset:], needle, true)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Strrpos 查找指定字符串在目标字符串中最后一次出现的位置.
func (ks *LkkString) Strrpos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(haystack, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// Strripos 查找指定字符串在目标字符串中最后一次出现的位置（不区分大小写）.
func (ks *LkkString) Strripos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = ks.LastIndex(haystack, needle, true)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// Ucfirst 将字符串的第一个字符转换为大写.
func (ks *LkkString) Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

// Lcfirst 将字符串的第一个字符转换为小写.
func (ks *LkkString) Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

// Ucwords 将字符串中每个词的首字母转换为大写.
func (ks *LkkString) Ucwords(str string) string {
	return strings.Title(str)
}

// Lcwords 将字符串中每个词的首字母转换为小写.
func (ks *LkkString) Lcwords(str string) string {
	buf := &bytes.Buffer{}
	lastIsSpace := true
	for _, r := range str {
		if unicode.IsLetter(r) {
			if lastIsSpace {
				r = unicode.ToLower(r)
			}

			lastIsSpace = false
		} else {
			lastIsSpace = false
			if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
				lastIsSpace = true
			}
		}

		buf.WriteRune(r)
	}

	return buf.String()
}

// Substr 截取字符串str的子串.
// start 为起始位置.若值是负数,返回的结果将从 str 结尾处向前数第 abs(start) 个字符开始.
// length 为截取的长度.若值时负数, str 末尾处的 abs(length) 个字符将会被省略.
// start/length的绝对值必须<=原字符串长度.
func (ks *LkkString) Substr(str string, start int, length ...int) string {
	total := len(str)
	if total == 0 {
		return ""
	}

	var sublen, end int
	max := total //最大的结束位置

	if len(length) == 0 {
		sublen = total
	} else {
		sublen = length[0]
	}

	if start < 0 {
		start = total + start
	}

	if sublen < 0 {
		sublen = total + sublen
		if sublen > 0 {
			max = sublen
		}
	}

	if start < 0 || sublen <= 0 || start >= max {
		return ""
	}

	end = start + sublen
	if end > max {
		end = max
	}

	return str[start:end]
}

// MbSubstr 返回(宽字符)字符串str的子串.
// start 为起始位置.若值是负数,返回的结果将从 str 结尾处向前数第 abs(start) 个字符开始.
// length 为截取的长度.若值时负数, str 末尾处的 abs(length) 个字符将会被省略.
// start/length的绝对值必须<=原字符串长度.
func (ks *LkkString) MbSubstr(str string, start int, length ...int) string {
	if len(str) == 0 {
		return ""
	}

	runes := []rune(str)
	total := len(runes)

	var sublen, end int
	max := total //最大的结束位置

	if len(length) == 0 {
		sublen = total
	} else {
		sublen = length[0]
	}

	if start < 0 {
		start = total + start
	}

	if sublen < 0 {
		sublen = total + sublen
		if sublen > 0 {
			max = sublen
		}
	}

	if start < 0 || sublen <= 0 || start >= max {
		return ""
	}

	end = start + sublen
	if end > max {
		end = max
	}

	return string(runes[start:end])
}

// SubstrCount 计算字符串出现的次数.
func (ks *LkkString) SubstrCount(str, substr string) int {
	return strings.Count(str, substr)
}

// Reverse 反转字符串.
func (ks *LkkString) Reverse(str string) string {
	n := len(str)
	runes := make([]rune, n)
	for _, r := range str {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}

// ChunkSplit 将字符串分割成小块.str为要分割的字符,chunklen为分割的尺寸,end为行尾序列符号.
func (ks *LkkString) ChunkSplit(str string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(str), []rune(end)
	length := uint(len(runes))
	if length <= 1 || length < chunklen {
		return str + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < length; i += chunklen {
		if i+chunklen > length {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

// Strlen 获取字符串长度.
func (ks *LkkString) Strlen(str string) int {
	return len(str)
}

// MbStrlen 获取宽字符串的长度,多字节的字符被计为 1.
func (ks *LkkString) MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

// Shuffle 随机打乱字符串.
func (ks *LkkString) Shuffle(str string) string {
	if str == "" {
		return str
	}

	runes := []rune(str)
	index := 0

	for i := len(runes) - 1; i > 0; i-- {
		index = rand.Intn(i + 1)

		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}

	return string(runes)
}

// Trim 去除字符串首尾处的空白字符（或者其他字符）.
func (ks *LkkString) Trim(str string, characterMask ...string) string {
	mask := getTrimMask(characterMask)
	return strings.Trim(str, mask)
}

// Ltrim 删除字符串开头的空白字符（或其他字符）.
func (ks *LkkString) Ltrim(str string, characterMask ...string) string {
	mask := getTrimMask(characterMask)
	return strings.TrimLeft(str, mask)
}

// Rtrim 删除字符串末端的空白字符（或者其他字符）.
func (ks *LkkString) Rtrim(str string, characterMask ...string) string {
	mask := getTrimMask(characterMask)
	return strings.TrimRight(str, mask)
}

// Chr 返回相对应于 ascii 所指定的单个字符.
func (ks *LkkString) Chr(ascii int) string {
	return string(ascii)
}

// Ord 将首字符转换为rune.
func (ks *LkkString) Ord(char string) rune {
	r, _ := utf8.DecodeRune([]byte(char))
	return r
}

// JsonEncode 对变量进行 JSON 编码.
// 依赖库github.com/json-iterator/go.
func (ks *LkkString) JsonEncode(val interface{}) ([]byte, error) {
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	return jsons.Marshal(val)
}

// JsonDecode 对 JSON 格式的字符串进行解码,注意val使用指针.
// 依赖库github.com/json-iterator/go.
func (ks *LkkString) JsonDecode(data []byte, val interface{}) error {
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	return jsons.Unmarshal(data, val)
}

// Addslashes 使用反斜线引用字符串.
func (ks *LkkString) Addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Stripslashes 反引用一个引用字符串.
func (ks *LkkString) Stripslashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < l && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Quotemeta 转义元字符集,包括 . \ + * ? [ ^ ] ( $ ).
func (ks *LkkString) Quotemeta(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Htmlentities 将字符转换为 HTML 转义字符.
func (ks *LkkString) Htmlentities(str string) string {
	return html.EscapeString(str)
}

// HtmlentityDecode 将HTML实体转换为它们对应的字符.
func (ks *LkkString) HtmlentityDecode(str string) string {
	return html.UnescapeString(str)
}

// Crc32 计算一个字符串的 crc32 多项式.
func (ks *LkkString) Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

// SimilarText 计算两个字符串的相似度,返回在两个字符串中匹配字符的数目;percent为相似程度百分数.
func (ks *LkkString) SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

// Explode 字符串分割.delimiters为分隔符,可选,支持多个.
func (ks *LkkString) Explode(str string, delimiters ...string) (res []string) {
	if str == "" {
		return
	}

	dLen := len(delimiters)
	if dLen == 0 {
		res = append(res, str)
	} else if dLen > 1 {
		var sl []string
		for _, v := range delimiters {
			if v != "" {
				sl = append(sl, v, KDelimiter)
			}
		}
		str = strings.NewReplacer(sl...).Replace(str)
		res = strings.Split(str, KDelimiter)
	} else {
		res = strings.Split(str, delimiters[0])
	}

	return
}

// Uniqid 获取一个带前缀、基于当前时间微秒数的唯一ID.
func (ks *LkkString) Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}

// UuidV4 获取UUID(Version4).
func (ks *LkkString) UuidV4() (string, error) {
	buf := make([]byte, 16)
	_, err := crand.Read(buf)

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16]), err
}

// VersionCompare 对比两个版本号字符串.
// 在第一个版本低于第二个时,version_compare() 返回 -1;如果两者相等,返回 0;第二个版本更低时则返回 1.
// operator允许的操作符有: <, lt, <=, le, >, gt, >=, ge, ==, =, eq, !=, <>, ne .
// 特定的版本字符串，将会用以下顺序处理：
// 未找到的任意字符串 < dev < alpha = a < beta = b < RC = rc < # < pl = p
// 用法:
// VersionCompare("1.2.3-alpha", "1.2.3RC7", '>=') ;
// VersionCompare("1.2.3-beta", "1.2.3pl", 'lt') ;
// VersionCompare("1.1_dev", "1.2any", 'eq') .
func (ks *LkkString) VersionCompare(version1, version2, operator string) bool {
	var vcompare func(string, string) int
	var canonicalize func(string) string
	var special func(string, string) int

	// version compare
	vcompare = func(origV1, origV2 string) int {
		if origV1 == "" || origV2 == "" {
			if origV1 == "" && origV2 == "" {
				return 0
			}
			if origV1 == "" {
				return -1
			}
			return 1
		}

		ver1, ver2, compare := "", "", 0
		if origV1[0] == '#' {
			ver1 = origV1
		} else {
			ver1 = canonicalize(origV1)
		}
		if origV2[0] == '#' {
			ver2 = origV2
		} else {
			ver2 = canonicalize(origV2)
		}
		n1, n2 := 0, 0
		for {
			p1, p2 := "", ""
			n1 = strings.IndexByte(ver1, '.')
			if n1 == -1 {
				p1, ver1 = ver1[:], ""
			} else {
				p1, ver1 = ver1[:n1], ver1[n1+1:]
			}
			n2 = strings.IndexByte(ver2, '.')
			if n2 == -1 {
				p2, ver2 = ver2, ""
			} else {
				p2, ver2 = ver2[:n2], ver2[n2+1:]
			}

			if p1 == "" || p2 == "" {
				break
			}

			if (p1[0] >= '0' && p1[0] <= '9') && (p2[0] >= '0' && p2[0] <= '9') { // all is digit
				l1, _ := strconv.Atoi(p1)
				l2, _ := strconv.Atoi(p2)
				if l1 > l2 {
					compare = 1
				} else if l1 == l2 {
					compare = 0
				} else {
					compare = -1
				}
			} else if !(p1[0] >= '0' && p1[0] <= '9') && !(p2[0] >= '0' && p2[0] <= '9') { // all not digit
				compare = special(p1, p2)
			} else { // part is digit
				if p1[0] >= '0' && p1[0] <= '9' { // is digit
					compare = special("#N#", p2)
				} else {
					compare = special(p1, "#N#")
				}
			}

			if compare != 0 || n1 == -1 || n2 == -1 {
				break
			}
		}

		if compare == 0 {
			if ver1 != "" {
				if ver1[0] >= '0' && ver1[0] <= '9' {
					compare = 1
				} else {
					compare = vcompare(ver1, "#N#")
				}
			} else if ver2 != "" {
				if ver2[0] >= '0' && ver2[0] <= '9' {
					compare = -1
				} else {
					compare = vcompare("#N#", ver2)
				}
			}
		}

		return compare
	}

	// canonicalize 规范化转换
	canonicalize = func(version string) string {
		ver := []byte(version)
		l := len(ver)
		var buf = make([]byte, l*2)
		j := 0
		for i, v := range ver {
			next := uint8(0)
			if i+1 < l { // Have the next one
				next = ver[i+1]
			}
			if v == '-' || v == '_' || v == '+' { // replace '-', '_', '+' to '.'
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else if (next > 0) &&
				(!(next >= '0' && next <= '9') && (v >= '0' && v <= '9')) ||
				(!(v >= '0' && v <= '9') && (next >= '0' && next <= '9')) { // Insert '.' before and after a non-digit
				buf[j] = v
				j++
				if v != '.' && next != '.' {
					buf[j] = '.'
					j++
				}
				continue
			} else if !((v >= '0' && v <= '9') ||
				(v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')) { // Non-letters and numbers
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else {
				buf[j] = v
				j++
			}
		}

		return string(buf[:j])
	}

	// compare special version forms 特殊版本号
	special = func(form1, form2 string) int {
		found1, found2, len1, len2 := -1, -1, len(form1), len(form2)
		// (Any string not found) < dev < alpha = a < beta = b < RC = rc < # < pl = p
		forms := map[string]int{
			"dev":   0,
			"alpha": 1,
			"a":     1,
			"beta":  2,
			"b":     2,
			"RC":    3,
			"rc":    3,
			"#":     4,
			"pl":    5,
			"p":     5,
		}

		for name, order := range forms {
			if len1 < len(name) {
				continue
			}
			if strings.Compare(form1[:len(name)], name) == 0 {
				found1 = order
				break
			}
		}
		for name, order := range forms {
			if len2 < len(name) {
				continue
			}
			if strings.Compare(form2[:len(name)], name) == 0 {
				found2 = order
				break
			}
		}

		if found1 == found2 {
			return 0
		} else if found1 > found2 {
			return 1
		} else {
			return -1
		}
	}

	compare := vcompare(version1, version2)

	switch operator {
	case "<", "lt":
		return compare == -1
	case "<=", "le":
		return compare != 1
	case ">", "gt":
		return compare == 1
	case ">=", "ge":
		return compare != -1
	case "==", "=", "eq":
		return compare == 0
	case "!=", "<>", "ne":
		return compare != 0
	default:
		panic("[VersionCompare] operator: invalid")
	}
}

// ToCamelCase 转为驼峰写法.
// 去掉包括下划线"_"和横杠"-".
func (ks *LkkString) ToCamelCase(str string) string {
	if len(str) == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	var r0, r1 rune
	var size int

	// leading connector will appear in output.
	for len(str) > 0 {
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if !isCaseConnector(r0) {
			r0 = unicode.ToUpper(r0)
			break
		}

		buf.WriteRune(r0)
	}

	for len(str) > 0 {
		r1 = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if isCaseConnector(r0) && isCaseConnector(r1) {
			buf.WriteRune(r1)
			continue
		}

		if isCaseConnector(r1) {
			r0 = unicode.ToUpper(r0)
		} else if unicode.IsLower(r1) && unicode.IsUpper(r0) {
			buf.WriteRune(r1)
		} else if unicode.IsUpper(r1) && unicode.IsLower(r0) {
			buf.WriteRune(r1)
		} else {
			r0 = unicode.ToLower(r0)
			buf.WriteRune(r1)
		}
	}

	buf.WriteRune(r0)
	return buf.String()
}

// ToSnakeCase 转为蛇形写法.
// 使用下划线"_"连接.
func (ks *LkkString) ToSnakeCase(str string) string {
	return camelCaseToLowerCase(str, '_')
}

// ToSnakeCase 转为串形写法.
// 使用横杠"-"连接.
func (ks *LkkString) ToKebabCase(str string) string {
	return camelCaseToLowerCase(str, '-')
}

// RemoveBefore 移除before之前的字符串;
// include为是否移除包括before本身;
// ignoreCase为是否忽略大小写.
func (ks *LkkString) RemoveBefore(str, before string, include, ignoreCase bool) string {
	i := ks.Index(str, before, ignoreCase)
	if i > 0 {
		if include {
			str = str[i+len(before):]
		} else {
			str = str[i:]
		}
	}
	return str
}

// RemoveAfter 移除after之后的字符串;
// include为是否移除包括after本身;
// ignoreCase为是否忽略大小写.
func (ks *LkkString) RemoveAfter(str, after string, include, ignoreCase bool) string {
	i := ks.Index(str, after, ignoreCase)
	if i > 0 {
		if include {
			str = str[0:i]
		} else {
			str = str[0 : i+len(after)]
		}
	}
	return str
}

// DBC2SBC 半角转全角.
func (ks *LkkString) DBC2SBC(s string) string {
	return width.Widen.String(s)
}

// SBC2DBC 全角转半角.
func (ks *LkkString) SBC2DBC(s string) string {
	return width.Narrow.String(s)
}

// Levenshtein 计算两个字符串之间的编辑距离.
func (ks *LkkString) Levenshtein(a, b *string) int {
	la := len(*a)
	lb := len(*b)

	if *a == *b {
		return 0
	} else if la > 255 || lb > 255 {
		return -1
	}

	d := make([]int, la+1)
	var lastdiag, olddiag, temp int
	for i := 1; i <= la; i++ {
		d[i] = i
	}
	for i := 1; i <= lb; i++ {
		d[0] = i
		lastdiag = i - 1
		for j := 1; j <= la; j++ {
			olddiag = d[j]
			min := d[j] + 1
			if (d[j-1] + 1) < min {
				min = d[j-1] + 1
			}
			if (*a)[j-1] == (*b)[i-1] {
				temp = 0
			} else {
				temp = 1
			}
			if (lastdiag + temp) < min {
				min = lastdiag + temp
			}
			d[j] = min
			lastdiag = olddiag
		}
	}
	return d[la]
}

// ClosestWord 获取与原字符串相似度最高的字符串,以及它们的编辑距离.
// word为原字符串,searchs为待查找的字符串数组.
func (ks *LkkString) ClosestWord(word string, searchs []string) (string, int) {
	distance := 10000
	res := ""
	for _, search := range searchs {
		newVal := ks.Levenshtein(&word, &search)
		if newVal == 0 {
			distance = 0
			res = search
			break
		}

		if newVal < distance {
			distance = newVal
			res = search
		}
	}

	return res, distance
}

// Utf8ToGbk UTF-8转GBK编码.
func (ks *LkkString) Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	return d, e
}

// GbkToUtf8 GBK转UTF-8编码.
func (ks *LkkString) GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	return d, e
}

// Utf8ToBig5 UTF-8转BIG5编码.
func (ks *LkkString) Utf8ToBig5(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	return d, e
}

// Big5ToUtf8 BIG5转UTF-8编码.
func (ks *LkkString) Big5ToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	return d, e
}

// FirstLetter 获取字符串首字母.
func (ks *LkkString) FirstLetter(str string) string {
	if str != "" {
		// 获取字符串第一个字符
		_, size := utf8.DecodeRuneInString(str)
		firstChar := str[:size]

		if ks.IsLetters(firstChar) {
			return firstChar
		} else if ks.IsChinese(firstChar) {
			// Utf8 转 GBK2312
			firstCharGbk, _ := ks.Utf8ToGbk([]byte(firstChar))

			// 获取第一个字符的16进制
			firstCharHex := hex.EncodeToString(firstCharGbk)

			// 16进制转十进制
			firstCharDec, _ := strconv.ParseInt(firstCharHex, 16, 0)

			// 十进制落在GB 2312的某个拼音区间即为某个字母
			firstCharDecimalRelative := firstCharDec - 65536
			if firstCharDecimalRelative >= -20319 && firstCharDecimalRelative <= -20284 {
				return "A"
			}
			if firstCharDecimalRelative >= -20283 && firstCharDecimalRelative <= -19776 {
				return "B"
			}
			if firstCharDecimalRelative >= -19775 && firstCharDecimalRelative <= -19219 {
				return "C"
			}
			if firstCharDecimalRelative >= -19218 && firstCharDecimalRelative <= -18711 {
				return "D"
			}
			if firstCharDecimalRelative >= -18710 && firstCharDecimalRelative <= -18527 {
				return "E"
			}
			if firstCharDecimalRelative >= -18526 && firstCharDecimalRelative <= -18240 {
				return "F"
			}
			if firstCharDecimalRelative >= -18239 && firstCharDecimalRelative <= -17923 {
				return "G"
			}
			if firstCharDecimalRelative >= -17922 && firstCharDecimalRelative <= -17418 {
				return "H"
			}
			if firstCharDecimalRelative >= -17417 && firstCharDecimalRelative <= -16475 {
				return "J"
			}
			if firstCharDecimalRelative >= -16474 && firstCharDecimalRelative <= -16213 {
				return "K"
			}
			if firstCharDecimalRelative >= -16212 && firstCharDecimalRelative <= -15641 {
				return "L"
			}
			if firstCharDecimalRelative >= -15640 && firstCharDecimalRelative <= -15166 {
				return "M"
			}
			if firstCharDecimalRelative >= -15165 && firstCharDecimalRelative <= -14923 {
				return "N"
			}
			if firstCharDecimalRelative >= -14922 && firstCharDecimalRelative <= -14915 {
				return "O"
			}
			if firstCharDecimalRelative >= -14914 && firstCharDecimalRelative <= -14631 {
				return "P"
			}
			if firstCharDecimalRelative >= -14630 && firstCharDecimalRelative <= -14150 {
				return "Q"
			}
			if firstCharDecimalRelative >= -14149 && firstCharDecimalRelative <= -14091 {
				return "R"
			}
			if firstCharDecimalRelative >= -14090 && firstCharDecimalRelative <= -13319 {
				return "S"
			}
			if firstCharDecimalRelative >= -13318 && firstCharDecimalRelative <= -12839 {
				return "T"
			}
			if firstCharDecimalRelative >= -12838 && firstCharDecimalRelative <= -12557 {
				return "W"
			}
			if firstCharDecimalRelative >= -12556 && firstCharDecimalRelative <= -11848 {
				return "X"
			}
			if firstCharDecimalRelative >= -11847 && firstCharDecimalRelative <= -11056 {
				return "Y"
			}
			if firstCharDecimalRelative >= -11055 && firstCharDecimalRelative <= -10247 {
				return "Z"
			}
		}
	}

	return ""
}

// Dstrpos 检查字符串str是否包含数组arr的元素之一,返回检查结果和匹配的字符串.
// chkCase为是否检查大小写.
func (ks *LkkString) Dstrpos(str string, arr []string, chkCase bool) (bool, string) {
	if len(str) == 0 || len(arr) == 0 {
		return false, ""
	}

	for _, v := range arr {
		if (chkCase && ks.Strpos(str, v, 0) != -1) || (!chkCase && ks.Stripos(str, v, 0) != -1) {
			return true, v
		}
	}

	return false, ""
}

// HideCard 隐藏证件号码.
func (ks *LkkString) HideCard(card string) string {
	res := "******"
	leng := len(card)
	if leng > 4 && leng <= 10 {
		res = card[0:4] + "******"
	} else if leng > 10 {
		res = card[0:4] + "******" + card[(leng-4):leng]
	}

	return res
}

// HideMobile 隐藏手机号.
func (ks *LkkString) HideMobile(mobile string) string {
	res := "***"
	leng := len(mobile)
	if leng > 7 {
		res = mobile[0:3] + "****" + mobile[leng-3:leng]
	}

	return res
}

// HideTrueName 隐藏真实名称(如姓名、账号、公司等).
func (ks *LkkString) HideTrueName(name string) string {
	res := "**"
	if name != "" {
		runs := []rune(name)
		leng := len(runs)
		if leng <= 3 {
			res = string(runs[0:1]) + res
		} else if leng < 5 {
			res = string(runs[0:2]) + res
		} else if leng < 10 {
			res = string(runs[0:2]) + "***" + string(runs[leng-2:leng])
		} else if leng < 16 {
			res = string(runs[0:3]) + "****" + string(runs[leng-3:leng])
		} else {
			res = string(runs[0:4]) + "*****" + string(runs[leng-4:leng])
		}
	}

	return res
}

// CountBase64Byte 统计base64字符串大小,字节.
func (ks *LkkString) CountBase64Byte(str string) (res int) {
	pos := strings.Index(str, ",")
	if pos > 10 {
		img := strings.Replace(str[pos:], "=", "", -1)
		res = int(float64(len(img)) * float64(3.0/4.0))
	}

	return
}

// Strpad 使用fill填充str字符串到指定长度max.
// ptype为填充类型,枚举值(PAD_LEFT,PAD_RIGHT,PAD_BOTH).
func (ks *LkkString) Strpad(str string, fill string, max int, ptype LkkPadType) string {
	runeStr := []rune(str)
	runeStrLen := len(runeStr)
	if runeStrLen >= max || max < 1 || len(fill) == 0 {
		return str
	}

	var leftsize int
	var rightsize int

	switch ptype {
	case PAD_BOTH:
		rlsize := float64(max-runeStrLen) / 2
		leftsize = int(rlsize)
		rightsize = int(rlsize + math.Copysign(0.5, rlsize))

	case PAD_LEFT:
		leftsize = max - runeStrLen

	case PAD_RIGHT:
		rightsize = max - runeStrLen
	}

	buf := make([]rune, 0, max)

	if ptype == PAD_LEFT || ptype == PAD_BOTH {
		for i := 0; i < leftsize; {
			for _, v := range []rune(fill) {
				buf = append(buf, v)
				if i >= leftsize-1 {
					i = leftsize
					break
				} else {
					i++
				}
			}
		}
	}

	buf = append(buf, runeStr...)

	if ptype == PAD_RIGHT || ptype == PAD_BOTH {
		for i := 0; i < rightsize; {
			for _, v := range []rune(fill) {
				buf = append(buf, v)
				if i >= rightsize-1 {
					i = rightsize
					break
				} else {
					i++
				}
			}
		}
	}

	return string(buf)
}

// StrpadLeft 字符串左侧填充,请参考Strpad.
func (ks *LkkString) StrpadLeft(str string, fill string, max int) string {
	return ks.Strpad(str, fill, max, PAD_LEFT)
}

// StrpadRight 字符串右侧填充,请参考Strpad.
func (ks *LkkString) StrpadRight(str string, fill string, max int) string {
	return ks.Strpad(str, fill, max, PAD_RIGHT)
}

// StrpadBoth 字符串两侧填充,请参考Strpad.
func (ks *LkkString) StrpadBoth(str string, fill string, max int) string {
	return ks.Strpad(str, fill, max, PAD_BOTH)
}

// Img2Base64 将字符串转换为base64图片.ext为图片扩展名,默认jpg.
func (ks *LkkString) Img2Base64(content []byte, ext ...string) string {
	var imgType string = "jpg"
	if len(ext) > 0 {
		imgType = strings.ToLower(ext[0])
	}

	return fmt.Sprintf("data:image/%s;base64,%s", imgType, base64.StdEncoding.EncodeToString(content))
}

// Jsonp2Json 将jsonp转为json串.
// Example: forbar({a:"1",b:2}) to {"a":"1","b":2}
func (ks *LkkString) Jsonp2Json(str string) (string, error) {
	start := strings.Index(str, "(")
	end := strings.LastIndex(str, ")")

	if start == -1 || end == -1 {
		return "", errors.New("invalid jsonp.")
	}

	start += 1
	if start >= end {
		return "", errors.New("invalid jsonp.")
	}

	res := str[start:end]

	return res, nil
}

// CountWords 统计字符串中单词的使用情况.
// 返回结果:单词总数;和一个字典,包含每个单词的单独统计.
// 因为没有分词,对中文尚未很好支持.
func (ks *LkkString) CountWords(str string) (int, map[string]int) {
	//过滤标点符号
	var buffer bytes.Buffer
	for _, r := range str {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
			buffer.WriteRune(' ')
		} else {
			buffer.WriteRune(r)
		}
	}

	var total int
	mp := make(map[string]int)
	words := strings.Fields(buffer.String())
	for _, word := range words {
		mp[word] += 1
		total++
	}

	return total, mp
}

// IsLetters 字符串是否全(英文)字母组成.
func (ks *LkkString) IsLetters(str string) bool {
	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return str != ""
}

// IsEmpty 字符串是否为空(包括空格).
func (ks *LkkString) IsEmpty(str string) bool {
	if str == "" || len(ks.Trim(str)) == 0 {
		return true
	}

	return false
}

// IsUpper 字符串是否全部大写.
func (ks *LkkString) IsUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return str != ""
}

// IsLower 字符串是否全部小写.
func (ks *LkkString) IsLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return str != ""
}

// HasLetter 字符串是否含有(英文)字母.
func (ks *LkkString) HasLetter(str string) bool {
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}

// IsUtf8 字符串是否UTF-8编码.
func (ks *LkkString) IsUtf8(str string) bool {
	return str != "" && utf8.ValidString(str)
}

// IsASCII 是否IsASCII字符串.
func (ks *LkkString) IsASCII(str string) bool {
	//return str != "" && RegAscii.MatchString(str)
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] > 127 {
			return false
		}
	}

	return str != ""
}

// IsMultibyte 字符串是否含有多字节字符.
func (ks *LkkString) IsMultibyte(str string) bool {
	return str != "" && RegMultiByte.MatchString(str)
}

// HasFullWidth 是否含有全角字符.
func (ks *LkkString) HasFullWidth(str string) bool {
	return str != "" && RegFullWidth.MatchString(str)
}

// HasHalfWidth 是否含有半角字符.
func (ks *LkkString) HasHalfWidth(str string) bool {
	return str != "" && RegHalfWidth.MatchString(str)
}

// IsEnglish 字符串是否纯英文.letterCase是否检查大小写,枚举值(CASE_NONE,CASE_LOWER,CASE_UPPER).
func (ks *LkkString) IsEnglish(str string, letterCase LkkCaseSwitch) bool {
	switch letterCase {
	case CASE_NONE:
		return ks.IsLetters(str)
	case CASE_LOWER:
		return str != "" && RegAlphaLower.MatchString(str)
	case CASE_UPPER:
		return str != "" && RegAlphaUpper.MatchString(str)
	default:
		return ks.IsLetters(str)
	}
}

// HasEnglish 是否含有英文字符,HasLetter的别名.
func (ks *LkkString) HasEnglish(str string) bool {
	return ks.HasLetter(str)
}

// HasChinese 字符串是否含有中文.
func (ks *LkkString) HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}

// IsChinese 字符串是否全部中文.
func (ks *LkkString) IsChinese(str string) bool {
	return str != "" && RegChineseAll.MatchString(str)
}

// IsChineseName 字符串是否中文名称.
func (ks *LkkString) IsChineseName(str string) bool {
	return str != "" && RegChineseName.MatchString(str)
}

// IsWord 是否词语(不以下划线开头的中文、英文、数字、下划线).
func (ks *LkkString) IsWord(str string) bool {
	return str != "" && RegWord.MatchString(str)
}

// HasSpecialChar 字符串是否含有特殊字符.
func (ks *LkkString) HasSpecialChar(str string) (res bool) {
	if str == "" {
		return
	}

	for _, r := range str {
		// IsPunct 判断 r 是否为一个标点字符 (类别 P)
		// IsSymbol 判断 r 是否为一个符号字符
		// IsMark 判断 r 是否为一个 mark 字符 (类别 M)
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
			res = true
			return
		}
	}

	return
}

// IsJSON 字符串是否合法的json格式.
func (ks *LkkString) IsJSON(str string) bool {
	length := len(str)
	if length == 0 {
		return false
	} else if (str[0] != '{' || str[length-1] != '}') && (str[0] != '[' || str[length-1] != ']') {
		return false
	}

	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsIP 检查字符串是否IP地址.
func (ks *LkkString) IsIP(str string) bool {
	return str != "" && net.ParseIP(str) != nil
}

// IsIPv4 检查字符串是否IPv4地址.
func (ks *LkkString) IsIPv4(str string) bool {
	ipAddr := net.ParseIP(str)
	// 不是合法的IP地址
	if ipAddr == nil {
		return false
	}

	return ipAddr.To4() != nil && strings.ContainsRune(str, '.')
}

// IsIPv6 检查字符串是否IPv6地址.
func (ks *LkkString) IsIPv6(str string) bool {
	ipAddr := net.ParseIP(str)
	return ipAddr != nil && strings.ContainsRune(str, ':')
}

// IsPort 字符串或数字是否端口号.
func (ks *LkkString) IsPort(val interface{}) bool {
	if KConv.IsInt(val) {
		port := KConv.ToInt(val)
		if port > 0 && port < 65536 {
			return true
		}
	}

	return false
}

// IsDNSName 是否DNS名称.
func (ks *LkkString) IsDNSName(str string) bool {
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return !ks.IsIP(str) && RegDNSname.MatchString(str)
}

// IsDialString 是否网络拨号字符串(形如127.0.0.1:80),用于net.Dial()检查.
func (ks *LkkString) IsDialString(str string) bool {
	h, p, err := net.SplitHostPort(str)
	if err == nil && h != "" && p != "" && (ks.IsDNSName(h) || ks.IsIP(h)) && ks.IsPort(p) {
		return true
	}

	return false
}

// IsMACAddr 是否MAC物理网卡地址.
func (ks *LkkString) IsMACAddr(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

// IsHost 字符串是否主机名(IP或DNS名称).
func (ks *LkkString) IsHost(str string) bool {
	return ks.IsIP(str) || ks.IsDNSName(str)
}

// IsEmail 检查字符串是否邮箱.参数validateTrue,是否验证邮箱主机的真实性.
func (ks *LkkString) IsEmail(email string, validateHost bool) (bool, error) {
	//长度检查
	length := len(email)
	at := strings.LastIndexByte(email, '@')
	if (length < 6 || length > 254) || (at <= 0 || at > length-3) {
		return false, fmt.Errorf("invalid email length")
	}

	//验证邮箱格式
	chkFormat := RegEmail.MatchString(email)
	if !chkFormat {
		return false, fmt.Errorf("invalid email format")
	}

	//验证主机
	if validateHost {
		host := email[at+1:]
		if _, err := net.LookupMX(host); err != nil {
			//因无法确定mx主机的smtp端口,所以去掉Hello/Mail/Rcpt检查邮箱是否存在
			//仅检查主机是否有效
			//TODO 仅对国内几家大的邮件厂家进行检查
			if _, err := net.LookupIP(host); err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

// IsMobilecn 检查字符串是否中国大陆手机号.
func (ks *LkkString) IsMobilecn(str string) bool {
	return str != "" && RegMobilecn.MatchString(str)
}

// IsTel 是否固定电话或400/800电话.
func (ks *LkkString) IsTel(str string) bool {
	return str != "" && RegTelephone.MatchString(str)
}

// IsPhone 是否电话号码(手机或固话).
func (ks *LkkString) IsPhone(str string) bool {
	return str != "" && RegPhone.MatchString(str)
}

// IsCreditNo 检查是否(15或18位)身份证号码,并返回经校验的号码.
func (ks *LkkString) IsCreditNo(str string) (bool, string) {
	chk := str != "" && RegCreditno.MatchString(str)
	if !chk {
		return false, ""
	}

	// 检查省份代码
	if _, chk = CreditArea[str[0:2]]; !chk {
		return false, ""
	}

	// 将15位身份证升级到18位
	leng := len(str)
	if leng == 15 {
		// 先转为17位,如果身份证顺序码是996 997 998 999,这些是为百岁以上老人的特殊编码
		if chk, _ = ks.Dstrpos(str[12:], []string{"996", "997", "998", "999"}, false); chk {
			str = str[0:6] + "18" + str[6:]
		} else {
			str = str[0:6] + "19" + str[6:]
		}

		// 再加上校验码
		code := append([]byte{}, creditChecksum(str))
		str += string(code)
	}

	// 检查生日
	birthday := str[6:10] + "-" + str[10:12] + "-" + str[12:14]
	chk, tim := KTime.IsDate2time(birthday)
	now := KTime.UnixTime()
	if !chk {
		return false, ""
	} else if tim >= now {
		return false, ""
	}

	// 18位身份证需要验证最后一位校验位
	if leng == 18 {
		str = strings.ToUpper(str)
		if str[17] != creditChecksum(str) {
			return false, ""
		}
	}

	return true, str
}

// IsAlphaNumeric 是否字母或数字.
func (ks *LkkString) IsAlphaNumeric(str string) bool {
	return str != "" && RegAlphaNumeric.MatchString(str)
}

// IsHexcolor 检查是否十六进制颜色,并返回带"#"的修正值.
func (ks *LkkString) IsHexcolor(str string) (bool, string) {
	chk := str != "" && RegHexcolor.MatchString(str)
	if chk && !strings.ContainsRune(str, '#') {
		str = "#" + strings.ToUpper(str)
	}
	return chk, str
}

// IsRGBcolor 检查字符串是否RGB颜色格式.
func (ks *LkkString) IsRGBcolor(str string) bool {
	return str != "" && RegRgbcolor.MatchString(str)
}

// IsBlank 是否空(空白)字符.
func (ks *LkkString) IsBlank(str string) bool {
	// Check length
	if len(str) > 0 {
		// Iterate string
		for i := range str {
			// Check about char different from whitespace
			// 227为全角空格
			if str[i] > 32 && str[i] != 227 {
				return false
			}
		}
	}
	return true
}

// IsWhitespaces 是否全部空白字符,不包括空字符串.
func (ks *LkkString) IsWhitespaces(str string) bool {
	return str != "" && RegWhitespaceAll.MatchString(str)
}

// HasWhitespace 是否带有空白字符.
func (ks *LkkString) HasWhitespace(str string) bool {
	return str != "" && RegWhitespaceHas.MatchString(str)
}

// IsBase64 是否base64字符串.
func (ks *LkkString) IsBase64(str string) bool {
	return str != "" && RegBase64.MatchString(str)
}

// IsBase64Image 是否base64编码的图片.
func (ks *LkkString) IsBase64Image(str string) bool {
	if str == "" || !strings.ContainsRune(str, ',') {
		return false
	}

	dataURI := strings.Split(str, ",")
	return RegBase64Image.MatchString(dataURI[0]) && RegBase64.MatchString(dataURI[1])
}

// IsRsaPublicKey 检查字符串是否RSA的公钥,keylen为密钥长度.
func (ks *LkkString) IsRsaPublicKey(str string, keylen int) bool {
	bb := bytes.NewBufferString(str)
	pemBytes, _ := ioutil.ReadAll(bb)

	// 获取公钥
	block, _ := pem.Decode(pemBytes)
	if block != nil && block.Type != "PUBLIC KEY" {
		return false
	}
	var der []byte
	var err error

	if block != nil {
		der = block.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(str)
		if err != nil {
			return false
		}
	}

	key, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return false
	}
	pubkey, ok := key.(*rsa.PublicKey)
	if !ok {
		return false
	}
	bitlen := len(pubkey.N.Bytes()) * 8
	return bitlen == int(keylen)
}

// IsUrl 检查字符串是否URL.
func (ks *LkkString) IsUrl(str string) bool {
	if str == "" || len(str) <= 3 || utf8.RuneCountInString(str) >= 2083 || strings.HasPrefix(str, ".") {
		return false
	}

	res, err := url.ParseRequestURI(str)
	if err != nil {
		return false //Couldn't even parse the url
	}
	if len(res.Scheme) == 0 {
		return false //No Scheme found
	}

	return true
}

// IsUrlExists 检查URL是否存在.
func (ks *LkkString) IsUrlExists(str string) bool {
	if !ks.IsUrl(str) {
		return false
	}

	client := &http.Client{}
	resp, err := client.Head(str)
	if err != nil {
		return false
	} else if resp.StatusCode == 404 {
		return false
	}

	return true
}

// IsMd5 是否md5值.
func (ks *LkkString) IsMd5(str string) bool {
	return str != "" && RegMd5.MatchString(str)
}

// IsSha1 是否Sha1值.
func (ks *LkkString) IsSha1(str string) bool {
	return str != "" && RegSha1.MatchString(str)
}

// IsSha256 是否Sha256值.
func (ks *LkkString) IsSha256(str string) bool {
	return str != "" && RegSha256.MatchString(str)
}

// IsSha512 是否Sha512值.
func (ks *LkkString) IsSha512(str string) bool {
	return str != "" && RegSha512.MatchString(str)
}

// StartsWith 字符串str是否以sub开头.
func (ks *LkkString) StartsWith(str, sub string, ignoreCase bool) bool {
	if str != "" && sub != "" {
		i := ks.Index(str, sub, ignoreCase)
		return i == 0
	}

	return false
}

// EndsWith 字符串str是否以sub结尾.
func (ks *LkkString) EndsWith(str, sub string, ignoreCase bool) bool {
	if str != "" && sub != "" {
		i := ks.LastIndex(str, sub, ignoreCase)
		return i != -1 && (len(str)-len(sub)) == i
	}

	return false
}

// RemoveEmoji 移除字符串中的表情符(使用正则,效率较低).
func (ks *LkkString) RemoveEmoji(str string) string {
	return RegEmoji.ReplaceAllString(str, "")
}

// Gravatar 获取Gravatar头像地址.
func (ks *LkkString) Gravatar(email string, size uint16) string {
	h := md5.New()
	_, _ = io.WriteString(h, email)
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d", h.Sum(nil), size)
}

// AtWho 查找被@的用户名.minLen为用户名最小长度,默认5.
func (ks *LkkString) AtWho(text string, minLen ...int) []string {
	var result = []string{}
	var username string
	var min int = 5
	if len(minLen) > 0 && minLen[0] > 0 {
		min = minLen[0]
	}

	for _, line := range strings.Split(text, "\n") {
		if len(line) == 0 {
			continue
		}
		for {
			index := strings.Index(line, "@")
			if index == -1 {
				break
			} else if index > 0 {
				r := rune(line[index-1])
				if unicode.IsUpper(r) || unicode.IsLower(r) {
					line = line[index+1:]
				} else {
					line = line[index:]
				}
			} else if index == 0 {
				// the "@" is first characters
				endIndex := strings.Index(line, " ")
				if endIndex == -1 {
					username = line[1:]
				} else {
					username = line[1:endIndex]
				}

				if len(username) >= min && RegUsernameen.MatchString(username) && !KArr.InStringSlice(username, result) {
					result = append(result, username)
				}

				if endIndex == -1 {
					break
				}

				line = line[endIndex:]
			}
		}
	}

	return result
}
