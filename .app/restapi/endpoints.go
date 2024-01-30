package restapi

import (
	"log"
	"net/http"
	"time"

	"github.com/365admin/sharepoint-webparts/app/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4emb"
)

type ListNotication struct {
	Value []struct {
		SubscriptionID     string    `json:"subscriptionId"`
		ClientState        string    `json:"clientState"`
		ExpirationDateTime time.Time `json:"expirationDateTime"`
		Resource           string    `json:"resource"`
		TenantID           string    `json:"tenantId"`
		SiteURL            string    `json:"siteUrl"`
		WebID              string    `json:"webId"`
	} `json:"value"`
}

const description = `
	
Service  for managing Microsoft 365 resources

## Getting started 

### Authentication
You need a credential key to access the API. The credential is issue by [niels.johansen@nexigroup.com](mailto:niels.johansen@nexigroup.com).

Use the credential key to get an access token through the /v1/authorize end point. The access token is valid for 10 minutes.

Pass the access token in the Authorization header as a Bearer token to access the API.
	`

func sharedSettings(s *web.Service) {
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)
	s.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON:     false,
		LogLevel: "info",
	})
	s.Use(httplog.RequestLogger(logger))

	s.Post("/authorize", signin())
}

func rateLimitByAppId(maxRequestsPerMinute int) func(next http.Handler) http.Handler {
	return httprate.Limit(
		maxRequestsPerMinute, // requests
		1*time.Minute,        // per duration
		// an oversimplified example of rate limiting by a custom header
		httprate.WithKeyFuncs(func(r *http.Request) (string, error) {

			token := r.Context().Value("auth").(core.Authorization).AppId
			return token, nil
		}),
	)
}
func rateLimitByIpAddress(maxRequestsPerMinute int) func(next http.Handler) http.Handler {
	return httprate.Limit(
		maxRequestsPerMinute, // requests
		1*time.Minute,        // per duration
		// an oversimplified example of rate limiting by a custom header
		httprate.WithKeyFuncs(func(r *http.Request) (string, error) {

			token := r.Context().Value("auth").(core.Authorization).AppId
			return token, nil
		}),
	)
}

func addAdminEndpoints(s *web.Service, jwtAuth func(http.Handler) http.Handler) {
	s.Route("/v1/admin", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtAuth, nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "Bearer", "", ""))
			r.Use(rateLimitByAppId(50))
			// r.Method(http.MethodGet, "/auditlogsummary", nethttp.NewHandler(GetAuditLogSummarys()))
			// r.Method(http.MethodGet, "/auditlogs/date/{date}/{hour}", nethttp.NewHandler(getAuditLogs()))
			// r.Method(http.MethodGet, "/auditlogs/powershell/{objectId}", nethttp.NewHandler(getAuditLogPowershell()))
			// r.Method(http.MethodPost, "/sharepoint/copylibrary", nethttp.NewHandler(copyLibrary()))
			// r.Method(http.MethodPost, "/sharepoint/copypage", nethttp.NewHandler(copyPage()))
			// r.Method(http.MethodPost, "/sharepoint/renamelibrary", nethttp.NewHandler(renameLibrary()))
			// r.Method(http.MethodGet, "/user/", nethttp.NewHandler(getUsers()))
			// r.Method(http.MethodPost, "/user/", nethttp.NewHandler(addUser()))
			// r.Method(http.MethodPatch, "/user/{upn}/credentials", nethttp.NewHandler(updateUserCredentials()))
			// r.MethodFunc(http.MethodPost, "/powershell", executePowerShell)

		})
	})

	s.Mount("/debug/admin", middleware.Profiler())
}

func Serve() {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Magicbox"
	s.OpenAPI.Info.WithDescription(description)
	s.OpenAPI.Info.Version = "v0.0.1"

	sharedSettings(s)

	addAdminEndpoints(s, Authenticator)
	s.Docs("/openapi/all", swgui.New)
	log.Println("Server starting")
	if err := http.ListenAndServe(":4300", s); err != nil {
		log.Fatal(err)
	}
}
