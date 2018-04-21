package crud

import (
  "fmt"
  "net/http"

  "github.com/jimmc/jracemango/domain"
)

type handler struct {
  config *Config
}

type Config struct {
  Prefix string
  DomainRepos domain.Repos
}

func NewHandler(c *Config) http.Handler {
  h := handler{config: c}
  mux := http.NewServeMux()
  mux.HandleFunc(h.crudPrefix("site"), h.site)
  return mux
}

func (h *handler) crudPrefix(s string) string {
  return fmt.Sprintf("%s%s/", h.config.Prefix, s)
}
