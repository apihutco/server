package mysql

func Exec(sql string) error {
	return db.Exec(sql).Error
}
