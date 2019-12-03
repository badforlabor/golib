/**
 * Auth :   liubo
 * Date :   2019/12/3 9:19
 * Comment: 文件
 */

package io

import (
	"golib/log"
	"os"
)

// 追加内容
func AppendFile(filename string, text string) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		log.Println(err)
	}
}
