package patreon

type Event string

const (
	Create Event = "members:pledge:create"
	Update Event = "members:pledge:update"
	Delete Event = "members:pledge:delete"
)
