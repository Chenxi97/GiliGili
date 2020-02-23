package session

import (
	"log"
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
	//首先查看是否已经存在
	id := ""
	sessionMap.Range(func(k, v interface{}) bool {
		ss := v.(*defs.Session)
		if ss.LoginName == un {
			//删除过期session
			ct := nowInMilli()
			if ttl, err := strconv.ParseInt(ss.TTL, 10, 64); ttl < ct || err != nil {
				log.Print(ttl, "<", ct)
				deleteExpiredSession(id)
				return true
			}
			id = ss.ID
			return false
		}
		return true
	})
	if len(id) != 0 {
		return id
	}
	//不然就新建一个session
	id, _ = utils.NewUUID()
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
	log.Print("r: ", r)
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.Session)
		ct := nowInMilli()
		if ttl, err := strconv.ParseInt(ss.TTL, 10, 64); ttl < ct || err != nil {
			log.Print(ttl, "<", ct, " ", k.(string))
			deleteExpiredSession(ss.ID)
		} else {
			sessionMap.Store(k, ss)
		}
		return true
	})
	log.Print("sessionMap: ", sessionMap)
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	//log.Print("IsSessionExpired:",ss,ok)
	if ok {
		ct := nowInMilli()
		if ttl, err := strconv.ParseInt(ss.(*defs.Session).TTL, 10, 64); ttl < ct || err != nil {
			log.Print(ttl, "<", ct)
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.Session).LoginName, false
	}
	return "", true
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
