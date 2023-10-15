package hash_tables

type hashTable struct {
	Size int
	Data [][][2]interface{}
}

func NewHashTable(size int) *hashTable {
	return &hashTable{Size: size, Data: make([][][2]interface{}, size)}
}

func (ht *hashTable) Set(key string, value interface{}) {
	addr := ht.hashFn(key)
	if ht.Data[addr] == nil {
		ht.Data[addr] = [][2]interface{}{}
	}
	// can use linked list here to delete item more efficient
	ht.Data[addr] = append(ht.Data[addr], [2]interface{}{key, value})
}

func (ht *hashTable) Get(key string) interface{} {
	currentBucket := ht.Data[ht.hashFn(key)]
	if currentBucket == nil {
		return nil
	}

	for _, item := range currentBucket {
		if item[0] == key {
			return item[1]
		}
	}

	return nil
}

func (ht *hashTable) Keys() []interface{} {
	keysArr := []interface{}{}

	for _, bucket := range ht.Data {
		if bucket == nil {
			continue
		}

		for _, item := range bucket {
			keysArr = append(keysArr, item[0])
		}
	}

	return keysArr
}

func (ht *hashTable) hashFn(key string) int {
	hash := 0
	for i, r := range key {
		hash = (hash + int(r)*i) % ht.Size
	}
	return hash
}
