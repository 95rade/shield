package main

import (
	"database/sql"
	"os"

	"github.com/jhunt/go-log"

	"github.com/starkandwayne/shield/db"
)

func migrateRetention(to, from *db.DB) {
	n := 0
	rs, err := from.Query(`
	   SELECT uuid, name, summary, expiry
	     FROM retention`)
	if err != nil {
		log.Errorf("failed to migrate retention; SELECT said: %s", err)
		os.Exit(3)
	}

	for rs.Next() {
		n += 1
		var a, b, c, d sql.RawBytes

		if err := rs.Scan(&a, &b, &c, &d); err != nil {
			log.Errorf("failed to read result from retention source table: %s", err)
			os.Exit(3)
		}

		err = to.Exec(`
		   INSERT INTO retention (uuid, name, summary, expiry)
                VALUES (?, ?, ?, ?)`,
			a, b, c, d)
		if err != nil {
			log.Errorf("failed to insert result into retention dest table: %s", err)
			os.Exit(3)
		}
	}

	log.Infof("migrated %d retention", n)
}
