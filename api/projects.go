package api

type ProjectType string

const (
	IP ProjectType = "IP"
	Feature = "Feature"
	Visualization = "Visualization"
	Game = "Game"
	Commercial = "Commercial"
)

type Project struct {
	ID 			int
	Name 		string
	Type 		ProjectType
	Settings 	*ProjectSettings
	Users 		[]User
	Description string
	Deliveries 	[]Deliverable
	StartedAt	string
	CompletedAt	string
}

type ProjectSettings struct {}

type RenderResolution struct {
	Width 	int
	Height 	int
}

type Deliverable struct {
	FPS 		int
	Resolution 	RenderResolution
	DueAt 		string
	VideoFormat	FileType
	AudioFormat	FileType
	Colorspace 	string
	TotalFrames int
}
