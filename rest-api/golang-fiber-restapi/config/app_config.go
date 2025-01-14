package config

import "os"


func ListenAndServe() string {
	LoadDotEnv()
	return os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")
}

func ApiSecret() string {
	LoadDotEnv()
	return os.Getenv("API_SECRET")
}