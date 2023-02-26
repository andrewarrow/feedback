package controller

type Vars struct {
	Title string
	Phone string
}

func NewVars(site *Site) *Vars {
	v := Vars{}
	v.Title = "Feedback"
	v.Phone = site.Phone
	return &v
}

func (v *Vars) Fill(r *Render) {
	v.Title = r.Vars.Title
	v.Phone = r.Vars.Phone
}
