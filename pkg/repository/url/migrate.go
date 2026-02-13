package url

import (
	"fmt"
	"strings"
	"url-shortening-service/pkg/domain"
)

func (r *Repository) Migrate() {
	table := strings.TrimSpace(r.TableName)
	if table == "" {
		panic("TABLE_NAME is required")
	}

	if r.DB == nil {
		panic("DB connection is nil")
	}

	if r.DB.Dialector != nil && r.DB.Dialector.Name() == "postgres" {
		if err := r.fixPostgresUrlTable(table); err != nil {
			panic(err)
		}
	}

	if err := r.DB.AutoMigrate(&domain.ApiResponse{}); err != nil {
		panic(err)
	}

}

func (r *Repository) fixPostgresUrlTable(table string) error {
	var exists bool
	if err := r.DB.Raw(
		`SELECT EXISTS (
			SELECT 1 FROM information_schema.tables
			WHERE table_schema = current_schema() AND table_name = ?
		)`,
		table,
	).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		return nil
	}

	// Revisión de tipos actuales.
	getType := func(col string) (string, error) {
		var dataType string
		err := r.DB.Raw(
			`SELECT data_type FROM information_schema.columns
			WHERE table_schema = current_schema() AND table_name = ? AND column_name = ?`,
			table,
			col,
		).Scan(&dataType).Error
		return strings.ToLower(strings.TrimSpace(dataType)), err
	}

	idType, err := getType("id")
	if err != nil {
		return err
	}
	userIDType, err := getType("user_id")
	if err != nil {
		return err
	}

	needsFix := (idType == "bigint" || idType == "integer") || (userIDType == "bigint" || userIDType == "integer")
	if !needsFix {
		return nil
	}

	var count int64
	if err := r.DB.Table(table).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return r.DB.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, table)).Error
	}

	if err := r.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
		return err
	}

	// Fix id
	if idType == "bigint" || idType == "integer" {
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ADD COLUMN IF NOT EXISTS id_new char(36)`, table)).Error; err != nil {
			return err
		}
		if err := r.DB.Exec(fmt.Sprintf(`UPDATE "%s" SET id_new = gen_random_uuid()::text WHERE id_new IS NULL`, table)).Error; err != nil {
			return err
		}
	}

	// Fix user_id (intentando mapear desde users.id_old)
	if userIDType == "bigint" || userIDType == "integer" {
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ADD COLUMN IF NOT EXISTS user_id_new char(36)`, table)).Error; err != nil {
			return err
		}
		// Si existe tabla users con id_old, intentamos mapear; si no, generamos UUIDs para no romper inserts/reads.
		if err := r.DB.Exec(
			fmt.Sprintf(
				`UPDATE "%s" t SET user_id_new = u.id
				 FROM users u
				 WHERE t.user_id_new IS NULL AND u.id_old = t.user_id`,
				table,
			),
		).Error; err != nil {
			// Si falla (por ejemplo, no existe users.id_old), no abortamos todavía.
		}
		if err := r.DB.Exec(fmt.Sprintf(`UPDATE "%s" SET user_id_new = gen_random_uuid()::text WHERE user_id_new IS NULL`, table)).Error; err != nil {
			return err
		}
	}

	// Dropear PK y renombrar columnas.
	var pkName string
	if err := r.DB.Raw(
		`SELECT tc.constraint_name
		FROM information_schema.table_constraints tc
		WHERE tc.table_schema = current_schema() AND tc.table_name = ? AND tc.constraint_type = 'PRIMARY KEY'`,
		table,
	).Scan(&pkName).Error; err != nil {
		return err
	}
	if strings.TrimSpace(pkName) != "" {
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" DROP CONSTRAINT IF EXISTS "%s"`, table, pkName)).Error; err != nil {
			return err
		}
	}

	// Renombres si aplican
	if idType == "bigint" || idType == "integer" {
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" RENAME COLUMN id TO id_old`, table)).Error; err != nil {
			return err
		}
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" RENAME COLUMN id_new TO id`, table)).Error; err != nil {
			return err
		}
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ALTER COLUMN id SET NOT NULL`, table)).Error; err != nil {
			return err
		}
	}
	if userIDType == "bigint" || userIDType == "integer" {
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" RENAME COLUMN user_id TO user_id_old`, table)).Error; err != nil {
			return err
		}
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" RENAME COLUMN user_id_new TO user_id`, table)).Error; err != nil {
			return err
		}
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ALTER COLUMN user_id SET NOT NULL`, table)).Error; err != nil {
			return err
		}
	}

	// Restaurar PK
	if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ADD PRIMARY KEY (id)`, table)).Error; err != nil {
		return err
	}

	return nil
}
