package crud

import (
  "net/http"

  "github.com/jimmc/jracemango/domain"
)

type levelCrud struct{
  h *handler
}

func (sc *levelCrud) EntityTypeName() string {
  return "level"
}

func (sc *levelCrud) NewEntity() interface{} {
  return &domain.Level{}
}

func (sc *levelCrud) Save(entity interface{}) error {
  var level *domain.Level = entity.(*domain.Level)
  return sc.h.config.DomainRepos.Level().Save(level)
}

func (sc *levelCrud) List(offset, limit int) ([]interface{}, error) {
  sites, err := sc.h.config.DomainRepos.Level().List(offset, limit)
  if err != nil {
    return nil, err
  }
  a := make([]interface{}, len(sites))
  for i, level := range sites {
    a[i] = level
  }
  return a, nil
}

func (sc *levelCrud) FindByID(ID string) (interface{}, error) {
  return sc.h.config.DomainRepos.Level().FindByID(ID)
}

func (sc *levelCrud) DeleteByID(ID string) error {
  return sc.h.config.DomainRepos.Level().DeleteByID(ID)
}

func (sc *levelCrud) UpdateByID(ID string, oldEntity, newEntity interface{}, diffs domain.Diffs) error {
  var oldLevel *domain.Level = oldEntity.(*domain.Level)
  var newLevel *domain.Level = newEntity.(*domain.Level)
  return sc.h.config.DomainRepos.Level().UpdateByID(ID, oldLevel, newLevel, diffs)
}

func (h *handler) level(w http.ResponseWriter, r *http.Request) {
  sc := &levelCrud{h}
  h.stdcrud(w, r, sc)
}