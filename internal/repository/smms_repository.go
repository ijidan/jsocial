package repository

import "github.com/ijidan/jsocial/internal/pkg/smms"

type SmMsRepository struct {
	smMs *smms.SmMs
}

func (s *SmMsRepository) UploadImage(filePath string) (smms.Image, error) {
	return s.smMs.UploadImage(filePath)
}

func NewSmMsRepository(smMs *smms.SmMs) *SmMsRepository {
	r := &SmMsRepository{
		smMs: smMs,
	}
	return r
}
