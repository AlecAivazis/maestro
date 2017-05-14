package main

type Project interface {
	Name() string
	Tickets() []*Ticket
}

type Status interface {
	Name() string
}

type Ticket interface {
	Name() string
	ID() string
	Status() Status
}

type ProjectInMemory struct {
	name    string
	tickets []*Ticket
}

func (p *ProjectInMemory) Name() string {
	return p.name
}

func (p *ProjectInMemory) Tickets() []*Ticket {
	return p.tickets
}
