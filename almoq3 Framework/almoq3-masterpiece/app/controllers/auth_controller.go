package controllers

import (
	"mytest/bootstrap"
	"mytest/app/models"
	"mytest/app/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthController struct{}

func (a *AuthController) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	if err := bootstrap.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Could not create user"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

func (a *AuthController) Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	var user models.User
	if err := bootstrap.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := token.SignedString([]byte("secret")) // Should use APP_KEY
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{"token": t, "user": user})
}

func (a *AuthController) Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	var dbUser models.User
	bootstrap.DB.First(&dbUser, userID)

	return c.JSON(dbUser)
}
