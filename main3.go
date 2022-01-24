package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.snapp.ir/dispatching/confisus/internal/config"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Configinfo struct {
	Config config.Config
	Logger *zap.Logger
	Tracer trace.Tracer
}

type Mohtawa struct {
	UserKind string `json:"user_kind"`
	Channels struct {
		Event string `json:"event"`
	} `json:"channels"`
	CleanSession bool   `json:"clean_session"`
	Host         string `json:"host"`
	PingInterval int    `json:"ping_interval"`
	Port         string `json:"port"`
	Protocol     string `json:"protocol"`
	Qos          int    `json:"qos"`
	Timeout      int    `json:"timeout"`
	TLS          bool   `json:"tls"`
	Topics       struct {
		Chat struct {
			Name string `json:"name"`
			Qos  int    `json:"qos"`
		} `json:"chat"`
		DriverLocation struct {
			Name string `json:"name"`
			Qos  int    `json:"qos"`
		} `json:"driver_location"`
		Events struct {
			Name string `json:"name"`
			Qos  int    `json:"qos"`
		} `json:"events"`
		Location struct {
			Interval int    `json:"interval"`
			Name     string `json:"name"`
			Qos      int    `json:"qos"`
		} `json:"location"`
	} `json:"topics"`
}

func (h Configinfo) Handle(c *fiber.Ctx) error {
	_, span := h.Tracer.Start(c.Context(), "handler.configinfo")
	defer span.End()

	configinfo := h.Config.Dispatching

	return c.JSON(configinfo)
}

func (h Configinfo) HandleRTCConfig(c *fiber.Ctx) error {
	_, span := h.Tracer.Start(c.Context(), "handler.rtc")
	defer span.End()
	mohtawa := &Mohtawa{}
	if err := c.BodyParser(mohtawa); err != nil {
		return err
	}

	// use id to generate topics
	ID := c.Params("id")
	if ID == "0" {
		mohtawa.UserKind = "Passenger"
	} else if ID == "1" {
		mohtawa.UserKind = "Driver"
	} else {
		fmt.Println("some err")
	}

	newres := mohtawa

	return c.Status(http.StatusOK).JSON(newres) //nolint:wrapcheck
}

// Register registers the routes of configinfo handler on given fiber group.
func (h Configinfo) Register(g fiber.Router) {
	g.Get("/config", h.Handle)
	g.Post("/rtc/:id", h.HandleRTCConfig)

}
