package domain

type StorageAuthentication struct {
	ObjectId string
	ClientId string
	Username string
	Password string
	Scopes   []string
}
