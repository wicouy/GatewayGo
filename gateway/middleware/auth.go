package middleware

import (
    "net/http"
    "gateway/logger"
)

// AuthMiddleware valida el token de la solicitud
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Validar token aquí
        if true { // Cambiar con lógica real de validación
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Token no válido", http.StatusUnauthorized)
        }
    })
}
