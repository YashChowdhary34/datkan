package handlers

import (
	"time"

	"github.com/datkan-ai/core/utils"
	"github.com/gin-gonic/gin"
)

// 400 - badRequest
// 500 - internalServiceError

const (
	maxCSVSize = 10 * 1024 * 1024
)

type CSVRow struct {
	Email             string    `csv:"email"`
	FirstPurchaseDate time.Time `csv:"first_purchase_date"`
	LastPurchaseDate  time.Time	`csv:"last_purchase_date"`
	TotalSpent				float64		`csv:"total_spent"`
}

func HandleCSVUpload(c *gin.Context) {
	file, err := c.FormFile("csv")
	if err != nil {
		c.JSON(400, gin.H{"error": "No CSV file uploaded"})
		return
	}

	if file.Size > maxCSVSize {
		c.JSON(400, gin.H{"error": "File exceeds 10MB limit"})
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "File processing failed"})
	}
	defer f.Close()

	if err := utils.ValidateCSVHeaders(f); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// reset reader after header validation
	f.Seek(0, 0) 
}