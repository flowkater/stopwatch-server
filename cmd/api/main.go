package main

import (
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
	userRepo := repository.NewUserMemoryRepository()
	stopwatchRepo := repository.NewStopwatchMemoryRepository()

	// 2. 인프라 구성요소
	clientManager := websocket.NewClientManager()

	// 3. 서비스
	userService := services.NewUserService(userRepo, clientManager)
	stopwatchService := services.NewStopwatchService(stopwatchRepo, clientManager)

	// 4. 핸들러
	wsHandler := handlers.NewWebSocketHandler(userService, stopwatchService, clientManager)

	// 5. 애플리케이션 구성
	app := fiber.New()

	// 6. 라우터 등록
	wsHandler.Register(app)

	// 7. 서버 시작
	log.Println("서버 시작: 포트", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
