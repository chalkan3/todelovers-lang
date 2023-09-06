package tvm

type pointer struct {
	adress int64
}

func newPointer(adress int64) *pointer {
	return &pointer{
		adress: adress,
	}
}

func (pp *pointer) GetAddress() int64 { return pp.adress }
