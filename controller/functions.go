package controller

import (
	"intern_template_v1/middleware"
	"intern_template_v1/sharedfunctions"
)

func get_All_Names() (map[string]any, error) {
	db := middleware.DBConn

	var result map[string]any
	if err := db.Raw(`SELECT sample_feature1.get_all_names()`).Scan(&result).Error; err != nil {
		return nil, err
	}

	sharedfunctions.ConvertStringToJSONMap(result)
	result = sharedfunctions.GetMap(result, "get_all_names")

	return result, nil
}

func insert_student_name(params map[string]any) (map[string]any, error) {
	db := middleware.DBConn

	var result map[string]any
	if err := db.Raw(`SELECT sample_feature1.insert_name(?)`, params).Scan(&result).Error; err != nil {
		return nil, err
	}

	sharedfunctions.ConvertStringToJSONMap(result)
	result = sharedfunctions.GetMap(result, "insert_name")

	return result, nil
}
