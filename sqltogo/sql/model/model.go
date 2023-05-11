package model

const indexPri = "PRIMARY"

type (

	// Column defines column in table
	Column struct {
		*DbColumn
		Index *DbIndex
	}

	// DbColumn defines column info of columns
	DbColumn struct {
		Name            string      `db:"COLUMN_NAME"`
		DataType        string      `db:"DATA_TYPE"`
		ColumnType      string      `db:"COLUMN_TYPE"`
		Extra           string      `db:"EXTRA"`
		Comment         string      `db:"COLUMN_COMMENT"`
		ColumnDefault   interface{} `db:"COLUMN_DEFAULT"`
		IsNullAble      string      `db:"IS_NULLABLE"`
		OrdinalPosition int         `db:"ORDINAL_POSITION"`
	}

	// DbIndex defines index of columns in information_schema.statistic
	DbIndex struct {
		IndexName  string `db:"INDEX_NAME"`
		NonUnique  int    `db:"NON_UNIQUE"`
		SeqInIndex int    `db:"SEQ_IN_INDEX"`
	}

	// ColumnData describes the columns of table
	ColumnData struct {
		Db      string
		Table   string
		Columns []*Column
	}

	// Table describes mysql table which contains database name, table name, columns, keys
	Table struct {
		Db      string
		Table   string
		Columns []*Column
		// Primary key not included
		UniqueIndex map[string][]*Column
		PrimaryKey  *Column
		NormalIndex map[string][]*Column
	}

	// IndexType describes an alias of string
	IndexType string

	// Index describes a column index
	Index struct {
		IndexType IndexType
		Columns   []*Column
	}
)
