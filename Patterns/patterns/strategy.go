package patterns

import (
	"fmt"
)

type Strategy interface {
	Execute()
}

type Context struct {
	strat Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strat = strategy
}

func (c *Context) ExecuteStrategy() {
	if c.strat != nil {
		c.strat.Execute()
	}
}

type Strategy1 struct{}

func (s Strategy1) Execute() {
	fmt.Println("Strategy1 executed")
}

type Strategy2 struct{}

func (s Strategy2) Execute() {
	fmt.Println("Strategy2 executed")
}

type Strategy3 struct{}

func (s Strategy3) Execute() {
	fmt.Println("Strategy3 executed")
}

func main_strategy() {
	fmt.Println()

	ctx := &Context{}

	ctx.SetStrategy(Strategy1{})
	ctx.ExecuteStrategy()

	ctx.SetStrategy(Strategy2{})
	ctx.ExecuteStrategy()

	ctx.SetStrategy(Strategy3{})
	ctx.ExecuteStrategy()

	fmt.Println()
}
