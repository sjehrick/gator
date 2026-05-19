// Package config provides config structs and functions for the main package
package config

type Config struct {
	dbURL           string `json:"db_url"`
	currentUserName string `json:"current_user_name"`
}
