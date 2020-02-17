package dbops

//AddVideoDeletionRecord 将要删除的文件id写入数据库
func AddVideoDeletionRecord(vid string) error {
	video := VideoDel{ID: vid}
	err = dbConn.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}
