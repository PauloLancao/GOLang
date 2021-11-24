package di

import (
	"github.com/paulolancao/go-contacts/repository"
	"github.com/paulolancao/go-contacts/services"
)

// DI struct
type DI struct {
	ContactService    *services.Service      // -> contactservice -> contactrepository
	ContactRepository *repository.Repository // -> contactrepository
}
