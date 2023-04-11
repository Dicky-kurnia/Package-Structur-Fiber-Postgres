package exception

import (
	"errors"
	"go-fiber-postgres/model"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

type CustomErr struct {
	status  int
	code    string
	message string
}

func (c *CustomErr) Status() int {
	return c.status
}

func (c *CustomErr) Code() string {
	return c.code
}

func (c *CustomErr) Error() string {
	return c.message
}

func NewError(status int, code, message string) error {
	return &CustomErr{
		status:  status,
		code:    code,
		message: message,
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var customError *CustomErr
	if errors.As(err, &customError) {
		return ctx.Status(customError.Status()).JSON(model.Response{
			Code:   customError.Status(),
			Status: customError.Code(),
			Error: map[string]interface{}{
				"general": customError.Error(),
			},
		})
	}
	_, ok := err.(ValidationError)
	if ok {
		var obj interface{}
		_ = json.Unmarshal([]byte(err.Error()), &obj)
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: model.BAD_REQUEST,
			Data:   nil,
			Error:  obj,
		})
	}

	if err.Error() == model.NOT_FOUND_ERR_TYPE {
		return ctx.Status(404).JSON(model.Response{
			Code:   404,
			Status: "Not Found",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.NOT_FOUND_ERR_TYPE,
			},
		})
	}

	if err.Error() == model.PRODUCT_NOT_FOUND {
		return ctx.Status(404).JSON(model.Response{
			Code:   404,
			Status: "Not Found",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.PRODUCT_NOT_FOUND,
			},
		})
	}

	if err.Error() == model.CART_NOT_FOUND {
		return ctx.Status(404).JSON(model.Response{
			Code:   404,
			Status: "Not Found",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.CART_NOT_FOUND,
			},
		})
	}

	if err.Error() == model.AUTHENTICATION_FAILURE_ERR_TYPE {
		return ctx.Status(401).JSON(model.Response{
			Code:   401,
			Status: "Bad Request",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.AUTHENTICATION_FAILURE_ERR_TYPE,
			},
		})
	}

	if err.Error() == model.AGENT_BLOCKED {
		return ctx.Status(401).JSON(model.Response{
			Code:   401,
			Status: "Agent Blocked",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.AGENT_BLOCKED,
			},
		})
	}

	if err.Error() == model.AGENT_INACTIVE {
		return ctx.Status(401).JSON(model.Response{
			Code:   401,
			Status: "Agent Inactive",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.AGENT_INACTIVE,
			},
		})
	}

	if err.Error() == model.AGENT_NOT_FOUND {
		return ctx.Status(404).JSON(model.Response{
			Code:   404,
			Status: "Agent Not Found",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.AGENT_NOT_FOUND,
			},
		})
	}

	if err.Error() == model.ORDER_NOT_FOUND {
		return ctx.Status(404).JSON(model.Response{
			Code:   404,
			Status: "Not Found",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.ORDER_NOT_FOUND,
			},
		})
	}

	if err.Error() == model.UNAUTHORIZATION {
		return ctx.Status(401).JSON(model.Response{
			Code:   401,
			Status: model.UNAUTHORIZATION,
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.UNAUTHORIZATION,
			},
		})
	}
	if err.Error() == model.UNAUTHORIZED {
		return ctx.Status(401).JSON(model.Response{
			Code:   401,
			Status: "Unauthorized",
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.UNAUTHORIZED,
			},
		})
	}

	if err.Error() == model.USERNAME_OR_PASSWORD_INVALID {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: model.BAD_REQUEST,
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.USERNAME_OR_PASSWORD_INVALID,
			},
		})
	}

	if err.Error() == model.PASSWORD_INVALID {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: model.BAD_REQUEST,
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.PASSWORD_INVALID,
			},
		})
	}

	if err.Error() == model.CONFIRM_PASSWORD_NOT_MATCH {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: model.BAD_REQUEST,
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.CONFIRM_PASSWORD_NOT_MATCH,
			},
		})
	}

	if err.Error() == model.OLD_PASSWORD_CANNOT_BE_MATCH {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: model.BAD_REQUEST,
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.OLD_PASSWORD_CANNOT_BE_MATCH,
			},
		})
	}

	if err.Error() == model.EXTENSION_NOT_ALLOWED {
		return ctx.Status(http.StatusBadRequest).JSON(model.Response{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.EXTENSION_NOT_ALLOWED,
			},
		})
	}

	if err.Error() == model.STOCK_EXCEEDES_LIMIT {
		return ctx.Status(http.StatusBadRequest).JSON(model.Response{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.STOCK_EXCEEDES_LIMIT,
			},
		})
	}

	if err.Error() == model.MUST_NUMBER_ERR_TYPE {
		return ctx.Status(http.StatusBadRequest).JSON(model.Response{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.MUST_NUMBER_ERR_TYPE,
			},
		})
	}

	return ctx.Status(500).JSON(model.Response{
		Code:   500,
		Status: model.INTERNAL_ERROR_ERR_TYPE,
		Data:   err.Error(),
	})
}
