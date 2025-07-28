package utils

// Ptr converts any value to its pointer
// Usage examples:
//
//	utils.Ptr(float64(123.45))     // returns *float64
//	utils.Ptr(int(42))             // returns *int
//	utils.Ptr("hello")             // returns *string
func Ptr[T any](value T) *T {
	return &value
}

// SafePtr safely dereferences a pointer with default value
// Usage examples:
//
//	utils.SafePtr(user.TFASecret, "")           // returns string
//	utils.SafePtr(user.Age, 0)                  // returns int
//	utils.SafePtr(user.Score, 0.0)              // returns float64
func SafePtr[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}
