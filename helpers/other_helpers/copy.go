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

func CopyStructPhoto(source *domain.PhotoUpdateInput) *domain.Photo {
	// Meng-copy data dari struct source ke struct target
	target := &domain.Photo{
		Title:    source.Title,
		PhotoURL: source.PhotoURL,
		Caption:  source.Caption,
	}

	return target
}
