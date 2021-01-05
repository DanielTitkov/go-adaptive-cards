package cards

// BoolPtr returns pointer to bool
func BoolPtr(b bool) *bool {
	return &b
}

// TruePtr returns pointer to true
func TruePtr() *bool {
	return BoolPtr(true)
}

// FalsePtr returns pointer to false
func FalsePtr() *bool {
	return BoolPtr(false)
}
