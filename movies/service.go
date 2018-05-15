package movies

import (
	"context"

	"database/sql"
	"github.com/rs/zerolog"
)

type Service interface {
	GetMovies(ctx context.Context)([]Movie, error)
	GetMovieById(ctx context.Context, id string) (Movie, error)
}

//implementation with database and logger
type moviesService struct {
	db *sql.DB
	logger zerolog.Logger
}

//constructor - we can later add initialization if needed
func NewService(db *sql.DB, logger zerolog.Logger) Service {
	return moviesService{
		db,
		logger,
	}
}

//implementation
func (m moviesService) GetMovies (ctx context.Context) ([]Movie, error) {
	rows, err := m.db.Query("select * from movies")
	if err != nil {
		return nil, err
	}
	movies := make([]Movie, 0)
	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Director, &movie.Year, &movie.Userid, &movie.CreatedOn, &movie.UpdatedOn)
		if err != nil {
			return nil, err
		}
		movies = append(movies, *movie)
	}
	return movies, nil
}

//implementation
func (m moviesService) GetMovieById (ctx context.Context, id string)(Movie, error) {
	var movie Movie
	rows := m.db.QueryRow("select * from movies where id = $1", id)
	err := rows.Scan(&movie.Id, &movie.Title, &movie.Director, &movie.Year, &movie.Userid, &movie.CreatedOn, &movie.UpdatedOn)
	if err != nil {
		return movie, err
	}
	return movie, nil
}
