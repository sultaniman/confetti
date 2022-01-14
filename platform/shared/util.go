package shared

func String(val string) *string {
	return &val
}

func Int(val int) *int {
	return &val
}

func Int64(val int64) *int64 {
	return &val
}

func IntToInt64(val int) *int64 {
	ival := int64(val)
	return &ival
}

func Bool(val bool) *bool {
	return &val
}

func UInt32(val uint32) *uint32 {
	return &val
}
