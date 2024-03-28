package repository

import (
	"database/sql"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type videoRepositoryDB struct {
	db *sql.DB
}

func NewRepositoryDB(db *sql.DB) videoRepositoryDB {
	return videoRepositoryDB{db: db}
}

func (v videoRepositoryDB) GetVideosId(video Video) ([]int, error) {
	videos := make([]int, 0)
	query := `SELECT Video.vid_id FROM Video 
	INNER JOIN Program ON Video.prog_id = Program.prog_id 
	INNER JOIN Department ON Program.dept_id = Department.dept_id 
	INNER JOIN Faculty ON Department.fac_id = Faculty.fac_id 
	INNER JOIN University ON Faculty.uni_id = University.uni_id 
	WHERE University.uni = ? AND Video.grad_year = ? AND Faculty.fac = ? AND Department.dept = ? AND Program.prog = ?`
	rows, err := v.db.Query(query, video.Uni, video.Year, video.Faculty, video.Department, video.Curriculum)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var vidId int
			err := rows.Scan(&vidId)
			if err != nil {
				return nil, err
			} else {
				videos = append(videos, vidId)
			}
		}
	}

	return videos, nil
}

func (v videoRepositoryDB) GetImage(id int) (string, error) {
	var path string
	query := "SELECT image_path FROM Video WHERE vid_id = ?"
	err := v.db.QueryRow(query, id).Scan(&path)
	if err != nil {
		return "error: ", err
	}
	baseDir := "/Users/nnp/Documents/GitHub/u-gmo-backend/"
	fullPath := filepath.Join(baseDir, path)
	// fmt.Println(fullPath)
	return fullPath, nil
}

func (v videoRepositoryDB) GetVideo(id int) (string, error) {
	var path string
	query := "SELECT vid_path FROM Video WHERE vid_id = ?"
	err := v.db.QueryRow(query, id).Scan(&path)
	if err != nil {
		return "error: ", err
	}
	baseDir := "/Users/nnp/Documents/GitHub/u-gmo-backend/"
	fullPath := filepath.Join(baseDir, path)
	// fmt.Println(fullPath)
	return fullPath, nil
}
