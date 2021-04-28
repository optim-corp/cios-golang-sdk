package xmath

func MaxStr(args ...string) string {
	if len(args) == 0 {
		return ""
	}
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}
func MinStr(args ...string) string {
	if len(args) == 0 {
		return ""
	}
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}

func MaxInt64(args ...int64) int64 {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}
func MinInt64(args ...int64) int64 {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}

func MaxInt32(args ...int32) int32 {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}
func MinInt32(args ...int32) int32 {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}

func MaxInt(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}
func MinInt(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min

}

func MaxFloat32(args ...float32) float32 {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}
func MinFloat32(args ...float32) float32 {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min

}

func MaxFloat64(args ...float64) float64 {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}
func MinFloat64(args ...float64) float64 {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}
