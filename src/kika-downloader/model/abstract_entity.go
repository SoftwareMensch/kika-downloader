package model

type abstractEntity struct {
	id string
}

// GetId get entity id
func (a *abstractEntity) GetId() string {
	return a.id
}
