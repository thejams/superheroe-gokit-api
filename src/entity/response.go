package entity

type NormalResponse struct {
	Ok string `json:"msg"`
}

type Superheroes struct {
	Superheroes []*Superheroe `json:"superheroes"`
	Error       error         `json:"error" validate:"omitempty"`
}

type SuperheroeResponse struct {
	Msg        string      `json:"msg"`
	Superheroe *Superheroe `json:"superheroe"`
	Error      error       `json:"error" validate:"omitempty"`
}
