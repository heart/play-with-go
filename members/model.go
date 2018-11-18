package members

//ILogin is interface of member object use to login
type ILogin interface {
	Email() string
	Password() string
}

//Member : struct to represent BNK member
type Member struct {
	Birthday         string   `json:"birthday"`
	BloodType        string   `json:"blood_type"`
	EnglishFirstName string   `json:"english_first_name"`
	EnglishLastName  string   `json:"english_last_name"`
	Facebook         string   `json:"facebook"`
	Height           int      `json:"height"`
	Hobby            string   `json:"hobby"`
	Instagram        string   `json:"instagram"`
	Like             []string `json:"like"`
	Nickname         string   `json:"nickname"`
	Province         string   `json:"province"`
	ThaiFirstName    string   `json:"thai_first_name"`
	ThaiLastName     string   `json:"thai_last_name"`
}

// Copy return copy of Member
func (m *Member) Copy() *Member {
	newMember := *m
	return &newMember
}
