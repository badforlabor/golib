/**
 * Auth :   liubo
 * Date :   2019/12/3 9:19
 * Comment: 文件
 */

package ultraio

import (
	"golib/ultralog"
	"os"
)

// 追加内容
func AppendFile(filename string, text string) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		ultralog.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		ultralog.Println(err)
	}
}
