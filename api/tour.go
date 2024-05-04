package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateTourRequestParams struct {
	Name            string                    `json:"name" binding:"required"`
	Duration        int64                     `json:"duration" binding:"required,min=1"`
	RatingsAverage  int64                     `json:"ratings_average" binding:"omitempty,min=0,max=5"`
	MaxGroupSize    int64                     `json:"max_group_size" binding:"required,min=1"`
	Difficulty      string                    `json:"difficulty" binding:"required,oneof=low medieum hard very_hard"`
	RatingsQuantity int64                     `json:"ratings_quantity" binding:"omitempty,min=0"`
	Price           int64                     `json:"price" binding:"required,min=0"`
	Summary         string                    `json:"summary" binding:"required,min=10"`
	Description     string                    `json:"description" binding:"required,min=20"`
	ImageCover      string                    `json:"imagecover" binding:"required,url"`
	Images          string                    `json:"images" binding:"omitempty,url"`
	StartDates      time.Time                 `json:"start-dates" binding:"required"`
	SecretTour      bool                      `json:"secret_tour"`
	StartLocation   CreateStartLocationParams `json:"start_location" binding:"required"`
	Location        CreateLocationParams      `json:"location" binding:"required"`
}

type CreateStartLocationParams struct {
	Lat         float64 `json:"lat" binding:"required"`
	Long        float64 `json:"long" binding:"required"`
	Address     string  `json:"address" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Type        string  `json:"type"  binding:"required,oneof=point square circle"`
}

type CreateLocationParams struct {
	Lat         float64 `json:"lat" binding:"required"`
	Long        float64 `json:"long" binding:"required"`
	Address     string  `json:"address" binding:"required"`
	Description string  `json:"description"`
	Day         int64   `json:"day" binding:"required"`
	Type        string  `json:"type"  binding:"required,oneof=point square circle"`
}

func (server *Server) createNewTourController(ctx *gin.Context) {

	var req CreateTourRequestParams
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location := db.CreateLocationParams{
		Lat:         req.Location.Lat,
		Long:        req.Location.Long,
		Address:     sql.NullString{String: req.Location.Address, Valid: true},
		Description: sql.NullString{String: req.Location.Description, Valid: true},
		Day:         req.Location.Day,
		Type:        db.LocationType(req.Location.Type),
	}

	startLocation := db.CreateStartLocationParams{
		Lat:         req.StartLocation.Lat,
		Long:        req.StartLocation.Long,
		Description: sql.NullString{String: req.StartLocation.Description, Valid: true},
		Type:        db.LocationType(req.StartLocation.Type),
		Address:     req.StartLocation.Address,
	}
	tour := db.CreateTourParams{
		Name:            req.Name,
		Duration:        req.Duration,
		RatingsAverage:  req.RatingsAverage,
		MaxGroupSize:    req.MaxGroupSize,
		Difficulty:      req.Difficulty,
		RatingsQuantity: req.RatingsQuantity,
		Price:           req.Price,
		Summary:         req.Summary,
		Description:     req.Description,
		Imagecover:      req.ImageCover,
		Images:          req.Images,
		StartDates:      req.StartDates,
		SecretTour:      sql.NullBool{Bool: false},
	}

	createdTour, err := server.store.CreateNewTourTransaction(ctx, db.LocationTrParams(location), db.StartlocationTrParams(startLocation), db.TourTrParams(tour))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// If there's no error, proceed with your business logic
	ctx.JSON(http.StatusOK, gin.H{"message": "Tour created successfully", "data": createdTour})
}
