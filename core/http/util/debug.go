package util

import (
	"fmt"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func DebugRoute(prefix string, route chi.Routes) {
	for _, method := range route.Routes() {
		patternIsClean := !strings.HasSuffix(method.Pattern, "/*")
		var cleanPattern string
		if patternIsClean {
			cleanPattern = method.Pattern
		} else {
			cleanPattern = strings.TrimSuffix(method.Pattern, "/*")
		}
		path := fmt.Sprintf("%s%s", prefix, cleanPattern)
		for k := range method.Handlers {
			if patternIsClean {
				log.Debug().Str("path", path).Str("method", k).Msg("Registered route")
			}
		}

		if method.SubRoutes != nil {
			DebugRoute(path, method.SubRoutes)
		}
	}
}
