package dbops

import (
	"log"
	"sync"

	"github.com/Chenxi97/GiliGili/api/defs"
)

func InsertSession(sid string, ttl string, uname string) error {
	session := defs.Session{ID: sid, TTL: ttl, LoginName: uname}
	err := dbConn.Create(&session).Error
	if err != nil {
		return err
	}
	return nil
}

func RetrieveSession(sid string) (*defs.Session, error) {
	session := defs.Session{}
	err := dbConn.Where("id = ?", sid).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	rows, err := dbConn.Find(&defs.Session{}).Rows()
	if err != nil {
		return m, err
	}
	for rows.Next() {
		var id, ttl, name string
		if er := rows.Scan(&id, &ttl, &name); er != nil {
			log.Printf("retrive sessions error: %s", er)
			break
		}
		m.Store(id, &defs.Session{ID: id, TTL: ttl, LoginName: name})
	}
	return m, nil
}

func DeleteSession(sid string) error {
	err := dbConn.Where("id = ?", sid).Delete(&defs.Session{}).Error
	if err != nil {
		return err
	}
	return nil
}
