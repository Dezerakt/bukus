package v1

import (
	"bukus/internal/entity"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func (obj *V1) store(ctx fiber.Ctx) error {

	err := obj.bookUC.Store(ctx.RequestCtx(), entity.Book{})
	if err != nil {
		obj.l.Error(err, "restapi - v1 - history")

		return errorResponse(ctx, http.StatusInternalServerError, "database problems")
	}

	return ctx.Status(http.StatusOK).JSON(nil)
}
