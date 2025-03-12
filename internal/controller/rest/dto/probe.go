package dto

// HealthResponse describe health probe of service.
type HealthResponse struct {
	Status string `json:"status"`
}

// ReadyResponse describe ready probe of service.
type ReadyResponse struct {
	Status string `json:"status"`
}
