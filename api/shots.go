package api

type Shot struct {
	Name string
	Status *Status
	Project *Project
	Sequence string
	CutIn int
	CutOut int
	HeadIn int
	TailOut int
	AssetItems map[string]*Asset
}