package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/robbailey3/website-api/exception"
)

func BadRequest(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(struct{ BaseResponse }{
		BaseResponse: BaseResponse{
			Timestamp: time.Now().Unix(),
			Error: &ErrorResponse{
				Code:    exception.BAD_REQUEST,
				Message: msg,
			},
		},
	})
}
