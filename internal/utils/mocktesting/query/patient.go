package query


func NewPatient() string {
	return "INSERT INTO `patients` (`code`,`national_id`,`full_name`,`address`,"+
			"`gender`,`birth_date`,`blood_type`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?)"
}


func UpdatePatient() string {
	return "UPDATE `patients` SET `national_id`=?,`full_name`=?,"+
		"`address`=?,`gender`=?,`birth_date`=?,`blood_type`=? "+
		"WHERE `patients`.`deleted_at` IS NULL AND `id` = ?"
}

func DeletePatient() string {
	return "UPDATE `patients` SET `deleted_at`=? WHERE `patients`.`id` = ? AND `patients`.`deleted_at` IS NULL"
}

func FindAllPatient() string {
	return "SELECT * FROM `patients` WHERE `patients`.`deleted_at` IS NULL"
}

func FindByIDPatient() string {
	return "SELECT * FROM `patients` WHERE id = ? AND `patients`.`deleted_at` IS NULL"
}

func FindPatientMedicRecord() string {
	return "SELECT MedicRecord.*,medical_sessions.*,MC.blood_tension AS MedicCheck__blood_tension,"+
		"MC.height AS MedicCheck__height,MC.weight AS MedicCheck__weight,MC.body_temperature AS MedicCheck__body_temperature,"+
		"U.full_name AS MedicCheck__nurse,User.full_name as doctor,MedicalFacility.name as facility,"+
		"`MedicRecord`.`id` AS `MedicRecord__id`,`MedicRecord`.`serial_number` AS `MedicRecord__serial_number`,"+
		"`MedicRecord`.`complaint` AS `MedicRecord__complaint`,`MedicRecord`.`diagnose` AS `MedicRecord__diagnose`,"+
		"`MedicRecord`.`prescription` AS `MedicRecord__prescription`,`MedicRecord`.`patient_code` AS `MedicRecord__patient_code`,"+
		"`MedicRecord`.`status` AS `MedicRecord__status`,`User`.`id` AS `User__id`,`User`.`code` AS `User__code`,"+
		"`User`.`full_name` AS `User__full_name`,`User`.`gender` AS `User__gender`,`User`.`email` AS `User__email`,"+
		"`User`.`password` AS `User__password`,`User`.`status` AS `User__status`,`User`.`role_id` AS `User__role_id`,"+
		"`User`.`medical_facility_id` AS `User__medical_facility_id`,`User`.`deleted_at` AS `User__deleted_at`,"+
		"`MedicalFacility`.`id` AS `MedicalFacility__id`,`MedicalFacility`.`name` AS `MedicalFacility__name` "+
		"FROM `medical_sessions` LEFT JOIN `medic_records` `MedicRecord` ON `medical_sessions`.`medic_record_id` = `MedicRecord`.`id` "+
		"LEFT JOIN `users` `User` ON `medical_sessions`.`user_code` = `User`.`code` AND `User`.`deleted_at` IS NULL "+
		"LEFT JOIN `medical_facilities` `MedicalFacility` ON `medical_sessions`.`medical_facility_id` = `MedicalFacility`.`id` "+
		"LEFT JOIN medic_checks MC ON MC.medic_record_id = MedicRecord.id LEFT JOIN users U ON U.code = MC.user_code "+
		"WHERE MedicRecord.patient_code = ? AND MedicRecord.status = 1"
}
