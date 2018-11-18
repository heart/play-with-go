package members

//ILogin is interface of member object use to login
type ILogin interface {
	Email() string
	Password() string
}

//Member : struct to represent member
type Member struct {
	ID           string
	Name         string
	UserEmail    string
	UserPassword string
	Enabled      bool
}

//Email of user
func (m *Member) Email() string {
	return m.UserEmail
}

//Password of user
func (m *Member) Password() string {
	return m.UserPassword
}

// Copy return copy of Member
func (m *Member) Copy() *Member {
	newMember := *m
	return &newMember
}
