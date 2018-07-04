package query

import (
  "net/http"

  "github.com/jimmc/jracemango/domain"
)

type seedingplanQuery struct{
  h *handler
}

func (sc *seedingplanQuery) EntityTypeName() string {
  return "seedingplan"
}

func (sc *seedingplanQuery) NewEntity() interface{} {
  return &domain.SeedingPlan{}
}

func (h *handler) seedingplan(w http.ResponseWriter, r *http.Request) {
  sq := &seedingplanQuery{h}
  h.stdquery(w, r, sq)
}