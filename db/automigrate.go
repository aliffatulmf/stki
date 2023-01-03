package db

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Tables []any

func MakeTable(rows ...any) (Tables, error) {
	var tabs Tables

	for _, row := range rows {
		if reflect.ValueOf(row).Kind() != reflect.Struct {
			return tabs, fmt.Errorf("MakeTable: %v type is not struct", reflect.TypeOf(row))
		}

		tabs = append(tabs, row)
	}

	return tabs, nil
}

type AutoMigrate struct {
	// migrator to perform io operations with database
	Migrator gorm.Migrator
	// a collection of tables consisting of a model struct
	Tables []any
}

func NewAutoMigrate(mig gorm.Migrator) *AutoMigrate {
	return &AutoMigrate{
		Migrator: mig,
	}
}

func (am *AutoMigrate) MigrateTable(t Tables) error {
	for _, model := range t {
		if am.Migrator.HasTable(model) {
			fmt.Printf("Deleting %+v\n", model)
			if err := am.Migrator.DropTable(model); err != nil {
				return fmt.Errorf("MigratorDropTable: %w", err)
			}
		}

		fmt.Printf("Trying to create %+v\n", model)
		if err := am.Migrator.CreateTable(model); err != nil {
			return fmt.Errorf("MigratorCreateTable: %w", err)
		}
	}

	return nil
}
