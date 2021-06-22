package entity

type NormalResponse struct {
	Ok string `json:"msg"`
}

type SuperheroesResponse struct {
	Superheroes []*Superheroe `json:"superheroes"`
}

type SuperheroeResponse struct {
	Msg        string      `json:"msg"`
	Superheroe *Superheroe `json:"superheroe"`
}
