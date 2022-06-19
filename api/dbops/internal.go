package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InserSession(sid string, ttl int64, uname string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlStr, uname)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetrieveSeesion(sid string) (*defs.Session, error) {
	session := &defs.Session{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM seesions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}

	var ttl, uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil || err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		session.TTL = res
		session.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return session, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}

	for rows.Next() {
		var sid, ttlStr, login_name string
		if err = rows.Scan(&sid, &ttlStr, &login_name); err != nil {
			log.Printf("retrive sessions error: %s\n", err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlStr, 10, 64); err == nil {
			session := &defs.Session{
				Username: login_name,
				TTL:      ttl,
			}
			m.Store(sid, session)
			log.Printf(" session id: %s, ttl: %d", sid, session.TTL)
		}
	}

	defer stmtOut.Close()
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	if _, err = stmtOut.Query(sid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}
