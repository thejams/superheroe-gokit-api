//Package entity provides all entities for this microservice
package entity

//Superheroe main struct for a superheroe
type Superheroe struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

//IDRequest
type IDRequest struct {
	Id string `json:"id"`
}
