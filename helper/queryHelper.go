package helper

import (
	"database/sql"
)

func AddLingkunganOrWilayahQueryHelper(idWilayah string, idLingkungan string, sqlScript string, tx *sql.Tx) (*sql.Rows, error) {
	var params []interface{}

	if idWilayah != "" {
		sqlScript += " AND b.id_wilayah = ?"
		params = append(params, idWilayah)
	}
	if idLingkungan != "" {
		sqlScript += " AND c.id_lingkungan = ?"
		params = append(params, idLingkungan)
	}

	return tx.Query(sqlScript, params...)

}

// function buat convert slice ids ke interface untuk mengisi parameter query
func ConvertToInterfaceSlice(ids []int32) []interface{} {
	ifaces := make([]interface{}, len(ids))
	for i, v := range ids {
		ifaces[i] = v
	}
	return ifaces
}