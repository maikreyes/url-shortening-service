package user

import "golang.org/x/crypto/bcrypt"

func (s *Service) EncryptPassword(pwsd string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(pwsd),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil

}
