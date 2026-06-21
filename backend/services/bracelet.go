package services

import (
	"jade-grading/database"
	"jade-grading/models"

	"gorm.io/gorm"
)

type BraceletService struct {
	gradingService *GradingService
	db             *gorm.DB
}

func NewBraceletService() *BraceletService {
	return &BraceletService{
		gradingService: NewGradingService(),
		db:             database.GetDB(),
	}
}

func (s *BraceletService) CreateBracelet(req *models.CreateBraceletRequest) (*models.JadeBracelet, error) {
	result := s.gradingService.CalculateGrade(req.Translucency, req.Fineness, req.BeadCount)

	bracelet := &models.JadeBracelet{
		Name:         req.Name,
		Material:     req.Material,
		Translucency: req.Translucency,
		Fineness:     req.Fineness,
		BeadCount:    req.BeadCount,
		Score:        result.Score,
		Grade:        result.Grade,
	}

	if err := s.db.Create(bracelet).Error; err != nil {
		return nil, err
	}

	return bracelet, nil
}

func (s *BraceletService) GetAllBracelets() ([]models.JadeBracelet, error) {
	var bracelets []models.JadeBracelet
	if err := s.db.Order("created_at DESC").Find(&bracelets).Error; err != nil {
		return nil, err
	}
	return bracelets, nil
}

func (s *BraceletService) GetBraceletByID(id uint64) (*models.JadeBracelet, error) {
	var bracelet models.JadeBracelet
	if err := s.db.First(&bracelet, id).Error; err != nil {
		return nil, err
	}
	return &bracelet, nil
}

func (s *BraceletService) DeleteBracelet(id uint64) error {
	if err := s.db.Delete(&models.JadeBracelet{}, id).Error; err != nil {
		return err
	}
	return nil
}
