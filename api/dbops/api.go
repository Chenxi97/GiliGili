package dbops

import (
	"time"

	"github.com/Chenxi97/GiliGili/api/defs"
	"github.com/Chenxi97/GiliGili/api/utils"
)

func AddUser(loginName string, pwd string) (*defs.User, error) {
	user := defs.User{LoginName: loginName, Pwd: pwd}
	err := dbConn.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(loginName string) (*defs.User, error) {
	user := defs.User{}
	err := dbConn.Where("login_name = ?", loginName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(loginName string, pwd string) error {
	err := dbConn.Where("login_name = ? AND pwd = ?", loginName, pwd).Delete(&defs.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func AddNewVideo(aid uint, name string) (*defs.VideoInfo, error) {
	//create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	ctime := time.Now().Format("Jan 02 2006, 15:04:05") //M D y, HH:MM:SS
	video := defs.VideoInfo{ID: vid, AuthorID: aid, Name: name, DisplayCtime: ctime}
	err = dbConn.Create(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	video := defs.VideoInfo{}
	err := dbConn.Where("id = ?", vid).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func DeleteVideoInfo(vid string) error {
	err := dbConn.Where("id = ?", vid).Delete(&defs.VideoInfo{}).Error
	if err != nil {
		return err
	}
	return nil
}

func AddNewComments(vid string, aid uint, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	comment := defs.Comment{ID: id, VideoID: vid, AuthorID: aid, Content: content}
	err = dbConn.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	res := []*defs.VideoInfo{}
	rows, err := dbConn.Table("comment").Order("video_info.create_time DESC").
		Select("video_info.id, video_info.author_id, video_info.name, video_info.display_ctime").
		Joins("left join user on video_info.author_id = users.id").
		Where("users.login_name=? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time<=FROM_UNIXTIME(?) ", uname, from, to).
		Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := &defs.VideoInfo{}
		if err := rows.Scan(&c); err != nil {
			return res, err
		}
		res = append(res, c)
	}

	return res, nil
}

func ListComments(vid string, from, to int) ([]*defs.CommentForList, error) {
	res := []*defs.CommentForList{}
	rows, err := dbConn.Table("comment").Order("comment.time DESC").Select("comment.id, user.login_name, comment.content").
		Joins("left join user on user.id = comment.author_id").
		Where("comment.video_id = ? AND comment.time > FROM_UNIXTIME(?) AND comment.time <= FROM_UNIXTIME(?) ", vid, from, to).
		Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.CommentForList{Comment: defs.Comment{ID: id, VideoID: vid, Content: content}, AuthorName: name}
		res = append(res, c)
	}

	return res, nil
}
