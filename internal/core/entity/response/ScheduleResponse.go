package response

type Session struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TimeStart string `json:"time_start"`
}

type SessionDetails struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`

	Staff *[]Worker `json:"staff_list" gorm:"one2many"`
}

type Worker struct {
	ID        uint   `json:"id" example:"1"`
	Code      string `json:"code" example:"DR00001"`
	FullName  string `json:"full_name" example:"Alsyad Ahmad"`
	Role      string `json:"roles" example:"Doctor"`
	Facility  string `json:"facility" example:"General"`
	Date		  string `json:"date" example:"2022-07-12"`
	SessionDetailsID		int `json:"-"`
}
