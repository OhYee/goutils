package condition

//go:generate gcg ./template/condition.json
type any = interface{}
type function = func()
