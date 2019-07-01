package main

import (
	"context"
	"encoding/json"
	"github.com/TonyXMH/MicroExample/micro-api/api/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api/proto"
	"go-micro/errors"
	"strings"
)

//主要实现我们在proto中定义的接口函数一个是Call和一个Bar

//首先定义Example和Foo两个空struct作为实现Call和Bar的载体

type Example struct{}
type Foo struct{}

//实现api.micro.go中ExampleHandler interface接口
func (e *Example) Call(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	log.Log("Example.Call")
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "arg error")
	}
	for k, v := range req.Header {
		log.Log("req header ", k, " : ", v)
	}
	rsp.StatusCode = 200
	b,_:=json.Marshal(map[string]string{
		"message":"we have received your req "+ strings.Join(name.Values," "),
	})
	rsp.Body = string(b)
	return nil
}

//实现api.micro.go中FooHandle interface接口
func (f *Foo) Bar(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	log.Log("Foo.Bar")
	if req.Method != "POST"{
		return errors.BadRequest("go.micro.api.example","require post")
	}
	ct,ok:=req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0{
		return  errors.BadRequest("go.micro.api.example","need content-type")
	}
	if ct.Values[0] != "application/json"{
		return errors.BadRequest("go.micro.api.example","expect application/json")
	}

	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body),&body)
	rsp.Body="received message "+ req.Body
	return nil
}

func main() {
	service := micro.NewService(micro.Name("go.micro.api.example"))
	service.Init()
	//调用api.micro.go中register注册函数
	api.RegisterExampleHandler(service.Server(), new(Example))
	api.RegisterFooHandler(service.Server(), new(Foo))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//运行顺序
// micro api --handler=api
// go run api.go
//构造get和post请求
//curl -H 'head-1: I am a header' "http://localhost:8080/example/call?name=john"
//curl -H 'Content-Type: application/json' -d '{data:1233}' http://localhost:8080/example/foo/bar