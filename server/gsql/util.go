package gsql

// Commit commits a transaction
func Commit(at Atomic) error {
	return at.End()
}

// RollbackIfFail rollbacks a transaction if the transaction is not successful
func RollbackIfFail(at Atomic, success bool) error {
	if success {
		return nil
	}
	return at.Fail()
}
