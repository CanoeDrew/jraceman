package query

import (
  "net/http"

  "github.com/jimmc/jraceman/domain"
)

type entryQuery struct{
  h *handler
}

func (sc *entryQuery) EntityTypeName() string {
  return "entry"
}

func (sc *entryQuery) NewEntity() interface{} {
  return &domain.Entry{}
}

func (sc *entryQuery) SummaryQuery(format string) string {
  return "select ID, '[' || ID || '] ' as summary from " + sc.EntityTypeName()
}

func (h *handler) entry(w http.ResponseWriter, r *http.Request) {
  sq := &entryQuery{h}
  h.stdquery(w, r, sq)
}
