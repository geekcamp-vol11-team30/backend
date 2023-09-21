package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Event struct {
	ID            ulid.ULID         `json:"id"`
	OwnerID       ulid.ULID         `json:"ownerId"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	DurationAbout string            `json:"durationAbout"`
	UnitSeconds   int               `json:"unitSeconds"`
	Units         []EventTimeUnit   `json:"units"`
	UserAnswers   []UserEventAnswer `json:"userAnswers"`
}
type EventRequest struct {
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	DurationAbout string                 `json:"durationAbout"`
	UnitSeconds   int                    `json:"unitDuration"`
	Units         []EventTimeUnitRequest `json:"units"`
}
type EventResponse struct {
	ID            string                    `json:"id"`
	OwnerID       string                    `json:"ownerId"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	DurationAbout string                    `json:"durationAbout"`
	UnitSeconds   int                       `json:"unitDuration"`
	Units         []EventTimeUnitResponse   `json:"units"`
	UserAnswers   []UserEventAnswerResponse `json:"userAnswers"`
}

type EventTimeUnit struct {
	ID          ulid.ULID `json:"id"`
	EventID     ulid.ULID `json:"-"`
	TimeSlot    time.Time `json:"timeSlot"`
	SlotSeconds int       `json:"slotSeconds"`
}
type EventTimeUnitRequest struct {
	TimeSlot time.Time `json:"timeSlot"`
}
type EventTimeUnitResponse struct {
	ID       string    `json:"id"`
	StartsAt time.Time `json:"startsAt"`
	EndsAt   time.Time `json:"endsAt"`
}

type Confirm struct {
	ID                   ulid.ULID `json:"id"`
	EventID              ulid.ULID `json:"-"`
	NotifyByEmail        bool      `json:"notifyByEmail"`
	NumberOfParticipants int       `json:"numberOfParticipants"`
	ConfirmationEmail    string    `json:"confirmationEmail"`
}
