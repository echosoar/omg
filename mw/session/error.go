package session

type SessionError struct {
	message string
}
func (sessError *SessionError) Error() string {
	return sessError.message;
}