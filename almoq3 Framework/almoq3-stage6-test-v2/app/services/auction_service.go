package services

import (
	"gorm.io/gorm"
)

type AuctionServiceService struct {
	db *gorm.DB
}

func NewAuctionServiceService(db *gorm.DB) *AuctionServiceService {
	return &AuctionServiceService{
		db: db,
	}
}

// Add your business logic methods here
// Example:
// func (s *AuctionServiceService) FindAll() ([]models.AuctionService, error) {
// 	var items []models.AuctionService
// 	err := s.db.Find(&items).Error
// 	return items, err
// }
