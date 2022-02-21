package collection_types

// KeyValueHolder Stores a pair of key & value of type K and V
// used in orderMap & syncedOrderMap
type KeyValueHolder[K comparable, V any] struct {
	key   K
	value V
}
