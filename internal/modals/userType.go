package modals

type UserType int

const (
	Student = iota
	Admin
)

func (ut UserType) String() string {
	names := [...]string{"Student", "Admin"}
	if ut < Student || ut > Admin {
		return "Invalid"
	} else {
		return names[ut]
	}
}
