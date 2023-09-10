package cast

func ToInt(cast interface{}) int { return cast.(int) }
func ToUInt8(cast interface{}) (uint8, bool) {
	c, ok := cast.(uint8)
	return c, ok
}

func ToByte(cast interface{}) byte        { return cast.(byte) }
func ToByteArray(cast interface{}) []byte { return cast.([]byte) }

func ToString(cast interface{}) (string, bool) {
	s, err := cast.(string)
	return s, err
}

func ToAlwaysInt(cast interface{}) int {
	castUInt8, ok := ToUInt8(cast)
	if ok {
		return int(castUInt8)
	}

	return ToInt(cast)
}
