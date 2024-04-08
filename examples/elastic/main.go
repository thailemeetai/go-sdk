package main

import (
	"github.com/sirupsen/logrus"
	goservice "github.com/thailemeetai/go-sdk"
	"github.com/thailemeetai/go-sdk/plugin/storage/sdkes"
)

func main() {
	service := goservice.New(
		goservice.WithName("demo"),
		goservice.WithVersion("1.0.0"),
	)
	_ = service.Init()
	newES := sdkes.NewES("test", "example")
	newES.InitFlags()
	err := newES.Run()
	if err != nil {
		logrus.Error("err: ", err)
	}
	_ = service.Start()
}
