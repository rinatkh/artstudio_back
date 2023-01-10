package consts

const (
	INQUEUE        = 1
	ALREADYSTUDENT = 2
	DONTANSWER     = 3
	DONTWANT       = 4
)

var Status = map[int]string{
	INQUEUE:        "в ожидании",
	ALREADYSTUDENT: "Уже занимается в ArtStudio",
	DONTANSWER:     "Не берет телефон",
	DONTWANT:       "Не хочет",
}

const (
	Student = "STUDENT"
	Teacher = "TEACHER"
	Admin   = "ADMIN"
)

const (
	MAN     = "M"
	WOMAN   = "W"
	NOTHING = ""
)
