package db

func InitDB() {
	ConnectDB()
	MigrateDB()
}
