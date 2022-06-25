package request

import uuid "github.com/satori/go.uuid"

type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	AuthorityId string
}
