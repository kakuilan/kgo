# Changelog

All notable changes to this project will be documented in this file.

## [v0.3.6]- 2022-11-29

#### Added

- 新增`KOS.DownloadFile`,文件下载

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.3.5]- 2022-10-10

#### Added

- 新增`KStr.StrOffset`,字符串整体偏移

#### Fixed

- fix ioutil.ReadAll Deprecated, replace to io.ReadAll
- fix ioutil.ReadFile Deprecated, replace to os.ReadFile

#### Changed

- none

#### Removed

- none

## [v0.3.4]- 2022-09-05

#### Added

- 新增`LkkFile.Md5Reader`
- 新增`LkkFile.ShaXReader`

#### Fixed

- none

#### Changed

- 重命名 `LkkFile.Md5` 为 `LkkFile.Md5File`
- 重命名 `LkkFile.ShaX` 为 `LkkFile.ShaXFile`

#### Removed

- none

## [v0.3.3]- 2022-08-01

#### Added

- none

#### Fixed

- 优化`KFile.CopyFile`
- 优化`KFile.FastCopy`
- 优化`KFile.CopyDir`
- 优化`KFile.DelDir`
- 优化`KFile.WriteFile`
- 优化`KFile.FileTree`
- 优化`KFile.Md5`
- 优化`KFile.IsZip`
- 优化`KFile.TarGz`
- 优化`KFile.UnTarGz`
- 优化os_darwin_cgo下`getProcessPathByPid`
- 优化os_darwin_cgo下`LkkOS.CpuUsage`
- 优化os_linux下`getPidByInode`
- 优化os_linux下`LkkOS.Uptime`
- 优化os_windows下`LkkOS.MemoryUsage`
- 优化os_windows下`LkkOS.CpuUsage`
- 优化os_windows下`LkkOS.DiskUsage`
- 优化os_windows下`LkkOS.Uptime`
- 优化os_windows下`LkkOS.GetBiosInfo`
- 优化os_windows下`LkkOS.GetBoardInfo`
- 优化os_windows下`LkkOS.GetCpuInfo`
- 优化os_windows下`LkkOS.IsProcessExists`
- 优化`isEmpty`
- 优化`shaXByte`
- 优化`pkcs7UnPadding`

#### Changed

- 修改`KFile.CopyLink`,增加`cover`文件覆盖参数
- 修改`LkkOS.HomeDir`,不再自行区分windows/unix,使用自带的`os.UserHomeDir`

#### Removed

- 删除os_darwin下`getPidByPort`
- 删除os_windows下`getPidByPort`

## [v0.3.2]- 2022-06-26

#### Added

- none

#### Fixed

- 修复`KConv.IsFloat`

#### Changed

- none

#### Removed

- none

## [v0.3.1]- 2022-06-21

#### Added

- none

#### Fixed

- 修复`KFile.Touch`,创建文件前检查文件是否存在
- 修复非cgo的darwin下`KOS.GetProcessExecPath`编译失败问题

#### Changed

- `SystemInfo`系统信息结构体增加`SystemArch`字段

#### Removed

- none

## [v0.3.0]- 2022-05-27

#### Added

- 新增`KStr.PasswordSafeLevel`,检查密码安全等级

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.2.9]- 2022-04-15

#### Added

- none

#### Fixed

- 修复`KStr.MatchEquations`

#### Changed

- 优化`LkkEncrypt.aesDecrypt`

#### Removed

- none

## [v0.2.8]- 2022-04-15

#### Added

- none

#### Fixed

- none

#### Changed

- rename `KArr.CopyToStruct` to `KArr.CopyStruct`

#### Removed

- none

## [v0.2.7]- 2022-04-14

#### Added

- none

#### Fixed

- 修改`reflectTypesMap`,获取匿名字段

#### Changed

- none

#### Removed

- none

## [v0.2.6]- 2022-04-14

#### Added

- 新增`KArr.CopyToStruct`,拷贝结构体

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.2.5]- 2022-03-31

#### Added

- none

#### Fixed

- 优化`KEncr.AuthCode`

#### Changed

- none

#### Removed

- none

## [v0.2.4]- 2022-03-05

#### Added

- 新增`KDbug.Stacks`,获取堆栈信息

#### Fixed

- 修改`KDbug.GetCallName`,参数`f`支持`uintptr`类型

#### Changed

- none

#### Removed

- none

## [v0.2.3]- 2022-01-21

#### Added

- 新增`KStr.UuidV5`,根据提供的字符,使用sha1生成36位哈希值

#### Fixed

- 修改`KStr.UuidV4`,添加version/variant信息

#### Changed

- none

#### Removed

- none

## [v0.2.2]- 2022-01-10

#### Added

- 新增`KEncr.RsaPrivateEncryptLong`使用RSA私钥加密长文本
- 新增`KEncr.RsaPublicDecryptLong`使用RSA公钥解密长文本

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.2.1]- 2022-01-05

#### Added

- 新增`KStr.ChunkBytes`将字节切片分割为多个小块
- 新增`KEncr.RsaPublicEncryptLong`使用RSA公钥加密长文本
- 新增`KEncr.RsaPrivateDecryptLong`使用RSA私钥解密长文本

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.2.0]- 2021-12-09

#### Added

- 新增`KStr.ToRunes`将字符串转为字符切片
- 新增`KConv.ToInterfaces`将变量转为接口切片
- 新增`KDbug.WrapError`错误包裹

#### Fixed

- 修改`KStr.IsASCII`根据字符串长度使用不同方法

#### Changed

- 重命名`KStr.IsHexcolor`为`IsHexColor`
- 重命名`KStr.IsRGBcolor`为`IsRgbColor`
- 将部分公开变量转为私有

#### Removed

- none

## [v0.1.9]- 2021-11-27

#### Added

- 新增`IsPointer`检查变量是否指针类型

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.1.8]- 2021-11-08

#### Added

- none

#### Fixed

- 修复`KDbug.DumpPrint`打印多变量问题
- 修改`KStr.Ucwords`因go1.18废弃strings.Title,使用cases.Title代替

#### Changed

- none

#### Removed

- none

## [v0.1.7]- 2021-10-15

#### Added

- none

#### Fixed

- 修复`KEncr.AuthCode`中keyb变化问题

#### Changed

- none

#### Removed

- none

## [v0.1.6]- 2021-8-21

#### Added

- 新增`KNum.NearLogarithm`,求对数临近值
- 新增`KNum.SplitNaturalNum`,将自然数按底数拆解

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.1.5]- 2021-6-11

- 重构版本，改动太多，懒得整理，就不再一一列出

## [v0.1.4]- 2020-12-1

#### Added

- 新增`KArr.ArrayIntersect`,求数组交集
- 新增`KArr.DeleteSliceItems`,删除切片多个元素
- 新增`KFile.FormatPath`,格式化路径
- 新增`KNum.IsNaturalRange`,是否自然数数组
- 新增`KNum.Log`,求任意底数的对数

#### Fixed

- 修改`KNum.IsNatural`,自然数包含0

#### Changed

- 修改`KArr.ArrayDiff`,增加比较方式参数`compareType`,返回字典
- 修改`KFile.FormatDir`,过滤特殊字符
- 修改`KNum.IsNan`,接收任意类型参数
- 修改`KNum.NumSign`,结果类型为int8
- 修改`KNum.Range`,支持生成降序的数组

#### Removed

- none

## [v0.1.3]- 2020-11-11

#### Added

- none

#### Fixed

- none

#### Changed

- 修改`KStr.Chr`,使用rune转换字符码
- 修改`KDbug.CallMethod`错误提示

#### Removed

- none

## [v0.1.2]- 2020-06-25

#### Added

- `KStr.Md5Byte`
- `KStr.ShaXByte`

#### Fixed

- none

#### Changed

- rename `md5Str` to `md5Byte`
- rename `shaXStr` to `shaXByte`

#### Removed

- none

## [v0.1.1]- 2020-06-23

#### Added

- `KStr.Serialize`
- `KStr.UnSerialize`
- `KStr.TrimBOM`

#### Fixed

- none

#### Changed

- none

#### Removed

- none

## [v0.1.0]- 2020-06-22

#### Added

- none

#### Fixed

- none

#### Changed

- 修改`KTime.Str2Timestruct`使用本地时区而非UTC

#### Removed

- none

## [v0.0.9]- 2020-06-20

#### Added

- none

#### Fixed

- none

#### Changed

- 修改`KEncr.Base64Encode`结果类型为[]byte
- 修改`KEncr.Base64Decode`参数类型为[]byte
- 修改`KEncr.Base64UrlEncode`结果类型为[]byte
- 修改`KEncr.Base64UrlDecode`结果类型为[]byte
- 修改`KEncr.AuthCode`参数和结果类型为[]byte
- 修改`KEncr.EasyEncrypt`参数和结果类型为[]byte
- 修改`KEncr.EasyDecrypt`参数和结果类型为[]byte
- 修改`KEncr.HmacShaX`结果类型为[]byte

#### Removed

- none

## [v0.0.8]- 2020-05-31

#### Added

- none

#### Fixed

- none

#### Changed

- `KOS.GetSystemInfo`增加`SystemOs`操作系统名称字段

#### Removed

- none

## [v0.0.7]- 2020-05-21

#### Added

- none

#### Fixed

- 修复`KStr.ToCamelCase`将首字母大写的驼峰串(如SayHello->Sayhello)转换错误问题

#### Changed

- none

#### Removed

- none

## [v0.0.6]- 2020-04-28

#### Added

- `KNum.AbsInt`

#### Fixed

- none

#### Changed

- rename `KNum.Abs` to `KNum.AbsFloat`

#### Removed

- none

## [v0.0.5]- 2020-03-20

#### Added

- `KArr.InInt64Slice`
- `KArr.InIntSlice`
- `KArr.InStringSlice`
- `KArr.IsEqualArray`
- `KConv.Byte2Hexs`
- `KConv.Hexs2Byte`
- `KConv.Runes2Bytes`
- `KEncr.AesCBCDecrypt`
- `KEncr.AesCBCEncrypt`
- `KEncr.AesCFBDecrypt`
- `KEncr.AesCFBEncrypt`
- `KEncr.AesCTRDecrypt`
- `KEncr.AesCTREncrypt`
- `KEncr.AesOFBDecrypt`
- `KEncr.AesOFBEncrypt`
- `KEncr.GenerateRsaKeys`
- `KEncr.RsaPrivateDecrypt`
- `KEncr.RsaPrivateEncrypt`
- `KEncr.RsaPublicDecrypt`
- `KEncr.RsaPublicEncrypt`
- `KFile.AppendFile`
- `KFile.GetFileMode`
- `KFile.ReadFirstLine`
- `KFile.ReadLastLine`
- `KNum.Percent`
- `KNum.RoundPlus`
- `KOS.GetBiosInfo`
- `KOS.GetBoardInfo`
- `KOS.GetCpuInfo`
- `KOS.IsProcessExists`
- `KStr.AtWho`
- `KStr.ClearUrlPrefix`
- `KStr.ClearUrlSuffix`
- `KStr.Gravatar`
- `KStr.IsWord`
- `KStr.RemoveEmoji`
- `KStr.UuidV4`
- `KTime.DaysBetween`
- `KTime.EndOfDay`
- `KTime.EndOfMonth`
- `KTime.EndOfWeek`
- `KTime.EndOfYear`
- `KTime.StartOfDay`
- `KTime.StartOfMonth`
- `KTime.StartOfWeek`
- `KTime.StartOfYear`

#### Fixed

- none

#### Changed

- `KFile.IsFile` 增加文件类型参数LkkFileType
- `KFile.WriteFile` 增加权限参数perm
- `KOS.Getenv` 增加默认值参数
- `KStr.Random` 移除time.Sleep
- `KTime.GetMonthDays` 放弃map,直接比较
- rename `KOS.GetProcessExeByPid` to `KOS.GetProcessExecPath`
- rename `KTime.Time` to `KTime.UnixTime`

#### Removed

- none

## [v0.0.4]- 2020-03-06

#### Added

- `KStr.Index`
- `KStr.LastIndex`

#### Fixed

- none

#### Changed

- `KStr.RemoveBefore` 增加参数ignoreCase
- `KStr.RemoveAfter` 增加参数ignoreCase
- `KStr.StartsWith` 增加参数ignoreCase,使用Index代替MbSubstr
- `KStr.EndsWith` 增加参数ignoreCase,使用LastIndex代替MbSubstr

#### Removed

- none

## [v0.0.3]- 2020-03-03

#### Added

- 增加常量`DYNAMIC_KEY_LEN`动态密钥长度

#### Fixed

- `KEncr.AuthCode` 修复bounds out of range错误

#### Changed

- `KEncr.AuthCode` 动态密钥长度改为8
- `KEncr.EasyEncrypt` 动态密钥长度改为8
- `KEncr.EasyDecrypt` 动态密钥长度改为8
- `KTime.CheckDate(month, day, year int)` to `CheckDate(year, month, day int)`
- `KNum.ByteFormat` 增加'delimiter'参数,为数字和单位间的分隔符

#### Removed

- none

## [v0.0.2]- 2020-02-09

#### Added

- `KArr.JoinInts`
- `KArr.JoinStrings`
- `KArr.Unique64Ints`
- `KArr.UniqueInts`
- `KArr.UniqueStrings`
- `KConv.ToBool`
- `KConv.IsNil`
- `KDbug.CallMethod`
- `KDbug.GetFuncDir`
- `KDbug.GetFuncFile`
- `KDbug.GetFuncPackage`
- `KDbug.GetMethod`
- `KDbug.HasMethod`
- `KFile.CountLines`
- `KFile.IsZip`
- `KFile.ReadInArray`
- `KFile.UnZip`
- `KFile.Zip`
- `KNum.Average`
- `KNum.AverageFloat64`
- `KNum.AverageInt`
- `KNum.FloatEqual`
- `KNum.GeoDistance`
- `KNum.MaxFloat64`
- `KNum.MaxInt`
- `KNum.RandFloat64`
- `KNum.RandInt64`
- `KNum:Sum`
- `KNum:SumFloat64`
- `KNum:SumInt`
- `KOS.ForceGC`
- `KOS.GetPidByPort`
- `KOS.GetProcessExeByPid`
- `KOS.TriggerGC`
- `KStr.CountWords`
- `KStr.EndsWith`
- `KStr.Img2Base64`
- `KStr.IsBlank`
- `KStr.IsEmpty`
- `KStr.IsLower`
- `KStr.IsMd5`
- `KStr.IsSha1`
- `KStr.IsSha256`
- `KStr.IsSha512`
- `KStr.IsUpper`
- `KStr.Jsonp2Json`
- `KStr.StartsWith`
- `KStr.ToKebabCase`
- `KStr.ToSnakeCase`
- `KTime.Day`
- `KTime.Hour`
- `KTime.Minute`
- `KTime.Month`
- `KTime.Second`
- `KTime.Str2Timestruct`
- `KTime.Year`

#### Fixed

- `KStr.Trim`, 当输入"0"时,结果为空的BUG.

#### Changed

- `KArr.Implode`, 增加对map的处理.
- `KEncr.EasyDecrypt`, 改进循环.
- `KEncr.EasyEncrypt`, 改进循环.
- `KNum.Max`, 接受任意类型的参数.
- `KNum.Min`, 接受任意类型的参数.
- `KNum.Sum`, 只对数值类型求和.
- `KOS.Pwd`, 弃用os.Args[0],改用os.Executable.
- `KStr.IsASCII`, 弃用正则判断.
- `KStr.IsEmail`, 去掉邮箱是否真实存在的检查.
- `KStr.MbSubstr`, 允许参数start/length为负数.
- `KStr.Substr`, 允许参数start/length为负数.
- rename `KArr.ArraySearchMutilItem` to `ArraySearchMutil`
- rename `KArr.MapMerge` to `MergeMap`
- rename `KArr.SliceMerge` to `MergeSlice`
- rename `KConv.Bin2dec` to `Bin2Dec`
- rename `KConv.Bin2hex` to `Bin2Hex`
- rename `KConv.ByteToFloat64` to `Byte2Float64`
- rename `KConv.ByteToInt64` to `Byte2Int64`
- rename `KConv.BytesSlice2Str` to `Bytes2Str`
- rename `KConv.Dec2bin` to `Dec2Bin`
- rename `KConv.Dec2hex` to `Dec2Hex`
- rename `KConv.Dec2oct` to `Dec2Oct`
- rename `KConv.Hex2bin` to `Hex2Bin`
- rename `KConv.Hex2dec` to `Hex2Dec`
- rename `KConv.Ip2long` to `Ip2Long`
- rename `KConv.Long2ip` to `Long2Ip`
- rename `KConv.Oct2dec` to `Oct2Dec`
- rename `KConv.Str2ByteSlice` to `Str2Bytes`
- rename `KConv.StrictStr2Float` to `Str2FloatStrict`
- rename `KConv.StrictStr2Int` to `Str2IntStrict`
- rename `KConv.StrictStr2Uint` to `Str2UintStrict`
- rename `KFile.Filemtime` to `GetModTime`
- rename `KFile.GetContents` to `ReadFile`
- rename `KFile.PutContents` to `WriteFile`
- rename `KStr.CamelName` to `ToCamelCase`
- rename `KStr.LowerCaseFirstWords` to `Lcwords`
- rename `KStr.StrShuffle` to `Shuffle`
- rename `KStr.Strrev` to `Reverse`
- rename `KStr.UpperCaseFirstWords` to `Ucwords`
- rename `KTime.Strtotime` to `Str2Timestamp`

#### Removed

- remove `KConv.Int2Bool`

*--end of file--*