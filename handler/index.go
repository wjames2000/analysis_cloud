package handler

import (
	"github.com/henrylee2cn/faygo"
	"fmt"
	"github.com/henrylee2cn/faygo/ext/uuid"
	"strings"
	u "analysis_cloud/units"
)


type UploadJson struct {
	Code      int64  `json:"code"`
	Msg       string `json:"msg"`
	AccessUrl string `json:"access_url"`
}

/**
* @Purpose:
* 文件上传
* @Method Name: UploadFile()
* @Param: 无
* @Author: 王剑
* @Return: code 错误代码
* 	   msg  错误信息
 */
func UploadFile() faygo.HandlerFunc {
	return func(c *faygo.Context) error {
		_, fh, err := c.FormFile("file")
		faygo.Debugf("upload file is fail, error info %s, fn %s", err, fh.Filename)
		if err != nil {
			faygo.Debugf("upload file is fail, error info %s", err)
			code, msg := u.AdapterError(err)
			return c.String(500, fmt.Sprintf(`"code":%d,"msg":"%s"`, code, msg))
		}

		fn := fmt.Sprintf("%s%s", uuid.New().String(), u.GetFileExt(fh.Filename))
		fn = strings.Replace(fn, "-", "", -1)
		_, err = c.SaveFile("file", false, fn)
		if err != nil {
			faygo.Debugf("Save file fail，error info : %v\n", err)
		}
		//判断文件后缀名，amr文件需要转换成mp3再上传
		//if u.GetFileExt(fn) == ".amr" {
		//	err, fn = u.ConvertToMP3(fmt.Sprintf("%s%s", faygo.UploadDir(), fn))
		//	if err != nil {
		//		faygo.Debugf("Convert amr file To MP3 is error, error info:%s", err)
		//	}
		//}
		//上传到阿里云Oss
		//fname := u.PutObject(fmt.Sprintf("%s%s", faygo.UploadDir(), fn), fn)
		//u.DeleteUploadFile(fn)
		uj := new(UploadJson)
		uj.Code = 0
		uj.Msg = "success"
		uj.AccessUrl = "http://" + fn
		return c.JSON(200, uj, false)
	}
}

