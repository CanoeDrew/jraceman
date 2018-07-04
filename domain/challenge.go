package domain

// ChallengeRepo describes how Challenge records are loaded and saved.
type ChallengeRepo interface {
  FindByID(ID string) (*Challenge, error)
  List(offset, limit int) ([]*Challenge, error)
  Save(*Challenge) (string, error)
  UpdateByID(ID string, oldChallenge, newChallenge *Challenge, diffs Diffs) error
  DeleteByID(ID string) error
}

// Challenge describes a challenge group.
type Challenge struct {
  ID string
  Name string
}