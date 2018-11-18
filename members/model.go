package members

//Member : struct to represent member
type Member struct {
	ID       string
	Name     string
	Email    string
	Password string
	Enabled  bool
}

// Copy return copy of Member
func (m *Member) Copy() *Member {
	newMember := *m
	return &newMember
}
