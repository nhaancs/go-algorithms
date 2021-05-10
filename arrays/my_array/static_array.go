package my_array

type myStringArray struct {
	items map[int]string
	capacity int
}

func NewMyStringArray(capacity int) *myStringArray {
	array := myStringArray{capacity: capacity, items: make(map[int]string)}
	for i := 0; i < array.capacity; i++ {
		array.items[i] = ""
	}
	return &array
}

// access
func (mia *myStringArray) Access(index int) string {
	if index > mia.capacity-1 {
		panic("index out of range")
	}
	return mia.items[index]
}
// insertion
func (mia *myStringArray) Insert(index int, value string) {
	if index > mia.capacity-1 {
		panic("index out of range")
	}
	for i := mia.capacity-1; i > index; i-- {
		mia.items[i] = mia.items[i-1]
	}
	mia.items[index] = value
}
// deletion
func (mia *myStringArray) Delete(index int) {
	if index > mia.capacity-1 {
		panic("index out of range")
	}
	for i := index; i < mia.capacity-1; i++ {
		mia.items[i] = mia.items[i+1]
		mia.items[i+1] = ""
	}
}
// search
func (mia *myStringArray) Search(value string) int {
	for k, v := range mia.items {
		if v == value {
			return k
		}
	}
	return -1
}
