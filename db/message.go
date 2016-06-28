package db

func (r *Repo) MessageById(dest interface{}, id string) error {
	return r.DB.QueryRow("SELECT id, name, body from message WHERE id = ? LIMIT 1").Scan(&dest)
}

func (r *Repo) MessageRecentLimit(dest interface{}, limit int) error {
	return r.DB.QueryRow("SELECT id, name, body from message ORDER BY id DESC LIMIT ?").Scan(&dest)
}

func (r *Repo) MessageRecent(dest interface{}) error {
	return r.MessageRecentLimit(&dest, 10)
}
