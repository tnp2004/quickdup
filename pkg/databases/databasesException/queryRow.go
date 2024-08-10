package databasesException

type QueryRow struct{}

func (e *QueryRow) Error() string {
	return "query row failed"
}
