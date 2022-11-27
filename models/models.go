package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"text;not null;default:null"`
	Name  string `json:"name" gorm:"text;not null;default:null"`
	Teams []Team
}

type Team struct {
	gorm.Model
	TeamName  string `json:"team_name" gorm:"text;not null;default:null"`
	Questions []Question
}

type Question struct {
	gorm.Model
	QuestionString string `json:"question_string" grom:"text;not null;default:null"`
}
