package database

import "database/sql"

func Init(init bool) (*sql.DB, error) {

		db, err := ConnectDb()
		if err != nil {
			return nil, err
		}

		if init {

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
	
			err = DbCreateTestUser(db)
			if err != nil {
				return nil, err
			}
	
			err = DbCreateLocationTable(db)
			if err != nil {
				return nil, err
			}
			
		}


		return db, nil


}