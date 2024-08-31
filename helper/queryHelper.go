package helper

import (
	"database/sql"
)

func AddLingkunganOrWilayahQueryHelper(idWilayah string, idLingkungan string, sqlScript string, tx *sql.Tx) (*sql.Rows, error) {
	var params []interface{}

	if idWilayah != "" {
		sqlScript += " AND id_wilayah = ?"
		params = append(params, idWilayah)
	}
	if idLingkungan != "" {
		sqlScript += " AND id_lingkungan = ?"
		params = append(params, idLingkungan)
	}

	return tx.Query(sqlScript, params...)

}