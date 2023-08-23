package main

import "golang.org/x/text/encoding/simplifiedchinese"

/*
UTF8ToGBK
功能：UTF8 转 GBK
参数：UTF8 编码的字符串
返回值：GBK编码的字符串和错误信息
*/
func UTF8ToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}
	return string(dst[:nDst]), nil
}
