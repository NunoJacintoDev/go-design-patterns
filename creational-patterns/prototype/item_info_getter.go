package prototype

type ItemInfoGetter interface {
	GetInfo() string
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}
