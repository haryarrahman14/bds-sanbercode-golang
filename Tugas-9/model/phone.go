package model

import (
	"fmt"
	"strings"
)

type Phone struct{
	Name, Brand	string
	Year		int
	Colors		[]string
 }

type DisplayPhoneInfo interface {
	GetInfo()	string
}

func (p Phone) GetInfo() string {
	colorStr := strings.Join(p.Colors, ", ")
	info := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.Name, p.Brand, p.Year, colorStr)
	return info
}

