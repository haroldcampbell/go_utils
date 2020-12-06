package utils

type StrArray []string

// var strArrayMutex sync.Mutex

func (items StrArray) Duplicate() StrArray {
	itemCount := len(items)
	newItems := make(StrArray, itemCount)
	for i := 0; i < itemCount; i++ {
		newItems = append(newItems, items[i])
	}
	return newItems
}

func (items StrArray) IndexOf(elm string) int {
	for k, v := range items {
		if elm == v {
			return k
		}
	}

	return -1
}

// DeleteItemAt deletes the item at the specified index
func (items StrArray) DeleteItemAt(i int) StrArray {
	if i < 0 || i >= len(items) {
		return items
	}

	return append(items[:i], items[i+1:]...)
}
