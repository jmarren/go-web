package server

import (
	"net/http"

	"github.com/jmarren/go-web/internal/handlers"
)

func addRoutes(
	mux *http.ServeMux,
	// logger *logging.Logger,
	// config Config,
	// tenantsStore *TenantsStore,
	// commentsStore *CommentsStore,
	// conversationService *ConversationService,
	// chatGPTService *ChatGPTService,
	// authProxy *authProxy,
) {

	/* Serve static Files */
	dir := http.Dir("/home/john-marren/templates/go-web/web/public")
	fs := http.FileServer(dir)

	home := handlers.NewHome()

	mux.Handle("GET /static/", http.StripPrefix("/static", fs))
	mux.Handle("GET /about", handlers.NewAbout())
	mux.Handle("GET /home", home)
	// http.FileServerFS

	// fs := http.FileServer(staticDir)
	// mux.Handle("GET /static/assets/", http.StripPrefix("/static/assets/", fs))
	mux.Handle("/", home)

	// mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
	// mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
	// mux.HandleFunc("/healthz", handleHealthzPlease(logger))
	// mux.Handle("/", http.NotFoundHandler())
}
