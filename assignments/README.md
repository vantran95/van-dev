restful/
├── cmd/
│   └── api/
│       └── main.go               # Application entry point
│       └── router.go             # Application router define
├── internal/
│   ├── config/                   # Configuration handling
│   ├── domain/                   # Business domain models
│   ├── repository/               # Data access layer
│   ├── service/                  # Business logic
│   ├── api/                      # REST API handlers
│   ├── websocket/                # WebSocket handlers
│   └── middleware/               # HTTP middleware
├── pkg/                          # Reusable packages
│   ├── logger/                   # Logging utilities
│   ├── validator/                # Input validation
│   └── errors/                   # Error handling
├── migrations/                   # Database migrations
├── scripts/                      # Build and deployment scripts
├── docker-compose.yml            # Local development environment
├── Dockerfile                    # Container definition
├── go.mod                        # Go module definition
└── README.md                     # Project documentation