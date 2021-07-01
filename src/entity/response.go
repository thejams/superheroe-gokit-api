package entity

type NormalResponse struct {
	Ok string `json:"msg"`
}

type Superheroes struct {
	Superheroes []*Superheroe `json:"superheroes,omitempty"`
	Error       error         `json:"error,omitempty"`
}

type SuperheroeResponse struct {
	Msg        string      `json:"msg,omitempty"`
	Superheroe *Superheroe `json:"superheroe,omitempty"`
	Error      error       `json:"error,omitempty"`
}
