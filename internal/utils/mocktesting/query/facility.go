package query


func NewFacility() string {
	return "INSERT INTO `medical_facilities` (`name`) VALUES (?)"
}


func UpdateFacility() string {
	return "UPDATE `medical_facilities` SET `name`=? WHERE `id` = ?"
}

func DeleteFacility() string {
	return "DELETE FROM `medical_facilities` WHERE `medical_facilities`.`id` = ?"
}

func FindAllFacility() string {
	return "SELECT * FROM `medical_facilities`"
}

func FindByIDFacility() string {
	return "SELECT * FROM `medical_facilities` WHERE id = ?"
}

func FindFacilityMember() string {
	return "SELECT users.*, roles.name as role FROM `users` join roles on roles.id = users.role_id "+
		"WHERE medical_facility_id = ? AND `users`.`deleted_at` IS NULL"
}
