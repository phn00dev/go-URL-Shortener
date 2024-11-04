package constructor

import (
	"github.com/phn00dev/go-URL-Shortener/internal/app"
	adminConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/admin/constructor"
)

func InitDependencies(dependencies *app.Dependencies) {
	adminConstructor.InitAdminRequirements(dependencies.DB)
}
