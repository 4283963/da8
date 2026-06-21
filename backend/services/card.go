package services

import (
	"fmt"
	"jade-grading/models"
	"strings"
	"time"
)

type CardService struct{}

func NewCardService() *CardService {
	return &CardService{}
}

func (s *CardService) GeneratePrintCard(b *models.JadeBracelet) (string, string) {
	cardWidth := 42

	beadCountText := "未填写"
	if b.BeadCount != nil {
		beadCountText = fmt.Sprintf("%d 颗", *b.BeadCount)
	}

	createdDate := b.CreatedAt.Format("2006-01-02 15:04")

	var sb strings.Builder

	sb.WriteString(s.drawBorderTop(cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawCenterLine("和 田 玉 手 串 质 地 评 定 卡", cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawSeparator(cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawTwoColumn("编号", b.Name, cardWidth))
	sb.WriteString(s.drawTwoColumn("材质", b.Material, cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawSeparator(cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawTwoColumn("透光度", fmt.Sprintf("%.2f 分", b.Translucency), cardWidth))
	sb.WriteString(s.drawTwoColumn("细度", fmt.Sprintf("%.2f 分", b.Fineness), cardWidth))
	sb.WriteString(s.drawTwoColumn("珠子颗数", beadCountText, cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawSeparator(cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawTwoColumn("综合评分", fmt.Sprintf("★ %.2f ★", b.Score), cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawCenterLine(fmt.Sprintf("【 评 定 等 级：%s 】", s.decorateGrade(b.Grade)), cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawSeparator(cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawRightLine("评定日期："+createdDate, cardWidth))
	sb.WriteString(s.drawRightLine("评定系统：和田玉智能评级 v1.0", cardWidth))
	sb.WriteString(s.drawLine("", cardWidth))
	sb.WriteString(s.drawBorderBottom(cardWidth))

	fileName := fmt.Sprintf("%s_%s_%s.txt",
		s.cleanFileName(b.Material),
		s.cleanFileName(b.Name),
		time.Now().Format("200601021504"),
	)

	return sb.String(), fileName
}

func (s *CardService) decorateGrade(grade string) string {
	switch grade {
	case "特级":
		return "☆☆☆ 特级 ☆☆☆"
	case "一级":
		return "★★ 一级 ★★"
	case "二级":
		return "★ 二级 ★"
	default:
		return grade
	}
}

func (s *CardService) cleanFileName(name string) string {
	var sb strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') ||
			(r >= 0x4e00 && r <= 0x9fff) {
			sb.WriteRune(r)
		}
	}
	if sb.Len() == 0 {
		return "card"
	}
	return sb.String()
}

func (s *CardService) drawBorderTop(width int) string {
	return "╔" + strings.Repeat("═", width-2) + "╗\n"
}

func (s *CardService) drawBorderBottom(width int) string {
	return "╚" + strings.Repeat("═", width-2) + "╝\n"
}

func (s *CardService) drawSeparator(width int) string {
	return "╠" + strings.Repeat("─", width-2) + "╣\n"
}

func (s *CardService) drawLine(text string, width int) string {
	return s.drawContentLine(text, width, false)
}

func (s *CardService) drawCenterLine(text string, width int) string {
	return s.drawContentLine(text, width, true)
}

func (s *CardService) drawContentLine(text string, width int, center bool) string {
	prefix := "║"
	suffix := "║"
	innerWidth := width - 2

	displayLen := s.visualLen(text)

	if displayLen > innerWidth {
		runes := []rune(text)
		for displayLen > innerWidth-3 && len(runes) > 0 {
			runes = runes[:len(runes)-1]
			displayLen = s.visualLen(string(runes))
		}
		text = string(runes) + "..."
		displayLen = s.visualLen(text)
	}

	var padding string
	if center {
		totalPad := innerWidth - displayLen
		leftPad := totalPad / 2
		rightPad := totalPad - leftPad
		padding = strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
	} else {
		padding = text + strings.Repeat(" ", innerWidth-displayLen)
	}

	return prefix + padding + suffix + "\n"
}

func (s *CardService) drawTwoColumn(label, value string, width int) string {
	innerWidth := width - 2
	leftColWidth := 14

	labelText := fmt.Sprintf(" %s：", label)
	labelDisplayLen := s.visualLen(labelText)
	labelPad := strings.Repeat(" ", leftColWidth-labelDisplayLen)

	leftPart := labelText + labelPad
	leftDisplayLen := s.visualLen(leftPart)

	valueDisplayLen := s.visualLen(value)
	rightPadLen := innerWidth - leftDisplayLen - valueDisplayLen
	if rightPadLen < 0 {
		rightPadLen = 0
	}

	content := leftPart + value + strings.Repeat(" ", rightPadLen)
	return "║" + content + "║\n"
}

func (s *CardService) drawRightLine(text string, width int) string {
	innerWidth := width - 2
	displayLen := s.visualLen(text)

	if displayLen >= innerWidth {
		return s.drawLine(text, width)
	}

	padLen := innerWidth - displayLen
	content := strings.Repeat(" ", padLen) + text
	return "║" + content + "║\n"
}

func (s *CardService) visualLen(s2 string) int {
	length := 0
	for _, r := range s2 {
		if r >= 0x4e00 && r <= 0x9fff {
			length += 2
		} else {
			length += 1
		}
	}
	return length
}
