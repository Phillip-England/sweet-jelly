package database

import (
	"database/sql"
)

func Init(resetDbEnv string) (*sql.DB, error) {

		db, err := ConnectDb()
		if err != nil {
			return nil, err
		}

		if resetDbEnv == "true" {

			err = DbDeleteAllTables(db)
			if err != nil {
				return nil, err
			}
		
			err = DbCreateUserTable(db)
			if err != nil {
				return nil, err
			}
	
			err = DbCreateSessionTable(db)
			if err != nil {
				return nil, err
			}
	
			err = DbCreateLocationTable(db)
			if err != nil {
				return nil, err
			}

			err = DbCreateTestLocations(db)
			if err != nil {
				return nil, err
			}

			err = DbCreateTestUsers(db)
			if err != nil {
				return nil, err
			}
	
			
		}

		return db, nil


}