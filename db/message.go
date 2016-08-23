package db

func (r *Repo) MessageById(dest *Message, id string) error {
	return r.DB.QueryRow("SELECT id, name, body from message WHERE id = ? LIMIT 1", id).Scan(&dest.Id, &dest.Name, &dest.Body)
}

func (r *Repo) RecentMessagesLimit(limit int) ([]Message, error) {
	var result []Message
	rows, err := r.DB.Query("SELECT id, name, body from message ORDER BY id DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item Message
		err = rows.Scan(&item.Id, &item.Name, &item.Body)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil

}

func (r *Repo) RecentMessages() ([]Message, error) {
	return r.RecentMessagesLimit(10)
}

func (r *Repo) CountMessage() (int, error) {
	var id string // for memo
	var count int
	err := r.DB.QueryRow("SELECT id, count(*) from message").Scan(&id, &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repo) PostMessage(name string, body string) error {
	count, err := r.CountMessage()
	if err != nil {
		return err
	}
	query := "INSERT INTO message(id, name, body) values(?, ?, ?)"
	_, err = r.DB.Exec(query, count+1, name, body)
	return err
}
