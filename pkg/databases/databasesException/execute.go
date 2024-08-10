package databasesException

type ExecTransaction struct{}

func (e *ExecTransaction) Error() string {
	return "execute a query failed"
}
