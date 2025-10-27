package routes

import (
	"minify/database"
	"minify/helpers"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL         string `json: "url"`
	CustomShort string `json: "short"`
	Expiry      string `json: "expiry"`
}
type response struct {
	URL         string        `json: "url"`
	CustomShort string        `json: "short"`
	Expiry      time.Duration `json: "expiry"`
	// Rate limit information
	XRateRemaining int           `json: "rate_limit"`
	XRateLimitRest time.Duration `json: "rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	// implement rate limiting here to check the IP of the user and reduce their quota by 1
	// Allow calling the api by 10 time only in 30 mins
	r2 := database.CreateClient(1)
	defer r2.Close()
	r2.Get(database.Ctx, c.IP()).Result()
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	// Step: 1 check if the URL is valid or not
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}
	// Step:2 check for Domain Error(i.e. if the user is trying to shorten our domain link)
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Haha, nice try :("})
	}
	// Step: 3 enforce https:// in the URL(SSL)
	body.URL = helpers.EnforceHTTP(body.URL)

	r2.Decr(database.Ctx, c.IP())

}
