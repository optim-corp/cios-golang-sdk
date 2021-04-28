package xmath

import "sort"

func SumStr(args ...string) (result string) {
	for _, arg := range args {
		result += arg
	}
	return
}
func SumInt(args ...int) (result int) {
	for _, arg := range args {
		result += arg
	}
	return
}
func SumInt32(args ...int32) (result int32) {
	for _, arg := range args {
		result += arg
	}
	return
}
func SumInt64(args ...int64) (result int64) {
	for _, arg := range args {
		result += arg
	}
	return
}
func SumFloat32(args ...float32) (result float32) {
	for _, arg := range args {
		result += arg
	}
	return
}
func SumFloat64(args ...float64) (result float64) {
	for _, arg := range args {
		result += arg
	}
	return
}

func AverageInt(args ...int) int {
	v := 0
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / len(args)
}
func AverageInt32(args ...int32) int32 {
	v := int32(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / int32(len(args))
}
func AverageInt64(args ...int64) int64 {
	v := int64(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / int64(len(args))
}
func AverageFloat32(args ...float32) float32 {
	v := float32(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / float32(len(args))
}
func AverageFloat64(args ...float64) float64 {
	v := float64(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / float64(len(args))
}

func MedianInt(args ...int) *int {
	length := len(args)
	if length == 0 {
		return nil
	}
	half := length / 2
	result := 0
	if sort.Slice(args, func(i, j int) bool { return args[i] > args[j] }); length%2 == 0 {
		result = (args[half] + args[half+1]) / 2
	} else {
		result = args[half+1]
	}
	return &result
}
func MedianInt32(args ...int32) *int32 {
	length := len(args)
	if length == 0 {
		return nil
	}
	half := length / 2
	result := int32(0)
	if sort.Slice(args, func(i, j int) bool { return args[i] > args[j] }); length%2 == 0 {
		result = (args[half] + args[half+1]) / 2
	} else {
		result = args[half+1]
	}
	return &result
}
func MedianInt64(args ...int64) *int64 {
	length := len(args)
	if length == 0 {
		return nil
	}
	half := length / 2
	result := int64(0)
	if sort.Slice(args, func(i, j int) bool { return args[i] > args[j] }); length%2 == 0 {
		result = (args[half] + args[half+1]) / 2
	} else {
		result = args[half+1]
	}
	return &result
}
func MedianFloat32(args ...float32) *float32 {
	length := len(args)
	if length == 0 {
		return nil
	}
	half := length / 2
	result := float32(0)
	if sort.Slice(args, func(i, j int) bool { return args[i] > args[j] }); length%2 == 0 {
		result = (args[half] + args[half+1]) / 2
	} else {
		result = args[half+1]
	}
	return &result
}
func MedianFloat64(args ...float64) *float64 {
	length := len(args)
	if length == 0 {
		return nil
	}
	half := length / 2
	result := float64(0)
	if sort.Slice(args, func(i, j int) bool { return args[i] > args[j] }); length%2 == 0 {
		result = (args[half] + args[half+1]) / 2
	} else {
		result = args[half+1]
	}
	return &result
}

func ModeInt(args ...int) *int {
	if len(args) == 0 {
		return nil
	}
	m := map[int]int{}
	result := 0
	max := -1
	for _, arg := range args {
		if v, ok := m[arg]; ok {
			m[arg] = v + 1
		} else {
			m[arg] = 0
		}
		max = MaxInt(max, m[arg])
		if max == m[arg] {
			result = arg
		}
	}
	return &result
}
func ModeInt32(args ...int32) *int32 {
	if len(args) == 0 {
		return nil
	}
	m := map[int32]int{}
	result := int32(0)
	max := -1
	for _, arg := range args {
		if v, ok := m[arg]; ok {
			m[arg] = v + 1
		} else {
			m[arg] = 0
		}
		max = MaxInt(max, m[arg])
		if max == m[arg] {
			result = arg
		}
	}
	return &result
}
func ModeInt64(args ...int64) *int64 {
	if len(args) == 0 {
		return nil
	}
	m := map[int64]int{}
	result := int64(0)
	max := -1
	for _, arg := range args {
		if v, ok := m[arg]; ok {
			m[arg] = v + 1
		} else {
			m[arg] = 0
		}
		max = MaxInt(max, m[arg])
		if max == m[arg] {
			result = arg
		}
	}
	return &result
}
func ModeFloat32(args ...float32) *float32 {
	if len(args) == 0 {
		return nil
	}
	m := map[float32]int{}
	result := float32(0)
	max := -1
	for _, arg := range args {
		if v, ok := m[arg]; ok {
			m[arg] = v + 1
		} else {
			m[arg] = 0
		}
		max = MaxInt(max, m[arg])
		if max == m[arg] {
			result = arg
		}
	}
	return &result
}
func ModeFloat64(args ...float64) *float64 {
	if len(args) == 0 {
		return nil
	}
	m := map[float64]int{}
	result := float64(0)
	max := -1
	for _, arg := range args {
		if v, ok := m[arg]; ok {
			m[arg] = v + 1
		} else {
			m[arg] = 0
		}
		max = MaxInt(max, m[arg])
		if max == m[arg] {
			result = arg
		}
	}
	return &result
}

func RangeInt(args ...int) *int {
	if len(args) == 0 {
		return nil
	}
	min := args[0]
	max := args[0]
	for i := 1; i < len(args); i++ {
		min = MinInt(min, args[i])
		max = MaxInt(max, args[i])
	}
	result := max - min
	return &result
}
func RangeInt32(args ...int32) *int32 {
	if len(args) == 0 {
		return nil
	}
	min := args[0]
	max := args[0]
	for i := 1; i < len(args); i++ {
		min = MinInt32(min, args[i])
		max = MaxInt32(max, args[i])
	}
	result := max - min
	return &result
}
func RangeInt64(args ...int64) *int64 {
	if len(args) == 0 {
		return nil
	}
	min := args[0]
	max := args[0]
	for i := 1; i < len(args); i++ {
		min = MinInt64(min, args[i])
		max = MaxInt64(max, args[i])
	}
	result := max - min
	return &result
}
func RangeFloat32(args ...float32) *float32 {
	if len(args) == 0 {
		return nil
	}
	min := args[0]
	max := args[0]
	for i := 1; i < len(args); i++ {
		min = MinFloat32(min, args[i])
		max = MaxFloat32(max, args[i])
	}
	result := max - min
	return &result
}
func RangeFloat64(args ...float64) *float64 {
	if len(args) == 0 {
		return nil
	}
	min := args[0]
	max := args[0]
	for i := 1; i < len(args); i++ {
		min = MinFloat64(min, args[i])
		max = MaxFloat64(max, args[i])
	}
	result := max - min
	return &result
}
