package strategy

import "testing"

func TestStrategy(t *testing.T) {
	sha := &SHA{}
	md5 := &MD5{}

	protect1 := NewPasswordProtector("yo", "123456", sha)
	protect1.Hash()
	// Change the algorithm, maintaining the same behavior
	protect1.SetAlgorithm(md5)
	protect1.Hash()

	protect2 := NewPasswordProtector("tu", "345678", md5)
	protect2.Hash()
	protect2.SetAlgorithm(sha)
	protect2.Hash()
}
