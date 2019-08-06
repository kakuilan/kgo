package kgo

import (
	"bytes"
	"fmt"
	"github.com/json-iterator/go"
	"hash/crc32"
	"html"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Nl2br 将换行符转换为br标签
func (ks *LkkString) Nl2br(html string) string {
	if html == "" {
		return ""
	}
	return strings.Replace(html, "\n", "<br />", -1)
}

// StripTags 过滤html和php标签
func (ks *LkkString) StripTags(html string) string {
	if html == "" {
		return ""
	}
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(html, "")
}

// Md5 获取字符串md5值,length指定结果长度32/16
func (ks *LkkString) Md5(str string, length uint8) string {
	return string(md5Str([]byte(str), length))
}

// Sha1 计算字符串的 sha1 散列值
func (ks *LkkString) Sha1(str string) string {
	return string(sha1Str([]byte(str)))
}

// Sha256 计算字符串的 sha256 散列值
func (ks *LkkString) Sha256(str string) string {
	return string(sha256Str([]byte(str)))
}

// Random 生成随机字符串;length为长度,stype为枚举(RAND_STRING_ALPHA,RAND_STRING_NUMERIC,RAND_STRING_ALPHANUM,RAND_STRING_SPECIAL,RAND_STRING_CHINESE)
func (ks *LkkString) Random(length uint8, stype LkkRandString) string {
	if length == 0 {
		return ""
	}

	var letter []rune
	alphas := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specials := "~!@#$%^&*()_+{}:|<>?`-=;,."

	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Nanosecond)

	switch stype {
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

// Strpos 查找字符串首次出现的位置
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

// Stripos  查找字符串首次出现的位置（不区分大小写）
func (ks *LkkString) Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	haystack = haystack[offset:]
	pos := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Ucfirst 将字符串的首字母转换为大写
func (ks *LkkString) Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

// Lcfirst 使一个字符串的第一个字符小写
func (ks *LkkString) Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

// Substr 返回字符串的子串
func (ks *LkkString) Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

// Strrev 反转字符串
func (ks *LkkString) Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ChunkSplit 将字符串分割成小块.body为要分割的字符,chunklen为分割的尺寸,end为行尾序列符号
func (ks *LkkString) ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	length := uint(len(runes))
	if length <= 1 || length < chunklen {
		return body + end
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

// Strlen 获取字符串长度
func (ks *LkkString) Strlen(str string) int {
	return len(str)
}

// MbStrlen 获取字符串的长度,多字节的字符被计为 1
func (ks *LkkString) MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

// StrShuffle 随机打乱一个字符串
func (ks *LkkString) StrShuffle(str string) string {
	runes := []rune(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, len(runes))
	for i, v := range r.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

// Trim 去除字符串首尾处的空白字符（或者其他字符）
func (ks *LkkString) Trim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {
		mask = " \\t\\n\\r\\0\\x0B　"
	} else {
		mask = characterMask[0]
	}
	return strings.Trim(str, mask)
}

// Ltrim 删除字符串开头的空白字符（或其他字符）
func (ks *LkkString) Ltrim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {
		mask = " \\t\\n\\r\\0\\x0B　"
	} else {
		mask = characterMask[0]
	}
	return strings.TrimLeft(str, mask)
}

// Rtrim 删除字符串末端的空白字符（或者其他字符）
func (ks *LkkString) Rtrim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {
		mask = " \\t\\n\\r\\0\\x0B　"
	} else {
		mask = characterMask[0]
	}
	return strings.TrimRight(str, mask)
}

// Chr 返回相对应于 ascii 所指定的单个字符
func (ks *LkkString) Chr(ascii int) string {
	return string(ascii)
}

// Ord 将首字符转换为rune
func (ks *LkkString) Ord(char string) rune {
	r, _ := utf8.DecodeRune([]byte(char))
	return r
}

// JsonEncode 对变量进行 JSON 编码
// 依赖库github.com/json-iterator/go
func (ks *LkkString) JsonEncode(val interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(val)
}

// JsonDecode 对 JSON 格式的字符串进行解码,注意val使用指针
// 依赖库github.com/json-iterator/go
func (ks *LkkString) JsonDecode(data []byte, val interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, val)
}

// Addslashes 使用反斜线引用字符串
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

// Stripslashes 反引用一个引用字符串
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

// Quotemeta 转义元字符集,包括 . \ + * ? [ ^ ] ( $ )
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

// Htmlentities 将字符转换为 HTML 转义字符
func (ks *LkkString) Htmlentities(str string) string {
	return html.EscapeString(str)
}

// HtmlentityDecode 将HTML实体转换为它们对应的字符
func (ks *LkkString) HtmlentityDecode(str string) string {
	return html.UnescapeString(str)
}

// Crc32 计算一个字符串的 crc32 多项式
func (ks *LkkString) Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

// SimilarText 计算两个字符串的相似度,返回在两个字符串中匹配字符的数目;percent为相似程度百分数
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

// Explode 使用一个字符串分割另一个字符串
func (ks *LkkString) Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

// Uniqid 获取一个带前缀、基于当前时间微秒数的唯一ID
func (ks *LkkString) Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}
