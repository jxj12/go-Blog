package gredis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"

	"gin-blog/pkg/setting"
)

var RedisConn *redis.Pool

func Setup()error{
//	建立连接池
	RedisConn = &redis.Pool{
		//DialContext:     nil,
		MaxIdle:         setting.RdisSetting.MaxIdle,//最大空闲连接数
		MaxActive:       setting.RdisSetting.MaxActive, //在给定时间内，允许分配的最大连接数（当为零时，没有限制）
		IdleTimeout:     setting.RdisSetting.IdleTimeout,//在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
		//Wait:            false,
		//MaxConnLifetime: 0,
		Dial:            func()(redis.Conn,error){ //提供创建和配置应用程序连接的一个函数
			//连接redis
			c,err :=redis.Dial("tcp",setting.RdisSetting.Host)
			if err !=nil{
				return nil,err
			}
			if setting.RdisSetting.Password !=""{
				if _,err :=c.Do("AUTH",setting.RdisSetting.Password);err!=nil{
					c.Close()
					return nil,err
				}
			}
			return c,err
		},
		TestOnBorrow:  func(c redis.Conn,t time.Time)error{
			_,err :=c.Do("PING")
			return err
		},
	}
	return nil
}



func Set(key string,data interface{},time int)(bool,error){
	conn :=RedisConn.Get() //获取链接池
	defer conn.Close()
	value,err :=json.Marshal(data)//将数据编码成json字符串
	if err !=nil{
		return false,err
	}
	reply,err :=redis.Bool(conn.Do("SET",key,value))//Bool是将命令回复转换为布尔值的助手
	conn.Do("EXPIRE",key,time)
	return reply,err
}
//检测值是否存在
func Exists(key string)bool{
	conn :=RedisConn.Get()
	defer conn.Close()
	reply,err :=redis.Bool(conn.Do("EXISTS",key))
	if err !=nil{
		return false
	}
	return reply

}

func Get(key string)([]byte,error){
	conn :=RedisConn.Get()
	defer  conn.Close()
	reply,err:=redis.Bytes(conn.Do("GET",key))
	if err != nil{
		return nil,err}
	return reply,err
}

func Delect(key string)([]byte,error){
	conn :=RedisConn.Get()
	defer conn.Close()
	reply,err :=redis .Bytes(conn.Do("DEL",key))
	if err !=nil{
		return nil,err
	}
	return reply,err
}

func LikeDeletes(key string)error{
	conn :=RedisConn.Get()//在连接池中获取一个活跃连接
	defer  conn.Close()
	keys,err :=redis.Strings(conn.Do("keys","*"+key+"*"))
	if err !=nil{
		return err
	}
	for _,key :=range keys{
		_,err=Delect(key)
		if err !=nil{
			return err
		}
	}
	return nil
}