package command

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)
	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}
		p.queue = make([]Command, 3)
	}
}
