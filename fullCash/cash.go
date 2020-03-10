package fullCash

//
// import (
// 	"runtime"
// )
//
// type ResSerializer interface {
// 	GetNames(int) ([]string, error)
// 	Serialize(data interface{}) error
// }
//
// type serializer struct {
// }

//
// func (s serializer) GetNames(namesNum []int) ([]string, error) {
// 	pc := make([]uintptr, 10) // at least 1 entry needed
// 	runtime.Callers(2, pc)
// 	f := runtime.FuncForPC(pc[0])
// 	f.Name()
//
// 	return nil, nil
// }
//
// func (s serializer) Serialize(data interface{}) error {
// 	panic("implement me")
// }
//
// func NewResSerializer() ResSerializer {
// 	return &serializer{}
// }
//
