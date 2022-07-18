package migrations

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	initSchemaMigrationID = "SCHEMA_INIT"
)

// MigrateFunc is the func signature for migrating.
type MigrateFunc func(*gorm.DB) error

// RollbackFunc is the func signature for roll backing.
type RollbackFunc func(*gorm.DB) error

// InitSchemaFunc is the func signature for initializing the schema.
type InitSchemaFunc func(*gorm.DB) error

// Options define options for all migrations.
type Options struct {
	// TableName is the migrations table.
	TableName string
	// IDColumnName is the name of column where the migrations id will be stored.
	IDColumnName string
	// IDColumnSize is the length of the migrations id column
	IDColumnSize int
	// UseTransaction makes GorMigrate execute migrations inside a single transaction.
	// Keep in mind that not all databases support DDL commands inside transactions.
	UseTransaction bool
	// ValidateUnknownMigrations will cause migrate to fail if there's unknown migrations
	// IDs in the database
	ValidateUnknownMigrations bool
}

// Migration represents a database migrations (a modification to be made on the database).
type Migration struct {
	// Version is the migrations identifier. Usually a timestamp like "201601021504".
	Version string
	// Migrate is a function that will br executed while running this migrations.
	Migrate MigrateFunc
	// Rollback will be executed on rollback. Can be nil.
	Rollback RollbackFunc
}

// GorMigrate represents a collection of all migrations of a database schema.
type GorMigrate struct {
	db         *gorm.DB
	tx         *gorm.DB
	options    *Options
	migrations []*Migration
	initSchema InitSchemaFunc
}

// ReservedIDError is returned when a migrations is using a reserved ID
type ReservedIDError struct {
	ID string
}

func (e *ReservedIDError) Error() string {
	return fmt.Sprintf(`gormigrate: Reserved migrations Version: "%s"`, e.ID)
}

// DuplicatedIDError is returned when more than one migrations have the same ID
type DuplicatedIDError struct {
	ID string
}

func (e *DuplicatedIDError) Error() string {
	return fmt.Sprintf(`gormigrate: Duplicated migrations Version: "%s"`, e.ID)
}

var (
	// DefaultOptions can be used if you don't want to think about options.
	DefaultOptions = &Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}

	// ErrRollbackImpossible is returned when trying to rollback a migrations
	// that has no rollback function.
	ErrRollbackImpossible = errors.New("gormigrate: It's impossible to rollback this migrations")

	// ErrNoMigrationDefined is returned when no migrations is defined.
	ErrNoMigrationDefined = errors.New("gormigrate: No migrations defined")

	// ErrMissingID is returned when the Version od migrations is equal to ""
	ErrMissingID = errors.New("gormigrate: Missing Version in migrations")

	// ErrNoRunMigration is returned when any run migrations was found while
	// running RollbackLast
	ErrNoRunMigration = errors.New("gormigrate: Could not find last run migrations")

	// ErrMigrationIDDoesNotExist is returned when migrating or rolling back to a migrations Version that
	// does not exist in the list of migrations
	ErrMigrationIDDoesNotExist = errors.New("gormigrate: Tried to migrate to an Version that doesn't exist")

	// ErrUnknownPastMigration is returned if a migrations exists in the DB that doesn't exist in the code
	ErrUnknownPastMigration = errors.New("gormigrate: Found migrations in DB that does not exist in code")
)

// New returns a new GorMigrate.
func New(db *gorm.DB, options *Options, migrations []*Migration) *GorMigrate {
	if options.TableName == "" {
		options.TableName = DefaultOptions.TableName
	}
	if options.IDColumnName == "" {
		options.IDColumnName = DefaultOptions.IDColumnName
	}
	if options.IDColumnSize == 0 {
		options.IDColumnSize = DefaultOptions.IDColumnSize
	}
	return &GorMigrate{
		db:         db,
		options:    options,
		migrations: migrations,
	}
}

// InitSchema sets a function that is run if no migrations is found.
// The idea is preventing to run all migrations when a new clean database
// is being migrating. In this function you should create all tables and
// foreign key necessary to your application.
func (g *GorMigrate) InitSchema(initSchema InitSchemaFunc) {
	g.initSchema = initSchema
}

// Migrate executes all migrations that did not run yet.
func (g *GorMigrate) Migrate() error {
	if !g.hasMigrations() {
		return ErrNoMigrationDefined
	}
	var targetMigrationID string
	if len(g.migrations) > 0 {
		targetMigrationID = g.migrations[len(g.migrations)-1].Version
	}
	return g.migrate(targetMigrationID)
}

// MigrateTo executes all migrations that did not run yet up to the migrations that matches `migrationID`.
func (g *GorMigrate) MigrateTo(migrationID string) error {
	if err := g.checkIDExist(migrationID); err != nil {
		return err
	}
	return g.migrate(migrationID)
}

func (g *GorMigrate) migrate(migrationID string) error {
	if !g.hasMigrations() {
		return ErrNoMigrationDefined
	}

	if err := g.checkReservedID(); err != nil {
		return err
	}

	if err := g.checkDuplicatedID(); err != nil {
		return err
	}

	g.begin()
	defer g.rollback()

	if err := g.createMigrationTableIfNotExists(); err != nil {
		return err
	}

	if g.options.ValidateUnknownMigrations {
		unknownMigrations, err := g.unknownMigrationsHaveHappened()
		if err != nil {
			return err
		}
		if unknownMigrations {
			return ErrUnknownPastMigration
		}
	}

	if g.initSchema != nil {
		canInitializeSchema, err := g.canInitializeSchema()
		if err != nil {
			return err
		}
		if canInitializeSchema {
			if err := g.runInitSchema(); err != nil {
				return err
			}
			return g.commit()
		}
	}

	for _, migration := range g.migrations {
		if err := g.runMigration(migration); err != nil {
			return err
		}
		if migrationID != "" && migration.Version == migrationID {
			break
		}
	}
	return g.commit()
}

// There are migrations to apply if either there's a defined
// initSchema function or if the list of migrations is not empty.
func (g *GorMigrate) hasMigrations() bool {
	return g.initSchema != nil || len(g.migrations) > 0
}

// Check whether any migrations is using a reserved Version.
// For now there's only have one reserved Version, but there may be more in the future.
func (g *GorMigrate) checkReservedID() error {
	for _, m := range g.migrations {
		if m.Version == initSchemaMigrationID {
			return &ReservedIDError{ID: m.Version}
		}
	}
	return nil
}

func (g *GorMigrate) checkDuplicatedID() error {
	lookup := make(map[string]struct{}, len(g.migrations))
	for _, m := range g.migrations {
		if _, ok := lookup[m.Version]; ok {
			return &DuplicatedIDError{ID: m.Version}
		}
		lookup[m.Version] = struct{}{}
	}
	return nil
}

func (g *GorMigrate) checkIDExist(migrationID string) error {
	for _, migrate := range g.migrations {
		if migrate.Version == migrationID {
			return nil
		}
	}
	return ErrMigrationIDDoesNotExist
}

// RollbackLast undo the last migrations
func (g *GorMigrate) RollbackLast() error {
	if len(g.migrations) == 0 {
		return ErrNoMigrationDefined
	}

	g.begin()
	defer g.rollback()

	lastRunMigration, err := g.getLastRunMigration()
	if err != nil {
		return err
	}

	if err := g.rollbackMigration(lastRunMigration); err != nil {
		return err
	}
	return g.commit()
}

// RollbackTo undoes migrations up to the given migrations that matches the `migrationID`.
// Migration with the matching `migrationID` is not rolled back.
func (g *GorMigrate) RollbackTo(migrationID string) error {
	if len(g.migrations) == 0 {
		return ErrNoMigrationDefined
	}

	if err := g.checkIDExist(migrationID); err != nil {
		return err
	}

	g.begin()
	defer g.rollback()

	for i := len(g.migrations) - 1; i >= 0; i-- {
		migration := g.migrations[i]
		if migration.Version == migrationID {
			break
		}
		migrationRan, err := g.migrationRan(migration)
		if err != nil {
			return err
		}
		if migrationRan {
			if err := g.rollbackMigration(migration); err != nil {
				return err
			}
		}
	}
	return g.commit()
}

func (g *GorMigrate) getLastRunMigration() (*Migration, error) {
	for i := len(g.migrations) - 1; i >= 0; i-- {
		migration := g.migrations[i]

		migrationRan, err := g.migrationRan(migration)
		if err != nil {
			return nil, err
		}

		if migrationRan {
			return migration, nil
		}
	}
	return nil, ErrNoRunMigration
}

// RollbackMigration undo a migrations.
func (g *GorMigrate) RollbackMigration(m *Migration) error {
	g.begin()
	defer g.rollback()

	if err := g.rollbackMigration(m); err != nil {
		return err
	}
	return g.commit()
}

func (g *GorMigrate) rollbackMigration(m *Migration) error {
	if m.Rollback == nil {
		return ErrRollbackImpossible
	}

	if err := m.Rollback(g.tx); err != nil {
		return err
	}

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", g.options.TableName, g.options.IDColumnName)
	return g.tx.Exec(sql, m.Version).Error
}

func (g *GorMigrate) runInitSchema() error {
	if err := g.initSchema(g.tx); err != nil {
		return err
	}
	if err := g.insertMigration(initSchemaMigrationID); err != nil {
		return err
	}

	for _, migration := range g.migrations {
		if err := g.insertMigration(migration.Version); err != nil {
			return err
		}
	}

	return nil
}

func (g *GorMigrate) runMigration(migration *Migration) error {
	if len(migration.Version) == 0 {
		return ErrMissingID
	}

	migrationRan, err := g.migrationRan(migration)
	if err != nil {
		return err
	}
	if !migrationRan {
		if err := migration.Migrate(g.tx); err != nil {
			return err
		}

		if err := g.insertMigration(migration.Version); err != nil {
			return err
		}
	}
	return nil
}

func (g *GorMigrate) createMigrationTableIfNotExists() error {
	if g.tx.Migrator().HasTable(g.options.TableName) {
		return nil
	}

	sql := fmt.Sprintf("CREATE TABLE %s (%s VARCHAR(%d) PRIMARY KEY)", g.options.TableName, g.options.IDColumnName, g.options.IDColumnSize)
	return g.tx.Exec(sql).Error
}

func (g *GorMigrate) migrationRan(m *Migration) (bool, error) {
	var count int64
	err := g.tx.
		Table(g.options.TableName).
		Where(fmt.Sprintf("%s = ?", g.options.IDColumnName), m.Version).
		Count(&count).
		Error
	return count > 0, err
}

// The schema can be initialised only if it hasn't been initialised yet
// and no other migrations has been applied already.
func (g *GorMigrate) canInitializeSchema() (bool, error) {
	migrationRan, err := g.migrationRan(&Migration{Version: initSchemaMigrationID})
	if err != nil {
		return false, err
	}
	if migrationRan {
		return false, nil
	}

	// If the Version doesn't exist, we also want the list of migrations to be empty
	var count int64
	err = g.tx.
		Table(g.options.TableName).
		Count(&count).
		Error
	return count == 0, err
}

func (g *GorMigrate) unknownMigrationsHaveHappened() (bool, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s", g.options.IDColumnName, g.options.TableName)
	rows, err := g.tx.Raw(sql).Rows()
	if err != nil {
		return false, err
	}
	defer rows.Close()

	validIDSet := make(map[string]struct{}, len(g.migrations)+1)
	validIDSet[initSchemaMigrationID] = struct{}{}
	for _, migration := range g.migrations {
		validIDSet[migration.Version] = struct{}{}
	}

	for rows.Next() {
		var pastMigrationID string
		if err := rows.Scan(&pastMigrationID); err != nil {
			return false, err
		}
		if _, ok := validIDSet[pastMigrationID]; !ok {
			return true, nil
		}
	}

	return false, nil
}

func (g *GorMigrate) insertMigration(id string) error {
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?)", g.options.TableName, g.options.IDColumnName)
	return g.tx.Exec(sql, id).Error
}

func (g *GorMigrate) begin() {
	if g.options.UseTransaction {
		g.tx = g.db.Begin()
	} else {
		g.tx = g.db
	}
}

func (g *GorMigrate) commit() error {
	if g.options.UseTransaction {
		return g.tx.Commit().Error
	}
	return nil
}

func (g *GorMigrate) rollback() {
	if g.options.UseTransaction {
		g.tx.Rollback()
	}
}
