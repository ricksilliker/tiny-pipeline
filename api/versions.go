package api

type VersionType int
const (
	Review VersionType = iota
	Work
	Publish
)

type FileType int
const (
	Alembic FileType = iota
	MayaScene
	FBX
	PSD
	Zip
	JPG
	PNG
	MOV
	PDF
)

type VersionStatus int
const (
	Standby = iota
	UpNext
	InReview
	Reviewed
)

type Version struct {
	ID int
	Type VersionType
	Project *Project
	SubmittedBy *User
	Media []*File
	Number int
	Status VersionStatus

}

type File struct {
	URI string
	Name string
	Type FileType
	StartFrame int
	EndFrame int
	Size int64
}