package pkg

type CommandInfo struct {
	Name        string
	Description string

	Maker func() CustomCommand2 `json:"-"`
}

type SimpleCommand interface {
	Trigger([]string, MessageEvent) Actions
}

type CustomCommand2 interface {
	SimpleCommand

	HasCooldown(User) bool
	AddCooldown(User)
}

type CommandMatcher interface {
	Register(aliases []string, command interface{}) interface{}
	Deregister(command interface{})
	DeregisterAliases(aliases []string)
	Match(text string) (interface{}, []string)
}

type CommandsManager interface {
	CommandMatcher
	OnMessage(event MessageEvent) Actions

	FindByCommandID(id int64) interface{}

	Register2(id int64, aliases []string, command interface{})
}
