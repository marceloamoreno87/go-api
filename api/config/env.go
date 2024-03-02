package config

import (
	"os"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type EnvironmentInterface interface {
	GetNameProject() string
	GetHost() string
	GetPort() string
	GetDBHost() string
	GetDBPort() string
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
	GetJWTSecretKey() string
	GetJWTExpiresIn() string
	GetMailHost() string
	GetMailPort() string
	GetMailUser() string
	GetMailPassword() string
	GetMailDriver() string
	GetMailerSendApiKey() string
	GetSendgridApiKey() string
}

type Env struct {
	nameProject      string
	port             string
	dbDriver         string
	dbSslMode        string
	dbHost           string
	dbPort           string
	dbUser           string
	dbPassword       string
	dbName           string
	jwtSecretKey     string
	jwtExpiresIn     string
	mailHost         string
	mailPort         string
	mailUser         string
	mailPassword     string
	mailDriver       string
	mailerSendApiKey string
	sendgridApiKey   string
}

var Environment *Env

func NewEnv() {
	newEnv := &Env{
		nameProject:      os.Getenv("NAME_PROJECT"),
		dbDriver:         os.Getenv("DB_DRIVER"),
		dbSslMode:        os.Getenv("DB_SSL_MODE"),
		port:             os.Getenv("PORT"),
		dbHost:           os.Getenv("DB_HOST"),
		dbPort:           os.Getenv("DB_PORT"),
		dbUser:           os.Getenv("DB_USER"),
		dbPassword:       os.Getenv("DB_PASSWORD"),
		dbName:           os.Getenv("DB_NAME"),
		jwtSecretKey:     os.Getenv("JWT_SECRET_KEY"),
		jwtExpiresIn:     os.Getenv("JWT_EXPIRES_IN"),
		mailHost:         os.Getenv("MAIL_HOST"),
		mailPort:         os.Getenv("MAIL_PORT"),
		mailUser:         os.Getenv("MAIL_USER"),
		mailPassword:     os.Getenv("MAIL_PASSWORD"),
		mailDriver:       os.Getenv("MAIL_DRIVER"),
		mailerSendApiKey: os.Getenv("MAILERSEND_API_KEY"),
		sendgridApiKey:   os.Getenv("SENDGRID_API_KEY"),
	}
	notify := newEnv.Validate()
	if notify.HasErrors() {
		panic(notify.Messages())
	}

	Environment = newEnv
}

func (e *Env) Validate() (notify *notification.Errors) {
	notify = notification.New()

	if e.nameProject == "" {
		notify.AddError("NameProject is required", "config.env.nameProject")
	}
	if e.port == "" {
		notify.AddError("Port is required", "config.env.port")
	}
	if e.dbDriver == "" {
		notify.AddError("DBDriver is required", "config.env.dbDriver")
	}
	if e.dbSslMode == "" {
		notify.AddError("DBSslMode is required", "config.env.dbSslMode")
	}
	if e.dbHost == "" {
		notify.AddError("DBHost is required", "config.env.dbHost")
	}
	if e.dbPort == "" {
		notify.AddError("DBPort is required", "config.env.dbPort")
	}
	if e.dbUser == "" {
		notify.AddError("DBUser is required", "config.env.dbUser")
	}
	if e.dbPassword == "" {
		notify.AddError("DBPassword is required", "config.env.dbPassword")
	}
	if e.dbName == "" {
		notify.AddError("DBName is required", "config.env.dbName")
	}
	if e.jwtSecretKey == "" {
		notify.AddError("JWTSecretKey is required", "config.env.jwtSecretKey")
	}
	if e.jwtExpiresIn == "" {
		notify.AddError("JWTExpiresIn is required", "config.env.jwtExpiresIn")
	}
	if e.mailHost == "" {
		notify.AddError("MailhogHost is required", "config.env.mailHost")
	}
	if e.mailPort == "" {
		notify.AddError("MailhogPort is required", "config.env.mailPort")
	}
	if e.mailUser == "" {
		notify.AddError("MailhogUser is required", "config.env.mailUser")
	}
	if e.mailPassword == "" {
		notify.AddError("MailhogPassword is required", "config.env.mailPassword")
	}
	if e.mailDriver == "" {
		notify.AddError("MailDriver is required", "config.env.mailDriver")
	}
	if e.mailerSendApiKey == "" {
		notify.AddError("MailerSendApiKey is required", "config.env.mailerSendApiKey")
	}
	if e.sendgridApiKey == "" {
		notify.AddError("SendgridApiKey is required", "config.env.sendgridApiKey")
	}
	return
}

func (e *Env) GetNameProject() string {
	return e.nameProject
}

func (e *Env) GetPort() string {
	return e.port
}

func (e *Env) GetDBHost() string {
	return e.dbHost
}

func (e *Env) GetDBDriver() string {
	return e.dbDriver
}

func (e *Env) GetDBSslMode() string {
	return e.dbSslMode
}

func (e *Env) GetDBPort() string {
	return e.dbPort
}

func (e *Env) GetDBUser() string {
	return e.dbUser
}

func (e *Env) GetDBPassword() string {
	return e.dbPassword
}

func (e *Env) GetDBName() string {
	return e.dbName
}

func (e *Env) GetJWTSecretKey() string {
	return e.jwtSecretKey
}

func (e *Env) GetJWTExpiresIn() string {
	return e.jwtExpiresIn
}

func (e *Env) GetMailHost() string {
	return e.mailHost
}

func (e *Env) GetMailPort() string {
	return e.mailPort
}

func (e *Env) GetMailUser() string {
	return e.mailUser
}

func (e *Env) GetMailPassword() string {
	return e.mailPassword
}

func (e *Env) GetMailDriver() string {
	return e.mailDriver
}

func (e *Env) GetMailerSendApiKey() string {
	return e.mailerSendApiKey
}

func (e *Env) GetSendgridApiKey() string {
	return e.sendgridApiKey
}
