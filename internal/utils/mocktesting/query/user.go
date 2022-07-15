package query

func FindUserByEmail() string {
	return "SELECT users.*, Role.name as role, MedicalFacility.name as facility,`Role`.`id` AS `Role__id`,"+
		"`Role`.`name` AS `Role__name`,`Role`.`code` AS `Role__code`,"+
		"`MedicalFacility`.`id` AS `MedicalFacility__id`,`MedicalFacility`.`name` AS `MedicalFacility__name` FROM `users` "+
		"LEFT JOIN `roles` `Role` ON `users`.`role_id` = `Role`.`id` LEFT JOIN `medical_facilities` `MedicalFacility` "+
		"ON `users`.`medical_facility_id` = `MedicalFacility`.`id` WHERE users.email = ? AND `users`.`deleted_at` IS NULL"
}

func NewUser() string {
	return "INSERT INTO `users` (`code`,`full_name`,`gender`,`email`,`password`,`status`,`role_id`,`medical_facility_id`,`deleted_at`) " +
		"VALUES (?,?,?,?,?,?,?,?,?)"
}

func FindUserByCode() string {
	return "SELECT users.*, Role.name as role, MedicalFacility.name as facility,`Role`.`id` AS `Role__id`,"+
		"`Role`.`name` AS `Role__name`,`Role`.`code` AS `Role__code`,`MedicalFacility`.`id` AS `MedicalFacility__id`,"+
		"`MedicalFacility`.`name` AS `MedicalFacility__name` FROM `users` LEFT JOIN `roles` `Role` "+
		"ON `users`.`role_id` = `Role`.`id` LEFT JOIN `medical_facilities` `MedicalFacility` "+
		"ON `users`.`medical_facility_id` = `MedicalFacility`.`id` WHERE users.code = ? AND `users`.`deleted_at` IS NULL"
}

func ChangePassword() string {
	return "UPDATE `users` SET `code`=?,`password`=? WHERE code = ? AND `users`.`deleted_at` IS NULL"
}

func UpdateUser() string {
	return "UPDATE `users` SET `full_name`=?,`gender`=?,`email`=?,`password`=?,`role_id`=?,`medical_facility_id`=? "+
		"WHERE `users`.`deleted_at` IS NULL AND `id` = ?"
}

func DeleteUser() string {
	return "UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL"
}

func FindAllUser() string {
	return "SELECT users.*,Role.name as role,MedicalFacility.name as facility,`Role`.`id` AS `Role__id`,"+
		"`Role`.`name` AS `Role__name`,`Role`.`code` AS `Role__code`,`MedicalFacility`.`id` AS `MedicalFacility__id`,"+
		"`MedicalFacility`.`name` AS `MedicalFacility__name` FROM `users` LEFT JOIN `roles` `Role` "+
		"ON `users`.`role_id` = `Role`.`id` LEFT JOIN `medical_facilities` `MedicalFacility` "+
		"ON `users`.`medical_facility_id` = `MedicalFacility`.`id` WHERE `users`.`deleted_at` IS NULL"
}

func FindUserByID() string {
	return "SELECT users.*, Role.name as role, MedicalFacility.name as facility,`Role`.`id` AS `Role__id`,"+
		"`Role`.`name` AS `Role__name`,`Role`.`code` AS `Role__code`,`MedicalFacility`.`id` AS `MedicalFacility__id`,"+
		"`MedicalFacility`.`name` AS `MedicalFacility__name` FROM `users` LEFT JOIN `roles` `Role` "+
		"ON `users`.`role_id` = `Role`.`id` LEFT JOIN `medical_facilities` `MedicalFacility` "+
		"ON `users`.`medical_facility_id` = `MedicalFacility`.`id` WHERE users.id = ? AND `users`.`deleted_at` IS NULL"
}
