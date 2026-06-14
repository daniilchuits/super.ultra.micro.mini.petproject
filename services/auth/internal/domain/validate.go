package domain

func Validate(cred Credentials) error {

	if len(cred.Login) < 6 {
		return ErrShortLogin
	} else if len(cred.Password) < 6 {
		return ErrShortPassword
	} else {
		return nil
	}
}
