package entity

type GetIDRequest struct {
	Id string `json:"id"`
}

type NormalResponse struct {
	Ok string `json:"ok"`
}

type SuperheroesResponse struct {
	Superheroes []*Superheroe `json:"superheroes"`
}

type SuperheroeResponse struct {
	Msg        string      `json:"msg"`
	Superheroe *Superheroe `json:"superheroe"`
}

type EmptyStruct struct{}
