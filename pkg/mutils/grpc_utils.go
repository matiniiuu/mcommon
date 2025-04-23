package mutils

func PBEnumConvertor[T ~string, PB ~int32](val T, pbmap map[string]int32, dft PB) *PB {
	v, ok := pbmap[string(val)]
	if !ok {
		return &dft
	}
	result := PB(v)
	return &result
}
