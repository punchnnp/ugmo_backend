package repository

type Video struct {
	Uni        string `db:"uni"`
	Year       string `db:"grad_year"`
	Faculty    string `db:"fac"`
	Department string `db:"dept"`
	Curriculum string `db:"prog"`
}

type VideoRepository interface {
	GetVideosId(Video) ([]int, error)
	GetVideo(int) (string, error)
	GetImage(int) (string, error)
}
