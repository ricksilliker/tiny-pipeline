package api

type Color struct {
	R int
	G int
	B int
}

type Phase struct {
	ID int
	Name string
	Color Color
}

type Status struct {
	ID int
	Name string
	Color
}

type Tracker struct {
	StartedAt string
	DueAt string
	ApprovedAt string
	Estimation int
	Task *Task
	Phase *Phase
}

type AssetTracker struct {
	Tracker
	Asset *Asset
}

type ShotTracker struct {
	Tracker
	Shot *Shot
	Rating string
}

type ClientTracker struct {
	SubmittedAt string
	ResponsedAt string
	Phase *Phase
	Status *Status
	Caveat string
	Notes []string
	ActionItems []string
	Color *Color
}

type AssetClientTracker struct {
	ClientTracker
	Asset *Asset
}

type ShotClientTracker struct {
	ClientTracker
	Shot *Shot
}