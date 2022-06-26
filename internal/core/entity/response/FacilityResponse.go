package response

type Facility struct {
	ID int `json:"id" example:"1"`
	Name string `json:"name" example:"General"`
}

type FacilityDetails struct {
	ID int `json:"id" example:"1"`
	Name string `json:"name" example:"General"`
	Staff []Staff `json:"staff"`
}

type Staff struct {
	Code string `json:"code" example:"DR00001"`
	FullName string `json:"full_name" example:"Alsyad Ahmad"`
	Role string `json:"role" example:"Doctor"`
}
