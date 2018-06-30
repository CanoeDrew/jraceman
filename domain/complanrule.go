package domain

// ComplanRuleRepo describes how ComplanRule records are loaded and saved.
type ComplanRuleRepo interface {
  FindByID(ID string) (*ComplanRule, error)
  List(offset, limit int) ([]*ComplanRule, error)
  Save(*ComplanRule) (string, error)
  UpdateByID(ID string, oldComplanRule, newComplanRule *ComplanRule, diffs Diffs) error
  DeleteByID(ID string) error
}

// ComplanRule describes a progression rule in a Complan.
type ComplanRule struct {
  ID string
  ComplanID string
  FromRound int
  FromSection int
  FromPlace int
  ToRound int
  ToSection int
  ToLane int
}
