package threads

type Metadata struct {
	id       int
	parentID int
	main     bool
}

func (m *Metadata) ID() int       { return m.id }
func (m *Metadata) ParentID() int { return m.id }
func (m *Metadata) Main() bool    { return m.main }

func (m *Metadata) SetParentID(parentID int) { m.parentID = parentID }

func NewMetadata(id int, parentID int) *Metadata {
	return &Metadata{
		id:       id,
		parentID: parentID,
	}
}
