package query

import (
  "fmt"
  "net/http"

  "github.com/jimmc/jracemango/domain"
)

type handler struct {
  config *Config
}

// Config provides configuration of the http handler for query calls.
type Config struct {
  Prefix string
  DomainRepos domain.Repos
}

// NewHandler creates the http handler for query calls.
func NewHandler(c *Config) http.Handler {
  h := handler{config: c}
  mux := http.NewServeMux()
  mux.HandleFunc(h.queryPrefix("area"), h.area)
  mux.HandleFunc(h.queryPrefix("challenge"), h.challenge)
  mux.HandleFunc(h.queryPrefix("competition"), h.competition)
  mux.HandleFunc(h.queryPrefix("complan"), h.complan)
  mux.HandleFunc(h.queryPrefix("complanrule"), h.complanrule)
  mux.HandleFunc(h.queryPrefix("complanstage"), h.complanstage)
  mux.HandleFunc(h.queryPrefix("exception"), h.exception)
  mux.HandleFunc(h.queryPrefix("gender"), h.gender)
  mux.HandleFunc(h.queryPrefix("laneorder"), h.laneorder)
  mux.HandleFunc(h.queryPrefix("level"), h.level)
  mux.HandleFunc(h.queryPrefix("meet"), h.meet)
  mux.HandleFunc(h.queryPrefix("person"), h.person)
  mux.HandleFunc(h.queryPrefix("progression"), h.progression)
  mux.HandleFunc(h.queryPrefix("registration"), h.registration)
  mux.HandleFunc(h.queryPrefix("registrationfee"), h.registrationfee)
  mux.HandleFunc(h.queryPrefix("scoringrule"), h.scoringrule)
  mux.HandleFunc(h.queryPrefix("scoringsystem"), h.scoringsystem)
  mux.HandleFunc(h.queryPrefix("seedinglist"), h.seedinglist)
  mux.HandleFunc(h.queryPrefix("seedingplan"), h.seedingplan)
  mux.HandleFunc(h.queryPrefix("simplan"), h.simplan)
  mux.HandleFunc(h.queryPrefix("simplanrule"), h.simplanrule)
  mux.HandleFunc(h.queryPrefix("simplanstage"), h.simplanstage)
  mux.HandleFunc(h.queryPrefix("site"), h.site)
  mux.HandleFunc(h.queryPrefix("stage"), h.stage)
  mux.HandleFunc(h.queryPrefix("team"), h.team)
  return mux
}

func (h *handler) queryPrefix(s string) string {
  return fmt.Sprintf("%s%s/", h.config.Prefix, s)
}
