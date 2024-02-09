package logger

import "log"

// LogRequest registra la solicitud de token
func LogRequest(userID string, tokenDuration int) {
    log.Printf("Token solicitado por %s, durará %d segundos", userID, tokenDuration)
}
