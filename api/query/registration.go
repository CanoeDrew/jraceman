package query

import (
  "net/http"

  "github.com/jimmc/jraceman/domain"
)

type registrationQuery struct{
  h *handler
}

func (sc *registrationQuery) EntityTypeName() string {
  return "registration"
}

func (sc *registrationQuery) NewEntity() interface{} {
  return &domain.Registration{}
}

func (sc *registrationQuery) SummaryQuery(format string) string {
  return "select ID, '[' || ID || '] ' as summary from " + sc.EntityTypeName()
}

func (h *handler) registration(w http.ResponseWriter, r *http.Request) {
  sq := &registrationQuery{h}
  h.stdquery(w, r, sq)
}
