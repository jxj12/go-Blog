package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"

	//"log"
	//"github.com/EDDYCJY/go-gin-example/models"
	"./gin-blog/models"
	//"time"
	//"fmt"
)

//定时操作


func main(){
	fmt.Println("starting....")
	c :=cron.New()
	fmt.Println("111111111111")
//如果按月执行等，和修改spec
	c.AddFunc("*/5*****", func() {
		fmt.Println("Run models.CleanAllTag...")
		//models.CleanAllTag()
		models.ClentAllTage()
		fmt.Println("44444444444")
	})
	c.AddFunc("*/15*****", func() {
		fmt.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.Start()
	fmt.Println("222222222222222")
	//defer c.Stop()

	t1:=time.NewTimer(time.Second*10) //会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	//阻塞 select 等待 channel
	for {
		select {
		case <- t1.C:
			t1.Reset(time.Second*10)//会重置定时器，让它重新开始计时
			fmt.Println("333333333")
		}
	}

}
