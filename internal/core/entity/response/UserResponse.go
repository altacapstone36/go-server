package response

type (
	User struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		FullName string `json:"full_name"`
		Password string `json:"-"`
		Gender string `json:"gender"`
		Role string `json:"roles"`
		Facility string `json:"facility"`
		Status int `json:"status"`
	}
)
