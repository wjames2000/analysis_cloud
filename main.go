/**
说明：主运行单元
作者：王剑
QQ : 87225511
 */
package main

import (
	"github.com/henrylee2cn/faygo"
	"analysis_cloud/router"
)

func main(){
	faygo.SetUpload("./upload", false, false)
	app := faygo.New("analysis_cloud", "1.0")
	router.Route(app)
	faygo.Run()
}