Example
=======

## hello

The config file of `hello.conf`

```hocon
packages = ["github.com/flow-contrib/example"]

app {
	name = "hello"
	usage = "This is a demo for run flow"

	commands {
		say {
			usage = "This command will print hello"

			default-config = { name = "gogap" }

			flow = ["example.hello", "example.hello@confA"]
			
			config = {
    		  confA = {
    			 name = "Zeal"
            }
		}
	}
}
```

```bash
go-flow run say --config hello.conf
```