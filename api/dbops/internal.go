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

// func RetrieveSession(name string) (*defs.Session, error) {
// 	session := defs.Session{}
// 	err := dbConn.Where("login_name = ?", name).First(&session).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &session, nil
// }

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	sessions := []defs.Session{}
	if err := dbConn.Find(&sessions).Error; err != nil {
		log.Printf("retrive sessions error: %s", err)
		return m, err
	}
	for _, v := range sessions {
		m.Store(v.ID, &v)
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
