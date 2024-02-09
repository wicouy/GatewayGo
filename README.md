## Estructura de Directorios:

/gateway
├── main.go
├── handlers
│ └── handler.go
├── middleware
│ └── auth.go
└── logger
└── logger.go

- main.go: Este archivo será el punto de entrada de nuestro Gateway. Aquí configuraremos el servidor HTTP y las rutas.
- handlers/handler.go: Contendrá los controladores para manejar las solicitudes.
- middleware/auth.go: Aquí implementaremos la lógica para validar tokens.
- logger/logger.go: Para registrar quién pidió un token y su duración.
