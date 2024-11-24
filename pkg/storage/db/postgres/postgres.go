package postgres

import (
	"database/sql"
	logger "effective_mobile/logs"
	"effective_mobile/pkg/storage/db"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Connection PostgresConn

type PostgresConn struct {
	conn    *sql.DB
	loggers logger.LogInterface
	migrate *migrate.Migrate
}

func InitConn(port int, host, user, password, dbname string, loggers logger.LogInterface) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataBase, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		loggers.Fatal(err.Error())
	}

	err = dataBase.Ping()
	if err != nil {
		loggers.Fatal(err.Error())
	}

	Connection.conn = dataBase
	Connection.loggers = loggers
	Connection.migrateUp()
	Connection.loggers.Info("Start db success")
}

func (p *PostgresConn) Close() {
	Connection.MigrateDown()
	err := p.conn.Close()
	if err != nil {
		p.loggers.Fatal(err.Error())
	}
}

func (p *PostgresConn) GetGroup(name string) (db.Group, error) {
	var group db.Group
	row := p.conn.QueryRow(`SELECT * FROM groups WHERE groups.name = $1`, name)

	err := row.Scan(&group.Id, &group.Name)
	return group, err
}

func (p *PostgresConn) GetGroupById(id int) (db.Group, error) {
	var group db.Group

	row := p.conn.QueryRow(`SELECT * FROM groups WHERE groups.id = $1`, id)

	err := row.Scan(&group.Id, &group.Name)

	return group, err
}

func (p *PostgresConn) DeleteSong(id int) error {
	_, err := p.conn.Exec(`DELETE FROM songs WHERE id = $1`, id)
	return err
}

func (p *PostgresConn) CreateSong(releaseDate, text, link, song, group string) error {
	gr, err := p.GetGroup(group)
	if err != nil {
		return err
	}
	_, err = p.conn.Exec(`INSERT INTO songs (release_date, text, link, song, group_id) VALUES ($1, $2, $3, $4, $5);`, releaseDate, text, link, song, gr.Id)
	return err
}

func (p *PostgresConn) GetSongs(pack int, filters db.SongFilter) ([]db.Song, error) {
	//мб потом добавить настраиваемый размер пагинации?
	offSet := 5 * (pack - 1)
	args := []interface{}{}
	argIndex := 1

	query := `SELECT id, release_date, text, link, group_id, song FROM songs WHERE 1=1`

	if filters.Group != "" {
		group, err := p.GetGroup(filters.Group)
		if err != nil {
			return nil, err
		}

		query += fmt.Sprintf(" AND group_id = $%d", argIndex)
		args = append(args, group.Id)
		argIndex++
	}

	if filters.ReleaseDate[0] != "" && filters.ReleaseDate[1] != "" {
		query += fmt.Sprintf(" AND release_date >= $%d AND release_date <= $%d", argIndex, argIndex+1)
		args = append(args, filters.ReleaseDate[0])
		args = append(args, filters.ReleaseDate[1])
		argIndex += 2
	}

	if filters.Song != "" {
		query += fmt.Sprintf(" AND song ILIKE $%d", argIndex)
		args = append(args, filters.Song)
		argIndex++
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, 5, offSet)

	rows, err := p.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []db.Song
	for rows.Next() {
		var groupId int

		var song db.Song
		if err = rows.Scan(&song.Id, &song.ReleaseDate, &song.Words, &song.Link, &groupId, &song.Song); err != nil {
			return nil, err
		}

		group, err := p.GetGroupById(groupId)
		if err != nil {
			return nil, err
		}

		song.Group = group

		songs = append(songs, song)
	}

	return songs, nil
}

func (p *PostgresConn) GetSong(id int) (db.Song, error) {
	var song db.Song
	var groupId int

	row := p.conn.QueryRow(`SELECT * FROM songs WHERE songs.id = $1`, id)

	err := row.Scan(&song.Id, &song.Song, &song.Words, &song.Link, &song.ReleaseDate, &groupId)

	gr, err := p.GetGroupById(groupId)
	if err != nil {
		return db.Song{}, err
	}

	song.Group = gr
	return song, err
}

func (p *PostgresConn) UpdateSong(id int, group, releaseDate, text, link string) error {
	groupObj, err := p.GetGroup(group)
	if err != nil {
		return err
	}

	_, err = p.conn.Exec(`UPDATE songs SET group_id = $1, release_date = $2, text = $3, link = $4 WHERE id = $5;`, groupObj.Id, releaseDate, text, link, id)
	return err
}

func (p *PostgresConn) migrateUp() {
	driver, err := postgres.WithInstance(p.conn, &postgres.Config{})
	if err != nil {
		p.loggers.Fatal(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/storage/db/postgres/migrations",
		"songs",
		driver,
	)
	p.migrate = m
	if err != nil {
		p.loggers.Fatal(err.Error())
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		p.loggers.Fatal(err.Error())
	}
	p.loggers.Info("Success migrations up")
}

func (p *PostgresConn) MigrateDown() {
	if err := p.migrate.Down(); err != nil && err != migrate.ErrNoChange {
		p.loggers.Fatal(err.Error())
	}

	p.loggers.Info("Success migrations down")
}
