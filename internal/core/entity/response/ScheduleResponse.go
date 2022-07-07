package response

type ScheduleList struct {
	DataCheck   string `json:"date_check"`
	UserCode    string `json:"doctor_code"`
	SessionList []SessionList
}

type SessionList struct {
	SessionID uint   `json:"session_id"`
	Code      string `json:"code"`
}
