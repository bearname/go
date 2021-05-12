package model

type Originator struct {
	State string
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{e.State}
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.State = m.getSavedState()
}

func (e *Originator) SetState(state string) {
	e.State = state
}

func (e *Originator) GetState() string {
	return e.State
}
