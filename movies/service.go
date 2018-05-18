package movies

import (
	"context"

	"database/sql"
	"github.com/rs/zerolog"
	"errors"
)

type Service interface {
	GetMovies(ctx context.Context) ([]Movie, error)
	GetMovieById(ctx context.Context, id string) (Movie, error)
	NewMovie(ctx context.Context, title string, director []string, year string, userid string) (string, error)
}

//implementation with database and logger
type moviesService struct {
	db     *sql.DB
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
func (m moviesService) GetMovies(ctx context.Context) ([]Movie, error) {
	rows, err := m.db.Query("select * from movies")
	if err != nil {
		return nil, err
	}
	movies := make([]Movie, 0)
	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Userid, &movie.CreatedOn, &movie.UpdatedOn)
		if err != nil {
			return nil, err
		}
		r, err := m.db.Query("select director from movie_directors where movie_id=$1", movie.Id)
		var director []string
		for r.Next()	{
			var d string
			err = r.Scan(&d)
			if err != nil {
				return nil, err
			}
			director = append(director, d)
		}
		movie.Director = director
		movies = append(movies, *movie)
	}
	return movies, nil
}

//implementation
func (m moviesService) GetMovieById(ctx context.Context, id string) (Movie, error) {
	var movie Movie
	rows := m.db.QueryRow("select * from movies where id = $1", id)
	err := rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Userid, &movie.CreatedOn, &movie.UpdatedOn)
	if err != nil {
		return movie, err
	}
	r, err := m.db.Query("select director from movie_directors where movie_id=$1", movie.Id)
	var director []string
	for r.Next()	{
		var d string
		err = r.Scan(&d)
		if err != nil {
			return movie, err
		}
		director = append(director, d)
	}
	movie.Director = director
	return movie, nil
}

//implementation
func (m moviesService) NewMovie(ctx context.Context, title string, director []string, year string, userid string) (string, error) {
	rows, err := m.db.Query("select * from movies where title='" + title + "'")
	if err != nil {
		//todo: add logging
		return "", err
	}
	if !rows.Next() {
		var id string
		err := m.db.QueryRow("insert into movies (title, year, userid) values($1,$2,$3) returning id", title, year, userid).Scan(&id)
		//res, err := stmt.Exec(movie.Title,movie.Director, movie.Year, movie.Userid)
		//id, err := res.LastInsertId()
		if err != nil {
			return "", err
		}
		for _, d := range director {
			_, err = m.db.Query("insert into movie_directors (movie_id, director) values($1,$2)", id, d)
			if err != nil {
				//rollback
				err1 := err
				_, err := m.db.Query("delete from movies where id=$1", id)
				if err != nil {
					return "", err
				}
				return "", err1
			}
		}
		//return strconv.FormatInt(id, 10), nil
		return id, nil
	} else {

		return "", errors.New("movie already exists")
	}
}
