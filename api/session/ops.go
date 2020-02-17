package session

import (
	"strconv"
	"sync"
	"time"

	"github.com/Chenxi97/GiliGili/api/dbops"
	"github.com/Chenxi97/GiliGili/api/defs"
	"github.com/Chenxi97/GiliGili/api/utils"
)

//sync.Map有利于读，不利于写
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 // Severside session valid time: 30 min
	ttlstr := strconv.FormatInt(ttl, 10)

	ss := &defs.Session{id, ttlstr, un}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttlstr, un)

	return id
}
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.Session)
		sessionMap.Store(k, ss)
		return true
	})
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ttl, err := strconv.ParseInt(ss.(*defs.Session).TTL, 10, 64); err == nil {
			if ttl < ct {
				deleteExpiredSession(sid)
				return "", true
			}
			return ss.(*defs.Session).LoginName, false
		}
	}
	return "", true
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
