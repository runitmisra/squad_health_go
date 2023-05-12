package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"text;not null;default:null"`
	Name  string `json:"name" gorm:"text;not null;default:null"`
	Teams []Team `json:"teams" gorm:"many2many:user_teams;"`
}

type Team struct {
	gorm.Model
	TeamName string `json:"team_name" gorm:"text;not null;default:null"`
}

type Feedback struct {
	gorm.Model
	QuestionString string `json:"question_string" grom:"text;default:null"`
	Response       string `json:"response_string" grom:"text;default:null"`
	UserID         uint
	TeamID         uint
	User           User
	Team           Team
}
