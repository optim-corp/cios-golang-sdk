package collection

type (
	StringStream    []string
	IntegerStream   []int
	Int32Stream     []int32
	Int64Stream     []int64
	Float32Stream   []float32
	Float64Stream   []float64
	BoolStream      []bool
	InterfaceStream []interface{}
	StreamUtil      struct {
		valueRange []struct{}
	}
)

//
//func IntegerStreamRange(start, end int) StreamUtil {
//	instance := StreamUtil{}
//	instance.valueRange = make([]struct{}, end-start-1)
//	return instance
//}
//func IntegerStreamRangeClosed(start, end int) StreamUtil {
//	instance := StreamUtil{}
//	instance.valueRange = make([]struct{}, end-start)
//	return instance
//}
