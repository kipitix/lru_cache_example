package handlers

import (
	"log"
	"lrucache/application/services"
	"lrucache/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_cacheTTL = time.Second * 5
)

type UserHandlers interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type userHandlersImpl struct {
	services services.Services
}

func NewUserHandlers(services services.Services) UserHandlers {
	return &userHandlersImpl{
		services: services,
	}
}

var _ UserHandlers = (*userHandlersImpl)(nil)

func (us *userHandlersImpl) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		log.Printf("failed to parse request body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail to parse request"})
		return
	}

	if err := us.services.UserRepository().InsertOrUpdateUser(ctx, user); err != nil {
		log.Printf("failed to insert user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to create user"})
	}
}

func (us *userHandlersImpl) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	email := c.Param("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty email parameter"})
		return
	}

	var cachedRecord domain.User

	if err := us.services.Storage().Get(ctx, email, &cachedRecord); err == nil {
		c.JSON(http.StatusOK, gin.H{"payload": cachedRecord})
		return
	} else {
		log.Printf("error on get data from cache: %s", err)
	}

	record, err := us.services.UserRepository().FindUserByEmail(ctx, email)
	if err != nil {
		log.Printf("failed to FindUsers: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to find users"})
		return
	}

	// Put or update cache
	if err := us.services.Storage().Set(ctx, email, record, _cacheTTL); err != nil {
		log.Printf("error on set data into cache: %s", err)
	}

	// Send response
	c.JSON(http.StatusOK, gin.H{"payload": record})
}
