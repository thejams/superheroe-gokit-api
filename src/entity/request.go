//Package entity provides all entities for this microservice
package entity

//Superheroe main struct for a superheroe
type Superheroe struct {
	ID        string `json:"id" validate:"omitempty"`
	Name      string `json:"name" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
}

//IDRequest
type IDRequest struct {
	Id string `json:"id"`
}
