package v1_0

import (
	//"fmt"

	"github.com/go-chi/chi"

	"github.com/chenjun-git/umbrella-common/monitor"
)

func RegisterRouter(r chi.Router) {
	registerRouter("v1.0", r)
}

func registerRouter(version string, r chi.Router) {
	r.Route("/api/v1.0/video", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", monitor.HttpHandlerWrapper("GetVideoById", getVideoById))
			r.Put("/", monitor.HttpHandlerWrapper("UpdateVideoById", updateVideoById))
			r.Delete("/", monitor.HttpHandlerWrapper("DeleteVideoById", deleteVideoById))
		})

		r.Post("/video", monitor.HttpHandlerWrapper("AddVideo", addVideo))
	})

	r.Route("/api/v1.0/videos", func(r chi.Router) {
		r.Get("/", monitor.HttpHandlerWrapper("GetVideos", getVideos))
	})

	r.Route("/api/v1.0/video_theme", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", monitor.HttpHandlerWrapper("GetVideoThemeById", getVideoThemeById))
			r.Put("/", monitor.HttpHandlerWrapper("UpdateVideoThemeById", updateVideoThemeById))
			r.Delete("/", monitor.HttpHandlerWrapper("DeleteVideoThemeById", deleteVideoThemeById))
		})

		r.Post("/", monitor.HttpHandlerWrapper("AddVideoTheme", addVideoTheme))
	})

	r.Route("/api/v1.0/video_themes", func(r chi.Router) {
		r.Get("", monitor.HttpHandlerWrapper("GetVideoThemes", getVideoThemes))
	})
}

func wrapCode(v interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    0,
		"message": "ok",
		"body":    v,
	}
}
