package dbops

func ReadVideoDeletionRecord(count int) ([]string, error) {
	var res []string
	rows, err := dbConn.Limit(count).Find(&VideoDel{}).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return res, err
		}
		res = append(res, id)
	}

	return res, nil
}

func DelVideoDeletionRecord(vid string) error {
	err := dbConn.Where("id = ?", vid).Delete(&VideoDel{}).Error
	if err != nil {
		return err
	}
	return nil
}
