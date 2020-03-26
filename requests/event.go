package requests

type GenericEventRequest struct {
	Base
	EventID int64 `json:"eventId,string"`
}
