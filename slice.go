package slice_tool

// IsEmpty if slice is empty
func IsEmpty[T any](s []T) bool {
	return len(s) == 0
}

// IsNotEmpty if slice is not empty
func IsNotEmpty[T any](s []T) bool {
	return len(s) != 0
}

// Remove removes element from slice, return new slice,won't change original slice
func Remove[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

// Insert elements to position, return new slice,won't change original slice
func Insert[T any](s []T, index int, item ...T) []T {
	return append(append(s[:index], item...), s[index:]...)
}

// Search find element in slice, return index, if not found return -1
func Search[T comparable](s []T, item T) int {
	for i, i2 := range s {
		if i2 == item {
			return i
		}
	}
	return -1
}

// Contains check if slice contains element
func Contains[T comparable](s []T, item T) bool {
	for _, i2 := range s {
		if i2 == item {
			return true
		}
	}
	return false
}

//Unique remove duplicate elements, return new slice,won't change original slice
func Unique[T comparable](s []T) []T {
	newArr := make([]T, 0)
	for i := 0; i < len(s); i++ {
		if Contains(newArr, s[i]) == false {
			newArr = append(newArr, s[i])
		}
	}
	return newArr
}

// Filter elements in slice by filter func, return new slice,won't change original slice
func Filter[T comparable](s []T, filter func(item T, index int) bool) []T {
	var newSlice = make([]T, 0)
	for i := 0; i < len(s); i++ {
		if filter(s[i], i) {
			newSlice = append(newSlice, s[i])
		}
	}
	return newSlice
}

// Reject elements in slice by reject func, return new slice,won't change original slice
func Reject[T comparable](s []T, reject func(item T, index int) bool) []T {
	var newSlice = make([]T, 0)
	for i := 0; i < len(s); i++ {
		if !reject(s[i], i) {
			newSlice = append(newSlice, s[i])
		}
	}
	return newSlice
}

// Copy make a copy of slice
func Copy[T comparable](s []T) []T {
	var slice2 = make([]T, len(s))
	copy(slice2, s)
	return slice2
}

// Each call func for each element in slice
func Each[T comparable](s []T, function func(item T, index int)) {
	for i, t := range s {
		function(t, i)
	}
}

// Map every element in slice by mapper func
func Map[T comparable, K any](s []T, mapper func(item T, index int) K) []K {
	var newSlice = make([]K, len(s))
	for i := 0; i < len(s); i++ {
		newSlice[i] = mapper(s[i], i)
	}
	return newSlice
}

// Every check every element in slice by check func
func Every[T comparable](s []T, function func(item T, index int) bool) bool {
	for i, t := range s {
		if !function(t, i) {
			return false
		}
	}
	return true
}

// ForPage return new slice by page and page size, page starts from 1
func ForPage[T comparable](s []T, page int, perPage int) []T {
	start := page * perPage
	end := min((page+1)*perPage, len(s))

	return s[start:end]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Nth start form offset, return nth element in slice
func Nth[T comparable](s []T, n int, offset int) []T {
	var newSlice = make([]T, 0)
	for i := offset; i < len(s); i++ {
		if (i-offset)%n == 0 {
			newSlice = append(newSlice, s[i])
		}
	}
	return newSlice
}

// Pad insert element to slice, return new slice,won't change original slice
func Pad[T comparable](s []T, count int, def T) []T {
	var newSlice = s
	pad := count - len(s)
	for i := 0; i < pad; i++ {
		newSlice = append(newSlice, def)
	}
	return newSlice
}

// Pop remove last element in slice and return it,will change original slice
func Pop[T comparable](s *[]T) T {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

// Push insert element to slice's tail,return new slice length,will change original slice
func Push[T comparable](s *[]T, val T) int {
	*s = append(*s, val)
	return len(*s)
}

// Shift remove first element in slice and return it,will change original slice
func Shift[T comparable](s *[]T) T {
	res := (*s)[0]
	*s = (*s)[1:]
	return res
}

// UnShift insert element to slice's head,return new slice length,will change original slice
func UnShift[T comparable](s *[]T, val T) int {
	*s = append([]T{val}, *s...)
	return len(*s)
}
