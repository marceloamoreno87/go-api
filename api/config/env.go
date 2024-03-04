package config

import (
	"os"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

var (
	Envs        = []string{"development", "production"}
	DbDrivers   = []string{"postgres", "mysql"}
	MailDrivers = []string{"smtp", "mailersend", "sendgrid"}
)

type EnvironmentInterface interface {
	GetNameProject() string
	GetPort() string
	GetDBHost() string
	GetDBPort() string
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
	GetDBDriver() string
	GetDBSslMode() string
	GetMailFrom() string
	GetJWTSecretKey() string
	GetJWTExpiresIn() string
	GetMailHost() string
	GetMailPort() string
	GetMailUser() string
	GetMailPassword() string
	GetMailDriver() string
	GetMailerSendApiKey() string
	GetSendgridApiKey() string
	GetFrontendUrl() string
	GetEnv() string
}

type Env struct {
	nameProject      string
	port             string
	env              string
	frontendUrl      string
	dbDriver         string
	dbSslMode        string
	dbHost           string
	dbPort           string
	dbUser           string
	dbPassword       string
	dbName           string
	jwtSecretKey     string
	jwtExpiresIn     string
	mailFrom         string
	mailHost         string
	mailPort         string
	mailUser         string
	mailPassword     string
	mailDriver       string
	mailerSendApiKey string
	sendgridApiKey   string
}

var Environment EnvironmentInterface

func NewEnv() {
	newEnv := &Env{
		nameProject:      os.Getenv("NAME_PROJECT"),
		env:              os.Getenv("ENV"),
		frontendUrl:      os.Getenv("FRONTEND_URL"),
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
		mailFrom:         os.Getenv("MAIL_FROM"),
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
	notify.CheckRequiredField(e.GetNameProject(), "NameProject", "config.env.nameProject")
	notify.CheckRequiredField(e.GetPort(), "Port", "config.env.port")
	notify.CheckRequiredField(e.GetDBHost(), "DBHost", "config.env.dbHost")
	notify.CheckRequiredField(e.GetDBDriver(), "DBDriver", "config.env.dbDriver")
	notify.CheckRequiredField(e.GetDBSslMode(), "DBSslMode", "config.env.dbSslMode")
	notify.CheckRequiredField(e.GetDBPort(), "DBPort", "config.env.dbPort")
	notify.CheckRequiredField(e.GetDBUser(), "DBUser", "config.env.dbUser")
	notify.CheckRequiredField(e.GetDBPassword(), "DBPassword", "config.env.dbPassword")
	notify.CheckRequiredField(e.GetDBName(), "DBName", "config.env.dbName")
	notify.CheckRequiredField(e.GetJWTSecretKey(), "JWTSecretKey", "config.env.jwtSecretKey")
	notify.CheckRequiredField(e.GetJWTExpiresIn(), "JWTExpiresIn", "config.env.jwtExpiresIn")
	notify.CheckRequiredField(e.GetMailFrom(), "MailFrom", "config.env.mailFrom")
	notify.CheckRequiredField(e.GetMailHost(), "MailHost", "config.env.mailHost")
	notify.CheckRequiredField(e.GetMailPort(), "MailPort", "config.env.mailPort")
	notify.CheckRequiredField(e.GetMailUser(), "MailUser", "config.env.mailUser")
	notify.CheckRequiredField(e.GetMailPassword(), "MailPassword", "config.env.mailPassword")
	notify.CheckRequiredField(e.GetMailDriver(), "MailDriver", "config.env.mailDriver")
	notify.CheckRequiredField(e.GetMailerSendApiKey(), "MailerSendApiKey", "config.env.mailerSendApiKey")
	notify.CheckRequiredField(e.GetSendgridApiKey(), "SendgridApiKey", "config.env.sendgridApiKey")
	notify.CheckRequiredField(e.GetFrontendUrl(), "FrontendUrl", "config.env.frontendUrl")
	notify.CheckRequiredField(e.GetEnv(), "Env", "config.env.env")
	notify.CheckIsContains(e.GetEnv(), Envs, "Env", "config.env.env")
	notify.CheckIsContains(e.GetDBDriver(), DbDrivers, "DBDriver", "config.env.dbDriver")
	notify.CheckIsContains(e.GetMailDriver(), MailDrivers, "MailDriver", "config.env.mailDriver")
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

func (e *Env) GetMailFrom() string {
	return e.mailFrom
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

func (e *Env) GetFrontendUrl() string {
	return e.frontendUrl
}

func (e *Env) GetEnv() string {
	return e.env
}
