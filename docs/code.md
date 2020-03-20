### 代码片段

	// TimerOnce 一次性定时器
	TimerOnce struct {
		Pt      *time.Timer               // 定时器指针
		Dr      time.Duration             // 时间间隔
		Fn      func(args ...interface{}) // 执行函数
		Args    interface{}               // 执行函数的参数
		Count   int                       // 已执行次数
		running bool                      // 是否在运行
	}
	// TimerCycle 循环定时器
	TimerCycle struct {
		Pt       *time.Ticker              // 定时器指针
		Dr       time.Duration             // 时间间隔
		Fn       func(args ...interface{}) // 执行函数
		Args     interface{}               // 执行函数的参数
		Max      int                       // 最大执行次数
		Count    int                       // 已执行次数
		LastTime time.Time                 // 上次执行时间
		running  bool                      // 是否在运行
	}
	// LkkTimers 定时器容器
	LkkTimers struct {
		Onces     map[int64]*TimerOnce  // 一次性定时器字典
		Cycles    map[int64]*TimerCycle // 循环定时器字典
		OnceRuns  uint                  // 一次性定时器已执行次数
		CycleRuns uint                  // 循环定时器已执行次数
	}
	
    // KTimer utilities
    KTimer *LkkTimers


type ITimer interface {
	SetAfter()
	ClearAfter()
	ClearAfterAll()
	SetInterval()
	ClearInterval()
	ClearIntervalAll()
	Start()
	Stop()
	Reset()
	Clear()
	CountOnces()
	CountCycles()
	CountTimers()
}

func (kt *LkkTime) GetTimer() *LkkTimers {
	if KTimer == nil {
		KTimer = &LkkTimers{}
	}
	return KTimer
}

emoji表情的处理
import "unicode/utf8"

func FilterEmoji(content string) string {
    new_content := ""
    for _, value := range content {
        _, size := utf8.DecodeRuneInString(string(value))
        if size <= 3 {
            new_content += string(value)
        }
    }
    return new_content
}

// 检查是否存在特殊符号
// 1. emoji字符
// 2. ascii控制字符
// 3. \ " '
// val:待检查的字符串
// 返回值:
// bool:true:有特殊字符 false:无特殊字符
func IfHaveSpecialChar(val string) bool {
	if len(val) <= 0 {
		return false
	}

	// 表情符号过滤
	// Wide UCS-4 build
	emojiReg, _ := regexp.Compile("[^\U00000000-\U0000FFFF]+")
	if emojiReg.Match([]byte(val)) {
		return true
	}

	// 排除控制字符和特殊字符
	for _, charItem := range val {
		// 排除控制字符
		if (charItem > 0 && charItem < 0x20) || charItem == 0x7F {
			return true
		}

		// 排除部分特殊字符：  \ " '
		switch charItem {
		case '\\':
			fallthrough
		case '"':
			fallthrough
		case '\'':
			return true
		}
	}

	return false
}

// strip tags in html string
func StripTags(src string) string {
	//去除style,script,html tag
	re := regexp.MustCompile(`(?s)<(?:style|script)[^<>]*>.*?</(?:style|script)>|</?[a-z][a-z0-9]*[^<>]*>|<!--.*?-->`)
	src = re.ReplaceAllString(src, "")

	//trim all spaces(2+) into \n
	re = regexp.MustCompile(`\s{2,}`)
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

