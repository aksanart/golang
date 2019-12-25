package config

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) ([]byte, error) {
	password := []byte(pass)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil

}

// nil means it is a match
func CheckHash(key string, hashDB string) error {
	password := []byte(key)
	hashedPassword := []byte(hashDB)

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil { // nil means it is a match
		return err
	}
	return nil
}
