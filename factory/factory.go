package factory

func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &Connection.MySql{}
	case 2:
		&connection.Postgres{}
	default:
		return nil
	}
}
