package units

import (
	"fmt"
	"path"
)

/**
* @Purpose:
* 适配错误信息
* @Method Name: adapterError()
* @Param: err 错误信息
* @Author: 王剑
* @Return: code 错误代码
* 	   msg  错误信息
 */
func AdapterError(err error) (code int, msg string) {
	if err != nil {
		s := fmt.Sprintf("%s", err)
		return 500, s
	}
	return 0, "success"
}

/**
* @Purpose:
* 获取文件后缀名
* @Method Name: GetFileExt()
* @Param: fn 文件名
* @Author: 王剑
* @Return: 带'.'的后缀名
**/
func GetFileExt(fn string) string {
	filenameWithSuffix := path.Base(fn)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	return fileSuffix
}
