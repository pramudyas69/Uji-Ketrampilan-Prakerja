package other_helpers

import "uji/domain"

func CopyStructUser(source *domain.UserUpdateInput) *domain.User {
	// Meng-copy data dari struct source ke struct target
	target := &domain.User{
		Email:    source.Email,
		Age:      source.Age,
		Username: source.Username,
		Password: source.Password,
	}

	return target
}
