package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/ceit-aut/policeman/internal/model"
	"github.com/ceit-aut/policeman/internal/port/http/request"
	"github.com/ceit-aut/policeman/internal/port/http/response"
	"github.com/ceit-aut/policeman/pkg/crypto"
	"github.com/gofiber/fiber/v2"
)

var (
	errParsingFailed     = errors.New("cannot parse request body")
	errUserAlreadyExists = errors.New("username is already exists")
	errHashPassword      = errors.New("failed to hash your password")
	errSaveUser          = errors.New("failed to save user")
	errNotFound          = errors.New("username not found")
	errWrongPassword     = errors.New("password is incorrect")
	errTokenGeneration   = errors.New("failed to generate token")
)

// Register a user inside application.
// takes a user info request.
func (h *Handler) Register(ctx *fiber.Ctx) error {
	// create a user request
	var userReq request.UserInfo

	// parse user request body
	if err := ctx.BodyParser(&userReq); err != nil {
		log.Println(err)

		return errParsingFailed
	}

	// check username exists
	if h.Repositories.Users.Exists(userReq.Username) {
		return errUserAlreadyExists
	}

	// hash user password
	hp, er := crypto.HashData(userReq.Password)
	if er != nil {
		log.Println(er)

		return errHashPassword
	}

	// create a user model
	u := model.User{
		Username: userReq.Username,
		Password: hp,
	}

	// save user
	if err := h.Repositories.Users.Insert(u); err != nil {
		log.Println(err)

		return errSaveUser
	}

	return ctx.SendStatus(http.StatusOK)
}

// Login a user by return a token.
// takes a user info request.
func (h *Handler) Login(ctx *fiber.Ctx) error {
	// create a user request
	var userReq request.UserInfo

	// parse user request body
	if err := ctx.BodyParser(&userReq); err != nil {
		log.Println(err)

		return errParsingFailed
	}

	// get user from database
	user := h.Repositories.Users.GetSingle(userReq.Username)
	if user == nil {
		return errNotFound
	}

	// check password
	if !crypto.IsEqual(userReq.Password, user.Password) {
		return errWrongPassword
	}

	// generating jwt token
	token, expireTime, err := h.JWT.GenerateJWT(user.Username, user.Password)
	if err != nil {
		log.Println(err)

		return errTokenGeneration
	}

	// creating a token response
	resp := response.TokenResponse{
		Token:   token,
		Expires: expireTime,
	}

	return ctx.JSON(resp)
}
