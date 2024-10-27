package query

type Command interface {
	Execute() map[string]interface{}
}
