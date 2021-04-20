package sort_example


type HashMapAndSortedMap struct {
	HashMap map[int]string `json:"hash_map"`
}

func NewHashMapAndSortedMap(limit int)*HashMapAndSortedMap  {
		return &HashMapAndSortedMap{
			HashMap: make(map[int]string,limit),
		}
}

func (h *HashMapAndSortedMap) ContainsKey(key int) bool  {
	_,isEmpty := h.HashMap[key]
	return isEmpty
}

func (h *HashMapAndSortedMap) DelValue(key int)  {
	delete(h.HashMap, key)
}

func (h *HashMapAndSortedMap) SetValue(key int,value string)  {
	h.HashMap[key] = value
}

