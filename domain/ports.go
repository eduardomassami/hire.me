package domain

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE

type Repository interface {
	Get(alias string) ([]*URL, error)
	GetMostUsed() ([]*URL, error)
	Save(url *URL) error
}

type Service interface {
	Get(alias string) ([]*URL, error)
	GetMostUsed() ([]*URL, error)
	SaveNoCustomAlias(url *URL) error
	SaveWithCustomAlias(url *URL) error
}
