package session

import (
	"log"
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(key, value interface{}) bool {
		session := value.(*defs.Session)
		sessionMap.Store(key, session)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano() / 100000
	ttl := ct + 30*60*1000
	session := &defs.Session{
		Username: un,
		TTL:      ttl,
	}

	sessionMap.Store(id, session)
	err := dbops.InserSession(id, ttl, un)
	if err != nil {
		log.Printf("%s\n", err)
		return ""
	}
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	session, ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano() / 100000
		if session.(*defs.Session).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return session.(*defs.Session).Username, false
	}
	return "", true
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	err := dbops.DeleteSession(sid)
	if err != nil {
		log.Printf("%s", err)
		return
	}
}