package dialect

import "reflect"

var dialectMap  = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(tp reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}
func GetDialect(name string) (Dialect, bool) {
	res, ok := dialectMap[name]
	return res, ok
}
