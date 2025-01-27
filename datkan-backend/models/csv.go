package models

import (
	"time"

	"github.com/google/uuid"
)

type RawCSVData struct {
	ID 						uuid.UUID `gorm:"type:uuid;primaryKey"`
	OriginalData 	string 		`gorm:"type:jsonb"`
	CreatedAt 		time.Time
	ProcessedAt 	*time.Time
}

type AnonCustomer struct {
	ID						uuid.UUID `gorm:"type:uuid;primaryKey`
	RFM_Recency 	int				`gorm:"index"`
	RFM_Frequency int 			`gorm:"index"`
	RFM_Monetary 	int				`gorm:"index"`
	FirstPurchase time.Time 
	LastPurchase  time.Time
	CustomerHash  string 		`gorm:"uniqueIndex,size:64"`
}