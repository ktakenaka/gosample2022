package sample

type ListSamplessRequest interface {
	GetSampleIDs() []uint32
	IsWithComments() bool
}
