package main

type Project interface {
	Name() string
	Tickets() []Ticket
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
	tickets []Ticket
}

type StatusInMemory struct {
	name string
}

func (p *ProjectInMemory) Name() string {
	return p.name
}

func (p *ProjectInMemory) Tickets() []Ticket {
	return p.tickets
}

type TicketInMemory struct {
	name   string
	id     string
	status *StatusInMemory
}

func (t *TicketInMemory) Name() string {
	return t.name
}

func (t *TicketInMemory) ID() string {
	return t.id
}

func (t *TicketInMemory) Status() Status {
	return t.status
}

func (s *StatusInMemory) Name() string {
	return s.name
}
