package docs

import "github.com/GoGerman/geo-task/module/courierfacade/models"

// добавить документацию для роута /api/status

// swagger:route GET /api/status courier GetStatus
// Get courier status
// Responses:
//   200: GetStatusRes200

// swagger:response GetStatusRes200
type CourierResponse struct {
	// in:body
	Body models.CourierStatus
}
