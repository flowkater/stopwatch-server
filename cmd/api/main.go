package main

import (
	"encoding/json"
	"log"

	"github.com/flowkater/stopwatch-server/internal/application/services"
	"github.com/flowkater/stopwatch-server/internal/infrastructure/repository"
	"github.com/flowkater/stopwatch-server/internal/infrastructure/websocket"
	"github.com/flowkater/stopwatch-server/internal/interfaces/handlers"
	"github.com/flowkater/stopwatch-server/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 설정 로드
	cfg := config.NewConfig()

	// 의존성 주입
	// 1. 레포지토리
	userStateRepo := repository.NewUserStateMemoryRepository()

	// 2. 인프라 구성요소
	clientManager := websocket.NewClientManager()

	// 3. 서비스
	userStateService := services.NewUserStateService(userStateRepo, clientManager)

	// 4. 핸들러
	wsHandler := handlers.NewWebSocketHandler(userStateService, clientManager)
	userStateHandler := handlers.NewUserStateHandler(userStateService)

	// 5. 애플리케이션 구성
	app := fiber.New(fiber.Config{
		// 기본 JSON 인코딩 설정
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// 모든 응답에 대한 기본 미들웨어 설정
	app.Use(func(c *fiber.Ctx) error {
		// 기본 Content-Type 설정
		c.Set("Content-Type", "application/json; charset=utf-8")
		return c.Next()
	})

	// 6. 라우터 등록
	wsHandler.Register(app)
	userStateHandler.Register(app)

	// 7. 서버 시작
	log.Println("서버 시작: 포트", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
