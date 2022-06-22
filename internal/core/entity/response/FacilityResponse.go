package response

type Facility struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type FacilityDetails struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Staff []Staff `json:"staff"`
}

type Staff struct {
	Code string `json:"code"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
}
