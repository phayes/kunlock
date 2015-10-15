package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS != "darwin" {
		log.Fatal("Only works with MacOS X currently. Sorry.")
	}
	_, err := os.Stat("/Volumes/KOBOeReader")
	if err != nil {
		log.Fatal("Could not find plugged-in KOBO eReader. Select your language, then select \"Don't have WiFi network?\", then plug your kobo into your Mac, then run kunlock.")
	}

	_, err = os.Stat("/Volumes/KOBOeReader/.kobo/KoboReader.sqlite")
	if err != nil {
		log.Fatal("Oops something went wrong. Perhaps this is a new type of KOBO that this program doesn't know how to work with. Error: ", err)
	}

	db, err := sql.Open("sqlite3", "/Volumes/KOBOeReader/.kobo/KoboReader.sqlite")
	if err != nil {
		log.Fatal("Failed to open KOBO device database:", err)
	}

	stmt, err := db.Prepare("INSERT INTO user(UserID, UserKey, UserDisplayName, UserEmail) values(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec("KUnLock", "KUnLock", "KUnLock@example.com", "KUnLock@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Kobo is unlocked. Unplug and enjoy.")
}
