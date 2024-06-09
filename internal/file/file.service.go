package file

type Service interface {
	Upload()
	Delete()
	Get()
}

type serviceImpl struct {
	repo Repository
	// logger *zap.Logger
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (r *serviceImpl) Upload() {}

func (r *serviceImpl) Delete() {}

func (r *serviceImpl) Get() {}
