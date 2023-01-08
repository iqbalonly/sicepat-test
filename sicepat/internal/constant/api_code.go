package constant

type Code int

const (
	Unknown             Code = 99
	Success             Code = 0
	BadRequestGeneral   Code = 100
	InternalServerError Code = 200
)

const (
	DateReturnPayload string = "2006-01-02 15:04:05"
)
