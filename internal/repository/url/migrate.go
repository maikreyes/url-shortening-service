package url

func (r *Repository) Migrate() {

	r.DB.Exec(`CREATE TABLE IF NOT EXISTS urls (
	id int NOT NULL auto_increment PRIMARY KEY,   
	url VARCHAR(100) NOT NULL,
	short_code VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
        DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP
);`)

}
