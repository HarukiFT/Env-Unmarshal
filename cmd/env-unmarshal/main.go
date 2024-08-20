// Usage example

package main

import (
	envunmarshal "env-unmarshal"
	"fmt"
	"os"
)

type Test struct {
	Line int `env:"line"`
}

func main() {
	env, err := envunmarshal.LoadEnv(".env")

	if err != nil {
		panic(err)
	}

	var test Test = Test{}
	env.Unmarshal(&test)

	fmt.Println(test.Line)

	env.ToEnv()

	fmt.Println(os.Getenv("line"))

}
