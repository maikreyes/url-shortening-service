package user

import (
	"fmt"
	"log"
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

	log.Println("[USER_REPOSITORY] Migrating user table...")

	if r.DB.Dialector != nil && r.DB.Dialector.Name() == "postgres" {
		if err := r.fixPostgresUserTable(table); err != nil {
			panic(err)
		}
	}

	if err := r.DB.AutoMigrate(&domain.User{}); err != nil {
		panic(err)
	}

}

func (r *Repository) fixPostgresUserTable(table string) error {
	// Si la tabla no existe, AutoMigrate la crea correctamente.
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

	// Detectar tipo actual de la columna id.
	var dataType string
	if err := r.DB.Raw(
		`SELECT data_type FROM information_schema.columns
		WHERE table_schema = current_schema() AND table_name = ? AND column_name = 'id'`,
		table,
	).Scan(&dataType).Error; err != nil {
		return err
	}

	dataType = strings.ToLower(strings.TrimSpace(dataType))
	if dataType != "bigint" && dataType != "integer" {
		return nil
	}

	// Si está vacía, es más seguro recrearla.
	var count int64
	if err := r.DB.Table(table).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return r.DB.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, table)).Error
	}

	// Reparación in-place: crear ids UUID-text y preservar el id anterior en id_old.
	if err := r.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
		return err
	}

	if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ADD COLUMN IF NOT EXISTS id_new char(36)`, table)).Error; err != nil {
		return err
	}
	if err := r.DB.Exec(fmt.Sprintf(`UPDATE "%s" SET id_new = gen_random_uuid()::text WHERE id_new IS NULL`, table)).Error; err != nil {
		return err
	}

	// Dropear FKs que apunten a users.id para poder cambiar PK/columna.
	type fkRow struct {
		TableName      string
		ConstraintName string
	}
	var fks []fkRow
	if err := r.DB.Raw(
		`SELECT tc.table_name AS table_name, tc.constraint_name AS constraint_name
		FROM information_schema.table_constraints tc
		JOIN information_schema.constraint_column_usage ccu
		  ON ccu.constraint_name = tc.constraint_name AND ccu.table_schema = tc.table_schema
		WHERE tc.table_schema = current_schema()
		  AND tc.constraint_type = 'FOREIGN KEY'
		  AND ccu.table_name = ? AND ccu.column_name = 'id'`,
		table,
	).Scan(&fks).Error; err != nil {
		return err
	}
	for _, fk := range fks {
		if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" DROP CONSTRAINT IF EXISTS "%s"`, fk.TableName, fk.ConstraintName)).Error; err != nil {
			return err
		}
	}

	// Dropear PK actual.
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

	// Renombrar columnas: id -> id_old, id_new -> id, y restaurar PK.
	if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" RENAME COLUMN id TO id_old`, table)).Error; err != nil {
		return err
	}
	if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" RENAME COLUMN id_new TO id`, table)).Error; err != nil {
		return err
	}
	if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ALTER COLUMN id SET NOT NULL`, table)).Error; err != nil {
		return err
	}
	if err := r.DB.Exec(fmt.Sprintf(`ALTER TABLE "%s" ADD PRIMARY KEY (id)`, table)).Error; err != nil {
		return err
	}

	return nil
}
