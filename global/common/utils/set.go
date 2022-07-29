package utils

//
//type void struct{}
//
//type Set[T any] struct {
//	members map[any]void
//}
//
//func (receiver *Set[T]) Set(members ...T) *Set[T] {
//	receiver.members = make(map[any]void)
//	for _, member := range members {
//		receiver.members[member] = void{}
//	}
//	return receiver
//}
//
//func (receiver Set[T]) Array() []T {
//	var members []T
//	for member, _ := range receiver.members {
//		members = append(members, member.(T))
//	}
//	return members
//}
