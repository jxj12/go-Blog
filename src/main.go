package main

import (
	"./gin-blog/pkg/setting"
	"./gin-blog/routers"
	"fmt"
	//"github.com/EDDYCJY/go-gin-example/models"
	//"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"./gin-blog/pkg/logging"
	"./gin-blog/models"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	r :=routers.InitRouter()
	//r.Run() // listen and serve on 0.0.0.0:8080
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}


	s.ListenAndServe()


}
