package router

import (
	"analysis_cloud/handler"
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/faygo/ext/middleware"
)

func Route(frame *faygo.Framework) {
	frame.Route(
		frame.NewNamedAPI("upload", "POST", "/upload", handler.UploadFile()).Use(middleware.CrossOrigin),
		//frame.NewNamedAPI("upload", "GET POST", "/view/upload", handler.Upload()),
		//frame.NewNamedAPI("fang kai  index", "GET POST", "/view/index", handler.Index()),
		//frame.NewNamedAPI("index Content", "GET POST", "/view/indexContent", handler.IndexContent()),
	)
}
