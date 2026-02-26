package controller

import (
	"intern_template_v1/model/errors"
	"intern_template_v1/model/response"
	"intern_template_v1/model/status"

	"github.com/gofiber/fiber/v2"
)

func Get_names(c *fiber.Ctx) error {

	result, err := get_All_Names()
	if err != nil {
		return c.Status(500).JSON(response.ResponseModel{
			RetCode: "500",
			Message: status.RetCode500,
			Data: errors.ErrorModel{
				Message:   "Failed to fetch names",
				IsSuccess: false,
				Error:     err,
			},
		})
	}

	return c.JSON(result)
}

func Insert_name(c *fiber.Ctx) error {

	params := map[string]any{}

	if err := c.BodyParser(&params); err != nil {
		return c.Status(401).JSON(response.ResponseModel{
			RetCode: "401",
			Message: status.RetCode401,
			Data: errors.ErrorModel{
				Message:   "Failed to parse request",
				IsSuccess: false,
				Error:     err,
			},
		})
	}

	result, err := insert_student_name(params)
	if err != nil {
		return c.Status(500).JSON(response.ResponseModel{
			RetCode: "500",
			Message: status.RetCode500,
			Data: errors.ErrorModel{
				Message:   "failed to fetch student name",
				IsSuccess: false,
				Error:     err,
			},
		})
	}

	return c.JSON(result)
}
