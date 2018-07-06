package crud

import (
  "fmt"
  "net/http"

  "github.com/jimmc/jracemango/domain"
)

type handler struct {
  config *Config
}

// Config provides configuration of the http handler for CRUD calls.
type Config struct {
  Prefix string
  DomainRepos domain.Repos
}

// NewHandler creates the http handler for CRUD calls.
func NewHandler(c *Config) http.Handler {
  h := handler{config: c}
  mux := http.NewServeMux()
  mux.HandleFunc(h.crudPrefix("area"), h.area)
  mux.HandleFunc(h.crudPrefix("challenge"), h.challenge)
  mux.HandleFunc(h.crudPrefix("competition"), h.competition)
  mux.HandleFunc(h.crudPrefix("complan"), h.complan)
  mux.HandleFunc(h.crudPrefix("complanrule"), h.complanrule)
  mux.HandleFunc(h.crudPrefix("complanstage"), h.complanstage)
  mux.HandleFunc(h.crudPrefix("entry"), h.entry)
  mux.HandleFunc(h.crudPrefix("event"), h.event)
  mux.HandleFunc(h.crudPrefix("exception"), h.exception)
  mux.HandleFunc(h.crudPrefix("gender"), h.gender)
  mux.HandleFunc(h.crudPrefix("laneorder"), h.laneorder)
  mux.HandleFunc(h.crudPrefix("level"), h.level)
  mux.HandleFunc(h.crudPrefix("meet"), h.meet)
  mux.HandleFunc(h.crudPrefix("person"), h.person)
  mux.HandleFunc(h.crudPrefix("progression"), h.progression)
  mux.HandleFunc(h.crudPrefix("registration"), h.registration)
  mux.HandleFunc(h.crudPrefix("registrationfee"), h.registrationfee)
  mux.HandleFunc(h.crudPrefix("scoringrule"), h.scoringrule)
  mux.HandleFunc(h.crudPrefix("scoringsystem"), h.scoringsystem)
  mux.HandleFunc(h.crudPrefix("seedinglist"), h.seedinglist)
  mux.HandleFunc(h.crudPrefix("seedingplan"), h.seedingplan)
  mux.HandleFunc(h.crudPrefix("simplan"), h.simplan)
  mux.HandleFunc(h.crudPrefix("simplanrule"), h.simplanrule)
  mux.HandleFunc(h.crudPrefix("simplanstage"), h.simplanstage)
  mux.HandleFunc(h.crudPrefix("site"), h.site)
  mux.HandleFunc(h.crudPrefix("stage"), h.stage)
  mux.HandleFunc(h.crudPrefix("team"), h.team)
  return mux
}

func (h *handler) crudPrefix(s string) string {
  return fmt.Sprintf("%s%s/", h.config.Prefix, s)
}
