package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ceit-aut/policeman/internal/model"
	"github.com/ceit-aut/policeman/internal/port/http/request"
	"github.com/ceit-aut/policeman/internal/port/http/response"
	"github.com/gofiber/fiber/v2"
)

var (
	errEmptyAddress = errors.New("address cannot be empty")
	errSaveEndpoint = errors.New("failed to save endpoint")
)

// RegisterEndpoint will add an endpoint to application.
// takes a endpoint info request.
func (h *Handler) RegisterEndpoint(ctx *fiber.Ctx) error {
	// create a user request
	var userReq request.EndpointInfo

	// parse user request
	if err := ctx.BodyParser(&userReq); err != nil {
		log.Println(err)

		return errParsingFailed
	}

	// check address
	if userReq.Address == "" {
		return errEmptyAddress
	}

	// creating a new endpoint model
	e := model.Endpoint{
		Username:    ctx.Locals("username").(string),
		Url:         userReq.Address,
		Threshold:   h.Threshold,
		FailedTimes: 0,
		CreateTime:  time.Now(),
	}

	// save endpoint into database
	if err := h.Repositories.Endpoints.Upsert(e); err != nil {
		log.Println(err)

		return errSaveEndpoint
	}

	return ctx.SendStatus(http.StatusOK)
}

// GetAllEndpoints for a user.
func (h *Handler) GetAllEndpoints(ctx *fiber.Ctx) error {
	// create endpoints list
	var endpoints []response.EndpointResponse

	// get all endpoints of a user
	list := h.Repositories.Endpoints.GetSingle(ctx.Locals("username").(string))

	// create responses
	for _, item := range list {
		er := response.EndpointResponse{
			Address:   item.Url,
			CreatedAt: item.CreateTime,
		}

		endpoints = append(endpoints, er)
	}

	return ctx.JSON(endpoints)
}

// GetEndpointStatus will return one endpoint status.
func (h *Handler) GetEndpointStatus(ctx *fiber.Ctx) error {
	// create list of requests
	var requests []response.EndpointRequest

	// get one endpoint requests
	list := h.Repositories.Requests.GetAll(ctx.Params("address"))

	// generate requests response
	for _, item := range list {
		er := response.EndpointRequest{
			Status: item.Code,
			Time:   item.CreateTime,
		}

		requests = append(requests, er)
	}

	return ctx.JSON(requests)
}

// GetEndpointWarnings will return all the warnings for an endpoint.
func (h *Handler) GetEndpointWarnings(ctx *fiber.Ctx) error {
	return nil
}
