package config

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/naoina/kocha"
	"github.com/naoina/kocha/log"
)

var (
	AppName   = "helloworld"
	AppConfig = &kocha.Config{
		Addr:          kocha.Getenv("KOCHA_ADDR", "127.0.0.1:9100"),
		AppPath:       rootPath,
		AppName:       AppName,
		DefaultLayout: "app",
		Template: &kocha.Template{
			PathInfo: kocha.TemplatePathInfo{
				Name: AppName,
				Paths: []string{
					filepath.Join(rootPath, "app", "view"),
				},
			},
			FuncMap: kocha.TemplateFuncMap{},
		},

		// Logger settings.
		Logger: &kocha.LoggerConfig{
			Writer:    os.Stdout,
			Formatter: &log.LTSVFormatter{},
			Level:     log.INFO,
		},

		// Middlewares.
		Middlewares: []kocha.Middleware{
			&kocha.RequestLoggingMiddleware{},
			&kocha.PanicRecoverMiddleware{},
			&kocha.FormMiddleware{},
			&kocha.SessionMiddleware{
				Name: "helloworld_session",
				Store: &kocha.SessionCookieStore{
					// AUTO-GENERATED Random keys. DO NOT EDIT.
					SecretKey:  "\u007f\xc5ƯZ\xa8\x91\x04FI\x92\xd5\xf7<\x8f\x91u\x99?\xa6+K\xc0\xca\x01\xe5\xf1\x9e\x00E\xe8;",
					SigningKey: "\xd9]~d\x03\xeav\xc1Y\x8b\xf7\x8b\xbe\xd4\f7",
				},

				// Expiration of session cookie, in seconds, from now.
				// Persistent if -1, For not specify, set 0.
				CookieExpires: time.Duration(90) * time.Hour * 24,

				// Expiration of session data, in seconds, from now.
				// Perssitent if -1, For not specify, set 0.
				SessionExpires: time.Duration(90) * time.Hour * 24,
				HttpOnly:       false,
			},
			&kocha.FlashMiddleware{},
			&kocha.DispatchMiddleware{},
		},

		MaxClientBodySize: 1024 * 1024 * 10, // 10MB
	}

	_, configFileName, _, _ = runtime.Caller(0)
	rootPath                = filepath.Dir(filepath.Join(configFileName, ".."))
)
