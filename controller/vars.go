package controller

type Vars struct {
	Title  string
	Header string
	Footer string
	Phone  string
}

func NewVars() Vars {
	v := Vars{}
	v.Title = "Feedback"
	return v
}
