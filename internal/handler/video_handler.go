package handler

import (
	videoService "fampay_youtube_fetcher/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type VideoHandler struct {
	service *videoService.VideoService
}

func NewVideoHandler(service *videoService.VideoService) *VideoHandler {
	return &VideoHandler{service: service}
}

func (h *VideoHandler) GetLatestVideos(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	videos, err := h.service.GetLatestVideos(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"videos": videos,
	})
}

func (h *VideoHandler) SearchVideos(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "search query is required",
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	videos, err := h.service.SearchVideos(query, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"videos": videos,
	})
}
