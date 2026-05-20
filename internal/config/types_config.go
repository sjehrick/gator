// Package config provides config structs and functions for the main package
package config

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}
