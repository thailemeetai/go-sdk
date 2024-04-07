/*
 * @author           Thai Le <thaile.meetai@gmail.com>
 * @copyright        Research Team <thaile.meetai@gmail.com>
 * @license          Apache-2.0
 */

package main

import (
	goservice "github.com/thailemeetai/go-sdk/go-sdk"
	"github.com/thailemeetai/go-sdk/go-sdk/plugin/storage/sdkclickhouse"
)

func main() {
	service := goservice.New(
		goservice.WithName("demo"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkclickhouse.NewClickHouseDB("clickhouse", "")),
	)

	service.Init()
	_ = service.Start()
}
