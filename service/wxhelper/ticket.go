package wxhelper

type Ticket struct {
	ticket    string
	expiresIn int64
}

func NewTicket(ticket string, expiresIn int64) Ticket {
	return Ticket{
		ticket:    ticket,
		expiresIn: expiresIn,
	}
}

func (t *Ticket) getExpiresIn() int64 {
	return t.expiresIn
}

func (t *Ticket) getTicket() string {
	return t.ticket
}
