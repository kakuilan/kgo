package kgo

// Struct2Map 结构体转为字典;tagName为要导出的标签名,可以为空,为空时将导出所有公开字段.
func (kc *LkkConvert) Struct2Map(obj interface{}, tagName string) (map[string]interface{}, error) {
	return struct2Map(obj, tagName)
}
