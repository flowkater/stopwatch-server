# Stopwatch Server

실시간 스탑워치 상태를 공유하는 WebSocket 서버입니다. 사용자들의 스탑워치 상태(실행 중, 경과 시간 등)를 실시간으로 공유합니다.

## 기술 스택

- Go 1.23+
- Fiber 웹 프레임워크
- WebSocket

## 프로젝트 구조

DDD(Domain-Driven Design) 아키텍처를 따릅니다:

```
.
├── cmd/
│   └── api/               # 애플리케이션 진입점
│       └── main.go        # 메인 함수
├── internal/
│   ├── domain/            # 도메인 모델과 비즈니스 규칙
│   │   ├── message/       # 메시지 도메인
│   │   ├── user/          # 사용자 도메인
│   │   └── stopwatch/     # 스탑워치 도메인
│   ├── application/       # 애플리케이션 서비스 및 유스케이스
│   │   └── services/      # 애플리케이션 서비스
│   ├── infrastructure/    # 외부 시스템 연동 코드
│   │   ├── repository/    # 레포지토리 구현체
│   │   └── websocket/     # 웹소켓 기능 구현
│   └── interfaces/        # 외부 인터페이스
│       └── handlers/      # 핸들러 (컨트롤러)
└── pkg/
    └── config/            # 설정 관련 코드
```

## 실행 방법

1. 저장소 클론:

```
git clone https://github.com/your-username/stopwatch-server.git
cd stopwatch-server
```

2. 의존성 설치:

```
go mod download
```

3. 서버 실행:

```
go run cmd/api/main.go
```

4. 서버가 3000번 포트에서 실행됩니다.

## WebSocket API

### 연결

```
ws://localhost:3000/ws
```

### 메시지 형식

```json
{
  "userId": "user-123",
  "online": true,
  "stopwatchRunning": true,
  "elapsedTime": 120.5
}
```

- `userId`: 사용자 ID
- `online`: 온라인 상태
- `stopwatchRunning`: 스탑워치 실행 중 여부
- `elapsedTime`: 경과 시간(초 단위)
