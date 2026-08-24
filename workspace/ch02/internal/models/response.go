package models

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type InfoResponse struct {
	Version       string  `json:"version"`
	GoVersion     string  `json:"go_version"`
	UptimeSeconds float64 `json:"uptime_seconds"`
}
