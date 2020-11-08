package ioc

type IoCError struct {
	message string
}
func (ioCError *IoCError) Error() string {
	return ioCError.message;
}