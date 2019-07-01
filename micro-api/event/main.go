package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api/proto"
	"github.com/micro/go-log"
)

//事件订阅对象的所有公有方法都会执行，只要方法的参数数ctx和event

type Event struct{}

func (e*Event)Fun1(ctx context.Context,event*go_api.Event)error  {
	log.Log("Fun1 coming Event Name ",event.Name)
	log.Log("Fun1 coming Event Data ",event.Data)
	return nil
}
func (e*Event)Fun2(ctx context.Context,event*go_api.Event)error  {
	log.Log("Fun2 coming Event Name ",event.Name)
	log.Log("Fun2 coming Event Data ",event.Data)
	return nil
}
func (e*Event)Fun3(ctx context.Context,event*go_api.Event)error  {
	log.Log("Fun3 coming Event Name ",event.Name)
	log.Log("Fun3 coming Event Data ",event.Data)
	return nil
}
//私有函数是不会执行的
func (e*Event)fun1(ctx context.Context,event*go_api.Event)error  {
	log.Log("private fun1 coming Event Name ",event.Name)
	log.Log("private fun1 coming Event Data ",event.Data)
	return nil
}

func main()  {
	service:=micro.NewService(micro.Name("user"))
	service.Init()
	micro.RegisterSubscriber("go.micro.evt.user",service.Server(),new(Event))
	if err:=service.Run();err!=nil{
		log.Fatal(err)
	}
}

//micro api --handler=event --namespace=go.micro.evt
//go run main.go
//curl -d '{"message":"hello micro"}' http://localhost:8080/user/login