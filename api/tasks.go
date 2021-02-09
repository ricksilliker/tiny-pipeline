package api

type Task struct {
	Name string
	Step *PipelineStep
	Assignee *User
	Reviewer *User
	Project *Project
	Status *Status
}

type ShotTask struct {
	Task
	Shot *Shot
}

type AssetTask struct {
	Task
	Asset *Asset
}

type PipelineStep struct {
	Name string
	Code string
}
