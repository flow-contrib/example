package hello

import (
	"fmt"

	"github.com/gogap/config"
	"github.com/gogap/context"
	"github.com/gogap/flow"
)

func init() {
	flow.RegisterHandler("example.hello", Hello)
}

func Hello(ctx context.Context, conf config.Configuration) (err error) {
	name := conf.GetString("name", "world")
	fmt.Printf("Hello: %s\n", name)

	flow.AppendOutput(ctx, flow.NameValue{Name: "name", Value: name})

	return
}
