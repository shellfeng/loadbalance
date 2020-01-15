package loadbalance

type Loader interface {
	Get() Item
	Add(items ...Item)
	Delete(index int)
}

type Item interface {
	GetWeight() int
}

// WeightRoundLoader
func WeightRoundLoader() Loader {
	return &WeightRound{}
}
