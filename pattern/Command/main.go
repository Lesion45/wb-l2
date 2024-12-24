package patterns

import "fmt"

type ICommand interface {
	Execute()
}

type Invoker struct {
	command ICommand
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) SetCommand(c ICommand) {
	i.command = c
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

type PrintCommand struct {
	text string
}

func (pc *PrintCommand) Execute() {
	fmt.Println(pc.text)
}

func NewPrintCommand(text string) *PrintCommand {
	return &PrintCommand{text: text}
}

type AddCommand struct {
	a int
	b int
}

func (ac *AddCommand) Execute() {
	fmt.Println(ac.a + ac.b)
}

func NewAddCommand(a int, b int) *AddCommand {
	return &AddCommand{a: a, b: b}
}

func invoker() {
	invoker := NewInvoker()
	invoker.SetCommand(NewPrintCommand("Hello"))
	invoker.ExecuteCommand()
	invoker.SetCommand(NewAddCommand(1, 2))
	invoker.ExecuteCommand()
}
