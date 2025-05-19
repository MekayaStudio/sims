package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034405CreateCbtKopBeritaTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034405CreateCbtKopBeritaTable) Signature() string {
	return "20250519034405_create_cbt_kop_berita_table"
}

// Up Run the migrations.
func (r *M20250519034405CreateCbtKopBeritaTable) Up() error {
	if !facades.Schema().HasTable("cbt_kop_berita") {
		return facades.Schema().Create("cbt_kop_berita", func(table schema.Blueprint) {
			table.ID("id_kop")
			table.String("header_1", 100).Nullable()
			table.String("header_2", 100).Nullable()
			table.String("header_3", 100).Nullable()
			table.String("header_4", 100).Nullable()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519034405CreateCbtKopBeritaTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_kop_berita")
}
