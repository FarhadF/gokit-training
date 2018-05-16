package movies

import (
	"context"

	"database/sql"
	"github.com/rs/zerolog"
	"errors"
)

type Service interface {
	GetMovies(ctx context.Context)([]Movie, error)
	GetMovieById(ctx context.Context, id string) (Movie, error)
	NewMovie(ctx context.Context, title string, director string, year string, userid string) (string, error)
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

//implementation
func (m moviesService) NewMovie (ctx context.Context, title string, director string, year string, userid string) (string, error) {
	rows, err := m.db.Query("select * from movies where title='" + title + "'")
	if err != nil {
		return "", err
	}
	if !rows.Next() {
		var id string
		err := m.db.QueryRow("insert into movies (title, director, year, userid) values($1,$2,$3,$4) returning id", title, director, year, userid).Scan(&id)
		//res, err := stmt.Exec(movie.Title,movie.Director, movie.Year, movie.Userid)
		//id, err := res.LastInsertId()
		if err != nil {
			return "", err
		}
		//return strconv.FormatInt(id, 10), nil
		return id, nil
	} else {

		return "", errors.New("movie already exists")
	}
}
