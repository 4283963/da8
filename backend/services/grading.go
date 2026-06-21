package services

import (
	"fmt"
	"math"
)

const (
	TranslucencyWeight = 0.50
	FinenessWeight     = 0.50
	StandardBeadCount  = 17
	BeadCountTolerance = 2
)

type GradeResult struct {
	Score float64 `json:"score"`
	Grade string  `json:"grade"`
}

type GradingService struct{}

func NewGradingService() *GradingService {
	return &GradingService{}
}

func (s *GradingService) CalculateGrade(translucency, fineness float64, beadCount int) *GradeResult {
	baseScore := translucency*TranslucencyWeight + fineness*FinenessWeight

	beadPenalty := 0.0
	if math.Abs(float64(beadCount-StandardBeadCount)) > BeadCountTolerance {
		excess := math.Abs(float64(beadCount-StandardBeadCount)) - BeadCountTolerance
		beadPenalty = excess * 1.0
	}

	finalScore := baseScore - beadPenalty
	if finalScore < 0 {
		finalScore = 0
	}
	if finalScore > 100 {
		finalScore = 100
	}

	grade := scoreToGrade(finalScore)

	return &GradeResult{
		Score: math.Round(finalScore*100) / 100,
		Grade: grade,
	}
}

func scoreToGrade(score float64) string {
	switch {
	case score >= 90:
		return "特级"
	case score >= 80:
		return "一级"
	case score >= 70:
		return "二级"
	case score >= 60:
		return "三级"
	default:
		return "等外"
	}
}

func (s *GradingService) ExplainGrade(score float64, grade string) string {
	return fmt.Sprintf("综合评分 %.2f 分，评定等级：%s", score, grade)
}
