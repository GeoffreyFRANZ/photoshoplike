package session

import (
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

type Store struct {
	cookies *sessions.CookieStore
	data    sync.Map
}

func New(key []byte) *Store {
	return &Store{cookies: sessions.NewCookieStore(key)}
}
func (s *Store) Create(r *http.Request, w http.ResponseWriter, d PixelsData) (*sessions.Session, error) {
	godotenv.Load()
	sess, err := s.cookies.New(r, os.Getenv("SESSION_KEY"))
	if err != nil {
		return nil, err
	}
	s.data.Store(sess.ID, d)
	err = s.cookies.Save(r, w, sess)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
func (s *Store) Load(r *http.Request) (PixelsData, bool) {
	godotenv.Load()
	sess, err := s.cookies.Get(r, os.Getenv("SESSION_KEY"))
	if err != nil {
		return PixelsData{}, false
	}
	data, ok := s.data.Load(sess.ID)
	if ok == false {
		return PixelsData{}, false
	}
	return data.(PixelsData), ok
}
