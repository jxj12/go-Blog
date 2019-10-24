package main

import (
	"./gin-blog/pkg/setting"
	"./gin-blog/routers"
	"fmt"
	"net/http"
)

func main() {

	r :=routers.InitRouter()
	//r.Run() // listen and serve on 0.0.0.0:8080
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()


}
