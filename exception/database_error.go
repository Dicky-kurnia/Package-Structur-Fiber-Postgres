package exception

type DatabaseError struct {
	Message string
}

func (databaseError DatabaseError) Error() string {
	return databaseError.Message
}
