package hello

import (
	"encoding/json"
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

	outputVal, _ := json.Marshal(name)

	flow.AppendOutput(ctx, flow.NameValue{
		Name:  "name",
		Value: outputVal,
	})

	return
}
