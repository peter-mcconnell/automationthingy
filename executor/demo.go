package executor

import "fmt"

type DemoExecutor struct{}

func (e *DemoExecutor) Execute(id string) {
	for i := 0; i < 100; i++ {
		fmt.Printf(id+" executing line %d\n", i)
	}
}
