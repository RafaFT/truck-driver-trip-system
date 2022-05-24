package usecase

type DriverRepository interface {
	CreateDriverRepo
	DeleteDriverRepo
	GetDriverByCPFRepo
	GetDriverRepo
	UpdateDriverRepo
}
