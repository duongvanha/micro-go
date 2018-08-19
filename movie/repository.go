package movie

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository struct {
}

func (r *Repository) GetByPage(page int) ([5]Movie, error) {
	connStr := "**"
	var moviesTmp [5]Movie

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return moviesTmp, err
	}

	rows, e := db.Query("SELECT * FROM movies LIMIT 5")

	if e != nil {
		return moviesTmp, e
	}

	for i := 0; rows.Next(); i++ {
		var namevi string
		var nameen string
		var image string
		var status string
		var imdb string
		var director string
		var country string
		var year string
		var date string
		var time string
		var cam string
		var quality string
		var sub string
		var categories string
		var manufacturer string
		var view string
		var url string
		var movies string

		err = rows.Scan(
			&namevi,
			&nameen,
			&image,
			&status,
			&imdb,
			&director,
			&country,
			&year,
			&date,
			&time,
			&cam,
			&quality,
			&sub,
			&categories,
			&manufacturer,
			&view,
			&url,
			&movies,
		)

		moviesTmp[i] = Movie{NameVi: namevi, NameEn: nameen, Image: image, Status: status, IMDb: imdb, Director: director, Country: country, Year: year, Date: date, Time: time, Cam: cam, Quality: quality, Sub: sub, Categories: categories, Manufacturer: manufacturer, View: view, Url: url}

	}
	return moviesTmp, nil
}
