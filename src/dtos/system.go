package dtos

type APIInfoResponse struct {
	Name        string   `json:"name" example:"Go Auth API"`
	Version     string   `json:"version" example:"1.1.0"`
	Description string   `json:"description" example:"JWT authentication API built with Gin and MongoDB."`
	DocsURL     string   `json:"docsUrl" example:"/swagger/index.html"`
	Endpoints   []string `json:"endpoints" example:"/health,/api/v1/users,/api/v1/users/login"`
}

type HealthResponse struct {
	Status    string `json:"status" example:"ok"`
	Service   string `json:"service" example:"Go Auth API"`
	Version   string `json:"version" example:"1.1.0"`
	Timestamp string `json:"timestamp" example:"2026-04-17T14:00:00Z"`
}
