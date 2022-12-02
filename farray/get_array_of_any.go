package farray

// getSpecificModel
// generic func wil take an index of a model and will
// return a specific data define in every case
type getSpecificModel[Model any, Result any] func(Model) Result

// KeepFunc
// generic func wil take an index of a model and will
// return a boolean result based on condition
type KeepFunc[E any] func(E) bool

// GetArrayOfAny
// take an array of any and will return am array of specific object
//
// Example:
//
//	type One struct {
//		id int
//	}
//	type Two struct {
//		id string
//	}
//	type Data struct {
//		o One
//		t Two
//	}
//
//	ex := []*Data{
//		{o: One{id: 1}, t: Two{id: "one"}},
//		{o: One{id: 2}},
//		{o: One{id: 3}, t: Two{id: "three"}},
//	}
//
//	result := GetArrayOfAny[Data, string](ex, func(item *Data) string {
//		return item.t.id
//	})
//
// result -> [one three]
func GetArrayOfAny[Model any, Result comparable](model []*Model, f getSpecificModel[*Model, Result]) (result []Result) {

	//check every index and return an element based on passed func
	for i := range model {
		result = append(result, f(model[i]))
	}
	//filter model to remove nil data
	result = Filter[Result](result, func(res Result) bool {
		return res != SetDefaultValueOfNil(res)
	})

	//return final data
	return
}

// Filter
// take an array and return the one who confirm the condition
// passed by the func KeepFunc
func Filter[Model any](model []Model, f KeepFunc[Model]) (result []Model) {

	// if condition is true save the item
	for _, item := range model {
		if f(item) {
			result = append(result, item)
		}
	}
	//return final data
	return
}

// SetDefaultValueOfNil
// return default value type
func SetDefaultValueOfNil(variable any) any {
	switch variable.(type) {
	case bool:
		return false
	case string:
		return ""
	case int:
		return 0
	default:
		return nil
	}
}
