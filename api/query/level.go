package query

import (
  "net/http"

  "github.com/jimmc/jracemango/domain"
)

type levelQuery struct{
  h *handler
}

func (sc *levelQuery) EntityTypeName() string {
  return "level"
}

func (sc *levelQuery) NewEntity() interface{} {
  return &domain.Level{}
}

func (h *handler) level(w http.ResponseWriter, r *http.Request) {
  sq := &levelQuery{h}
  h.stdquery(w, r, sq)
}