package config

import (
	"endevix-cli-go/internal/config/parser"
	"time"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

var parsersMap = map[Format]ParserFunc{
	FormatJSON: parser.JsonToMap,
}

type ParserFunc func(string) (map[string]interface{}, error)

type Client struct {
	CompanyName  string `json:"company_name" validate:"required,min=2,max=100"`
	MainContact  string `json:"main_contact" validate:"required,min=2,max=100"`
	ContactEmail string `json:"contact_email" validate:"required,email"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=7,max=20"`
}

type Project struct {
	ProjectID       string    `json:"project_id" validate:"required,min=3,max=50"`
	ProjectName     string    `json:"project_name" validate:"required,min=3,max=100"`
	Status          string    `json:"status" validate:"required,oneof=pending in_progress completed cancelled"`
	Priority        string    `json:"priority" validate:"required,oneof=low medium high critical"`
	Budget          float64   `json:"budget" validate:"gte=0"`
	Currency        string    `json:"currency" validate:"required,len=3"`
	IsActive        bool      `json:"is_active"`
	StartDate       time.Time `json:"start_date" validate:"required"`
	Client          Client    `json:"client" validate:"required"`
	Tags            []string  `json:"tags" validate:"omitempty,dive,min=2,max=30"`
	AdditionalNotes string    `json:"additional_notes" validate:"max=500"`
}
