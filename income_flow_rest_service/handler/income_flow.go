package handler

import (
	"github.com/gofiber/fiber/v3"
	jsoniter "github.com/json-iterator/go"
	"income_flow_rest_service/model"
)

var state = []*model.Good{
	{
		ID:          1,
		Name:        "Iphone 16 pro",
		Description: "Новый айфон",
		Count:       1,
		Weight:      0.2,
	},
	{
		ID:          2,
		Name:        "Dyson v12",
		Description: "Пылесос от дайсон",
		Count:       2,
		Weight:      3.0,
	},
}

func (h *Handler) GetOneGood(c fiber.Ctx) error {
	var good model.Good
	err := jsoniter.Unmarshal(c.Body(), &good)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(model.Error{
			Status:  fiber.StatusUnprocessableEntity,
			Code:    "marshal_error",
			Message: err.Error(),
		})
	}

	var sendBody *model.Good
	for _, item := range state {
		if item.ID == good.ID {
			sendBody = item
		}
	}
	if sendBody == nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Error{
			Status:  fiber.StatusNotFound,
			Code:    "there_is_no_good",
			Message: "the good was not found",
		})
	}

	return c.JSON(sendBody)
}

func (h *Handler) GetAllGoods(c fiber.Ctx) error {
	return c.JSON(state)
}

func (h *Handler) CreateGood(c fiber.Ctx) error {
	var good model.Good
	err := jsoniter.Unmarshal(c.Body(), &good)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(model.Error{
			Status:  fiber.StatusUnprocessableEntity,
			Code:    "marshal_error",
			Message: err.Error(),
		})
	}

	state = append(state, &good)

	return c.JSON(good)
}

func (h *Handler) UpdateGood(c fiber.Ctx) error {
	var good model.Good
	err := jsoniter.Unmarshal(c.Body(), &good)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(model.Error{
			Status:  fiber.StatusUnprocessableEntity,
			Code:    "marshal_error",
			Message: err.Error(),
		})
	}

	var sendBody *model.Good
	for _, item := range state {
		if item.ID == good.ID {
			item.Name = good.Name
			item.Description = good.Description
			item.Count = good.Count
			item.Weight = good.Weight
			sendBody = item
		}
	}
	if sendBody == nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Error{
			Status:  fiber.StatusNotFound,
			Code:    "there_is_no_good",
			Message: "the good was not found",
		})
	}

	return c.JSON(sendBody)
}

func (h *Handler) DeleteGood(c fiber.Ctx) error {
	var good model.Good
	err := jsoniter.Unmarshal(c.Body(), &good)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(model.Error{
			Status:  fiber.StatusUnprocessableEntity,
			Code:    "marshal_error",
			Message: err.Error(),
		})
	}

	var idx int
	var item *model.Good
	for i := range state {
		if state[i].ID == good.ID {
			idx = i
			item = state[i]
		}
	}
	if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Error{
			Status:  fiber.StatusNotFound,
			Code:    "there_is_no_good",
			Message: "the good was not found",
		})
	}

	state = append(state[:idx], state[idx+1:]...)

	return c.JSON(good)
}
