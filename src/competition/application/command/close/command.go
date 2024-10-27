package close

import "fmt"

type Command struct {
	//
}

func New() *Command {
	return &Command{
		//
	}
}

func (c *Command) Execute() {
	fmt.Println("Hello, world!")
}
