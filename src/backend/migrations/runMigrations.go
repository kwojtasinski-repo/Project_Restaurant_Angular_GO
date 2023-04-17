package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
)

var currentMigrationVersion uint

func UpMigrations(config config.Config, migrationsToApply string) {
	log.Println("Prepare up migrations...")

	migrationsApply := 0
	if len(migrationsToApply) > 0 {
		var err error
		migrationsApply, err = strconv.Atoi(migrationsToApply)
		if err != nil {
			panic(err)
		}

		if migrationsApply <= 0 {
			panic("Invalid number cannot use negative or zero value")
		}
	}

	if err := createDatabaseIfNotExists(config); err != nil {
		panic(err)
	}

	if err := createDbUser(config); err != nil {
		panic(err)
	}

	migrate, err := createMigrationObject(config)
	if err != nil {
		panic(err)
	}
	defer migrate.Close()

	if err != nil {
		panic(err)
	}

	checkVersion(*migrate)
	log.Println("Migrating up database schema. Current version ", currentMigrationVersion)
	if migrationsApply == 0 {
		log.Println("Migrating up all versions")
		err = migrate.Up()
	} else {
		log.Println(getScriptMigrationText(migrationsApply))
		err = migrate.Steps(migrationsApply)
	}
	err = handleErrors(err)

	if err != nil {
		panic(err)
	}
}

func DownMigrations(config config.Config, migrationsToApply string) {
	log.Println("Prepare down migrations...")

	migrationsApply := 0
	if len(migrationsToApply) > 0 {
		var err error
		migrationsApply, err = strconv.Atoi(migrationsToApply)
		if err != nil {
			panic(err)
		}

		if migrationsApply >= 0 {
			panic("Invalid number cannot use positive or zero value")
		}
	}

	migrate, err := createMigrationObject(config)
	if err != nil {
		panic(err)
	}
	defer migrate.Close()

	checkVersion(*migrate)
	log.Println("Migrating down database schema. Current version ", currentMigrationVersion)
	if migrationsApply == 0 {
		log.Println("Migrating down all versions")
		err = migrate.Down()
	} else {
		log.Println(getScriptMigrationText(migrationsApply))
		err = migrate.Steps(migrationsApply)
	}
	err = handleErrors(err)

	if err != nil {
		panic(err)
	}
}

func createDatabaseIfNotExists(config config.Config) error {
	log.Println("CREATE DATABASE IF NOT EXISTS " + config.Database.Name)
	db, err := sql.Open("mysql", config.DatabaseMigration.Username+":"+config.DatabaseMigration.Password+"@tcp(localhost:3306)/")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.Database.Name)
	if err != nil {
		return err
	}

	return nil
}

func createDbUser(configFile config.Config) error {
	log.Println("CREATE USER IF NOT EXISTS ", configFile.Database.Username)
	db, err := sql.Open("mysql", configFile.DatabaseMigration.Username+":"+configFile.DatabaseMigration.Password+"@tcp(localhost:3306)/")
	if err != nil {
		return err
	}
	defer db.Close()

	queryCreateUser := "CREATE USER IF NOT EXISTS " + configFile.Database.Username + " IDENTIFIED BY '" + configFile.Database.Password + "';"
	log.Println("Running query", queryCreateUser)
	_, err = db.Exec(queryCreateUser)
	if err != nil {
		return err
	}

	log.Println("Added grants for user " + configFile.Database.Username)
	queryGrant := "GRANT select, update, insert, delete ON " + configFile.Database.Name + ".* to " + configFile.Database.Username + ";"
	log.Println("Running query", queryGrant)
	_, err = db.Exec(queryGrant)
	if err != nil {
		return err
	}

	return nil
}

func createMigrationObject(configFile config.Config) (*migrate.Migrate, error) {
	db, err := sql.Open("mysql", configFile.DatabaseMigration.Username+":"+configFile.DatabaseMigration.Password+"@tcp(localhost:3306)/"+configFile.Database.Name)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	path := filepath.Join(config.GetRootPath(), "migrations")
	splited := strings.Split(path, ":"+string(os.PathSeparator))
	pathSplited := splited[1]
	migrate, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///"+pathSplited),
		configFile.Database.Name, driver)

	if err != nil {
		return nil, err
	}

	return migrate, nil
}

func checkVersion(migration migrate.Migrate) {
	var isDirty bool
	var err error
	currentMigrationVersion, isDirty, err = migration.Version()
	if err == migrate.ErrNilVersion {
		log.Print("Migrations: There is no migrations applied")
		return
	}

	if err != nil {
		panic(err)
	}
	if isDirty {
		log.Print("Current migration ", currentMigrationVersion, " version is dirty check if database need clean up and change flag in schema_migrations to false")
		os.Exit(0)
	}
}

func handleErrors(err error) error {
	if err == migrate.ErrNoChange {
		log.Println("Migrations: ", err.Error())
		return nil
	}

	if err == migrate.ErrNilVersion {
		log.Println("Migrations: ", err.Error())
		return nil
	}

	return err
}

func getScriptMigrationText(migrationsApply int) string {
	if migrationsApply == 1 {
		return fmt.Sprintf("Migrating up %v version", migrationsApply)
	} else if migrationsApply == -1 {
		return fmt.Sprintf("Migrating down %v version", migrationsApply)
	} else if migrationsApply > 1 {
		return fmt.Sprintf("Migrating up %v versions", migrationsApply)
	} else {
		return fmt.Sprintf("Migrating down %v versions", migrationsApply)
	}
}
