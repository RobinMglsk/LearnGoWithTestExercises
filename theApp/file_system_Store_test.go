package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T){

	t.Run("league from a reader", func(t *testing.T){
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)
		
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T){
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store win for existing players", func(t *testing.T){
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		player := "Cleo"
		store.RecordWin(player)

		got := store.GetPlayerScore(player)
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("store win for new players", func(t *testing.T){
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		
		player := "Pepper"
		store.RecordWin(player)
		
		got := store.GetPlayerScore(player)
		want := 1
		
		assertScoreEquals(t, got, want)
	})
	
	t.Run("works with an empty file", func(t *testing.T){
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()
		
		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})
	
	t.Run("league sorted", func(t *testing.T){
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		
		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)
		
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t testing.TB, err error){
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}