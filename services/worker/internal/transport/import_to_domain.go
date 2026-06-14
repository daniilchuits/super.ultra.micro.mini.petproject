package transport

import "worker/internal/domain"

func ImportToDomain(name Filename) domain.Filename {
	return domain.Filename{Name: name.Name}
}
