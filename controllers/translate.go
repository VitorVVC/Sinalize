package controllers

import (
	"SinalizeB/models"
	"SinalizeB/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type TranslateController struct {
	validator *validator.Validate
	db        *gorm.DB
}

func NewTranslateController() *TranslateController {
	// TODO -> EnvString
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.Translate{})

	return &TranslateController{
		validator: validator.New(),
		db:        db,
	}
}

func (t *TranslateController) Create(ctx echo.Context) error {
	var data models.Translate

	if err := ctx.Bind(&data); err != nil {
		utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse body")
	}

	if err := t.validator.Struct(&data); err != nil {
		utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to validate body")
	}

	if err := t.db.Create(&data).Error; err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to save data")
	}

	return utils.HTTPCreated(ctx, data)
}
