package v1

import (
	"Bukus/internal/entity"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func (obj *V1) store(ctx fiber.Ctx) error {
	uc := obj.uc.Book

	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		return errorResponse(ctx, http.StatusUnauthorized, "unauthorized")
	}

	err := uc.Store(ctx.RequestCtx(), userID, entity.Book{})
	if err != nil {
		obj.l.Error(err, "restapi - v1 - history")

		return errorResponse(ctx, http.StatusInternalServerError, "database problems")
	}

	return ctx.Status(http.StatusOK).JSON(nil)
}
