package database

import (
  "fmt"
  "net/http"

  "github.com/jimmc/jraceman/domain"

  "github.com/jimmc/auth/auth"
  "github.com/jimmc/auth/permissions"
)

type handler struct {
  config *Config
}

const (
  editDatabase = permissions.Permission("edit_database")
)

// Config provides configuration of the http handler for our calls.
type Config struct {
  Prefix string
  DomainRepos domain.Repos
  AuthHandler *auth.Handler
}

// NewHandler creates the http handler for our calls.
func NewHandler(c *Config) http.Handler {
  h := handler{config: c}
  mux := http.NewServeMux()
  mux.HandleFunc(h.apiPrefix("upgrade"), c.AuthHandler.RequirePermissionFunc(h.upgrade, editDatabase))
  mux.HandleFunc(h.config.Prefix, h.blank)
  return mux
}

func (h *handler) blank(w http.ResponseWriter, r *http.Request) {
  // TODO - only list paths that the user has access to.
  http.Error(w, "Try /api/database/upgrade", http.StatusForbidden)
}

func (h *handler) apiPrefix(s string) string {
  return fmt.Sprintf("%s%s/", h.config.Prefix, s)
}
