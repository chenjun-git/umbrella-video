package v1_0

import (
	"fmt"

	"github.com/go-chi/chi"

	"github.com/chenjun-git/umbrella-common/monitor"
)

func RegisterRouter(r chi.Router) {
	registerRouter("v1.0", r)
}

func registerRouter(version string, r chi.Router) {
}