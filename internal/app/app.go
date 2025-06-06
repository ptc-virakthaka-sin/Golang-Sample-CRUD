package app

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"learn-fiber/api/auth"
	"learn-fiber/api/department"
	"learn-fiber/api/student"
	"learn-fiber/config"
	"learn-fiber/internal/ierror"
	"learn-fiber/internal/middleware"
	"learn-fiber/internal/util/gracefulshutdown"
	"learn-fiber/internal/util/validator"
	"learn-fiber/pkg/database"
	"learn-fiber/pkg/logger"
	"learn-fiber/pkg/redis"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
	logger.Init(config.Cfg.Log.Level)
	db, err := database.New()
	if err != nil {
		logger.L.Fatal("database connection error: %v", err)
	}
	if err = redis.New(); err != nil {
		logger.L.Fatal("redis connection error: %v", err)
	}
	go redis.StartWorker(db, "worker-1")
	validator.Init()
	app := fiber.New(fiber.Config{
		IdleTimeout:       time.Duration(config.Cfg.IdleTimeout) * time.Second,
		EnablePrintRoutes: config.Cfg.EnablePrintRoute,
		ErrorHandler:      ierror.HandleErrorResponse(),
		JSONDecoder:       sonic.Unmarshal,
		JSONEncoder:       sonic.Marshal,
	})

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(favicon.New())
	app.Use(middleware.Limiter())
	app.Use(requestid.New(requestid.Config{Generator: uuid.NewString}))

	initRootHandler(app)
	initHealthCheck(app, db)
	api := app.Group("/api")

	initAPIV1(api, db)

	gracefulshutdown.Listen(app, []func() error{
		func() error {
			fmt.Println("database disconnect")
			sql, err := db.DB()
			if err != nil {
				return err
			}
			return sql.Close()
		},
		func() error {
			fmt.Println("redis disconnect")
			return redis.Client.Close()
		},
	})
}

func initRootHandler(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"version": config.Cfg.App.Version,
			"name":    config.Cfg.App.Name,
		})
	})
}

func initHealthCheck(app fiber.Router, db *gorm.DB) {
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			sql, err := db.DB()
			if err != nil {
				return false
			}
			if err := sql.Ping(); err != nil {
				return false
			}
			if _, err = redis.Client.Ping(c.Context()).Result(); err != nil {
				return false
			}
			return true
		},
		ReadinessEndpoint: "/ready",
	}))
}

func initAPIV1(app fiber.Router, db *gorm.DB) {
	v1 := app.Group("/v1")
	v1.Use(flogger.New(flogger.Config{
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
	}))

	// Operational
	auth.AddRoutes(v1, db)
	student.AddRoutes(v1, db)
	department.AddRoutes(v1, db)
}
