package db

import "database/sql"

type PostgresRepository struct {
	db *sql.DB
}

type UserRepository struct {
	PostgresRepository
}

type SongRepository struct {
	PostgresRepository
}

type AlbumRepository struct {
	PostgresRepository
}

func NewUserRepository(url string) (*UserRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		PostgresRepository{db: db},
	}, nil
}

func NewSongRepository(url string) (*SongRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &SongRepository{
		PostgresRepository{db: db},
	}, nil
}

func NewAlbumRepository(url string) (*AlbumRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &AlbumRepository{
		PostgresRepository{db: db},
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}
