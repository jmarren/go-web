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

	mux.Handle("/", handlers.NewHome())

	// mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
	// mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
	// mux.HandleFunc("/healthz", handleHealthzPlease(logger))
	// mux.Handle("/", http.NotFoundHandler())
}
