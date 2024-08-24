package strategy

import "fmt"

type PasswordProtector struct {
	User      string
	Password  string
	Algorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(password string)
}

func NewPasswordProtector(user, password string, algorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		User:      user,
		Password:  password,
		Algorithm: algorithm,
	}
}

func (p *PasswordProtector) SetAlgorithm(algorithm HashAlgorithm) {
	p.Algorithm = algorithm
}

// Global hash method, applies the hash algorithm to the password
func (p *PasswordProtector) Hash() {
	// Delegates the hash to the algorithm
	p.Algorithm.Hash(p.Password)
}

type SHA struct{}

func (s *SHA) Hash(password string) {
	fmt.Printf("Hashing using SHA the password %s\n", password)
}

type MD5 struct{}

func (m *MD5) Hash(password string) {
	fmt.Printf("Hashing using MD5 the password %s\n", password)
}
