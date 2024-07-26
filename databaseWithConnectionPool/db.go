package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main() {

	dbUri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_DBUSER"),
		os.Getenv("POSTGRES_DBPASS"),
		os.Getenv("POSTGRES_DBHOST"),
		os.Getenv("POSTGRES_DBPORT"),
		os.Getenv("POSTGRES_DBNAME_GO"),
	)
	connectionToDatabase(dbUri)
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("*********Enter the Artists Name********")
	scanner.Scan()
	albums := must(getAblumsByAritstName(scanner.Text()))
	fmt.Printf("Albums found are : %v \n", albums)

	fmt.Println("*********Enter the Artists Id********")
	scanner.Scan()
	album := must(getAblumsById(must(strconv.ParseInt(scanner.Text(), 10, 64))))
	fmt.Printf("Albums found are : %v \n", album)

	fmt.Println("*********Adding a new album********")
	albID := must(addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}))
	fmt.Printf("ID of added album: %v\n", albID)
	
}

func connectionToDatabase(dbUri string) {
	fmt.Println(dbUri)
	var err error
	db, err = pgxpool.Connect(context.Background(), dbUri)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected via pgxPool")
}

func must[T any](result T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getAblumsByAritstName(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query(context.Background(), "SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	//close the connection
	defer rows.Close();
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, album)
	}
	// any err in reading
	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

func getAblumsById(id int64) (Album, error) {
	var album Album
	row := db.QueryRow(context.Background(), "Select * from album where id = $1", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows{
			return album, fmt.Errorf("albumsByArtist %q: %v", id, err)
		}
		return album, fmt.Errorf("albumsByArtist %q: %v", id, err)
	}
	return album, nil
}

func addAlbum(alb Album) (int64, error) {
	var id int64
    err := db.QueryRow(
		context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)",
		alb.Title, alb.Artist, alb.Price).Scan(&id)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}