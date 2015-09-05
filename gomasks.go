package gomasks

type Bitmask uint32

func (f Bitmask) HasFlag(flag Bitmask) bool {
	return f&flag != 0
}

func (f *Bitmask) AddFlag(flag Bitmask) {
	*f = (*f) | flag
}

func (f *Bitmask) RemoveFlag(flag Bitmask) {
	*f = (*f) &^ flag
}
