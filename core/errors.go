package core

type InvalidAuthToken struct {
	Issue string
}

func (i InvalidAuthToken) Error() string {
	return i.Issue
}
