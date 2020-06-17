# goqry

SQL Query Builder for Go

--
    import "."


## Usage

#### type Field

```go
type Field struct {
}
```

Field is a field

#### func (*Field) Eq

```go
func (f *Field) Eq(v string) Fielder
```
Eq sets the field filter

#### func (*Field) EqField

```go
func (f *Field) EqField(fld Fielder) Fielder
```
EqField sets the field filter

#### func (*Field) Gr

```go
func (f *Field) Gr(v string) Fielder
```
Gr sets the field filter

#### func (*Field) GrEq

```go
func (f *Field) GrEq(v string) Fielder
```
GrEq sets the field filter

#### func (*Field) GrEqField

```go
func (f *Field) GrEqField(fld Fielder) Fielder
```
GrEqField sets the field filter

#### func (*Field) GrField

```go
func (f *Field) GrField(fld Fielder) Fielder
```
GrField sets the field filter

#### func (*Field) Indent

```go
func (f *Field) Indent(ln, in string) string
```
Indent returns a formatted string ln = line delimeter in = indention delimeter

#### func (*Field) IsNumeric

```go
func (f *Field) IsNumeric() bool
```
IsNumeric returns true if field is numeric

#### func (*Field) Ls

```go
func (f *Field) Ls(v string) Fielder
```
Ls sets the field filter

#### func (*Field) LsEq

```go
func (f *Field) LsEq(v string) Fielder
```
LsEq sets the field filter

#### func (*Field) LsEqField

```go
func (f *Field) LsEqField(fld Fielder) Fielder
```
LsEqField sets the field filter

#### func (*Field) LsField

```go
func (f *Field) LsField(fld Fielder) Fielder
```
LsField sets the field filter

#### func (*Field) MakeNum

```go
func (f *Field) MakeNum() Fielder
```
MakeNum tells the field to behave as numeric

#### func (*Field) Name

```go
func (f *Field) Name() string
```
Name will return the qualified name of the field

#### func (*Field) NotEq

```go
func (f *Field) NotEq(v string) Fielder
```
NotEq sets the field filter

#### func (*Field) NotEqField

```go
func (f *Field) NotEqField(fld Fielder) Fielder
```
NotEqField sets the field filter

#### func (*Field) NotNull

```go
func (f *Field) NotNull() Fielder
```
NotNull sets the field filter

#### func (*Field) Null

```go
func (f *Field) Null() Fielder
```
Null sets the field filter

#### func (*Field) SQL

```go
func (f *Field) SQL() string
```
SQL will return the sql for the field

#### type Fielder

```go
type Fielder interface {
	Eq(v string) Fielder
	NotEq(v string) Fielder
	Gr(v string) Fielder
	Ls(v string) Fielder
	GrEq(v string) Fielder
	LsEq(v string) Fielder
	EqField(fld Fielder) Fielder
	NotEqField(fld Fielder) Fielder
	GrField(fld Fielder) Fielder
	LsField(fld Fielder) Fielder
	GrEqField(fld Fielder) Fielder
	LsEqField(fld Fielder) Fielder
	Null() Fielder
	NotNull() Fielder
	Name() string
	SQL() string
	Indent(lndelim, indelim string) string
	IsNumeric() bool
	MakeNum() Fielder
}
```

Fielder is a sql FIELD interface

#### type Join

```go
type Join struct {
}
```

Join is a join

#### func (*Join) Field

```go
func (j *Join) Field(fld string) Fielder
```
Field returns a field for a join

#### func (*Join) Indent

```go
func (j *Join) Indent(ln, in string) string
```
Indent returns formatted SQL

#### func (*Join) SQL

```go
func (j *Join) SQL() string
```
SQL will return a sql string

#### func (*Join) To

```go
func (j *Join) To() Tabler
```
To will return the to table

#### type JoinType

```go
type JoinType int
```

JoinType is the type of a join

```go
const (
	JoinInner JoinType = iota
	JoinFull
	JoinLeft
	JoinRight
)
```
Join types

#### type Joiner

```go
type Joiner interface {
	Field(fld string) Fielder
	To() Tabler
	SQL() string
	Indent(lndelim, indelim string) string
}
```

Joiner is a sql JOIN interface

#### type Table

```go
type Table struct {
}
```

Table is a table

#### func  NewTable

```go
func NewTable(name string) *Table
```
NewTable will return a table

#### func (*Table) Field

```go
func (t *Table) Field(fld string) Fielder
```
Field returns a field for a join

#### func (*Table) FullJoin

```go
func (t *Table) FullJoin(tbl Tabler) Joiner
```
FullJoin will return a joiner to another table

#### func (*Table) Indent

```go
func (t *Table) Indent(ln, in string) string
```
Indent return formatted SQL

#### func (*Table) InnerJoin

```go
func (t *Table) InnerJoin(tbl Tabler) Joiner
```
InnerJoin will return a inner joiner to another table

#### func (*Table) LeftJoin

```go
func (t *Table) LeftJoin(tbl Tabler) Joiner
```
LeftJoin will return a left joiner to another table

#### func (*Table) Name

```go
func (t *Table) Name() string
```
Name returns the table name

#### func (*Table) RightJoin

```go
func (t *Table) RightJoin(tbl Tabler) Joiner
```
RightJoin will return a table join

#### func (*Table) SQL

```go
func (t *Table) SQL() string
```
SQL returns a sql query

#### func (*Table) Select

```go
func (t *Table) Select(flds ...string) Tabler
```
Select will return fields in the query for the table

#### func (*Table) Selections

```go
func (t *Table) Selections() []string
```
Selections returns the SELECT portion of the table

#### func (*Table) Where

```go
func (t *Table) Where(flds ...Fielder) Wherer
```
Where returns a where clause

#### func (*Table) WhereString

```go
func (t *Table) WhereString(ln, in string) string
```
WhereString returns the WHERE portion of the table

#### type Tabler

```go
type Tabler interface {
	Field(fld string) Fielder
	Select(flds ...string) Tabler
	InnerJoin(tbl Tabler) Joiner
	LeftJoin(tbl Tabler) Joiner
	RightJoin(tbl Tabler) Joiner
	FullJoin(tbl Tabler) Joiner
	Where(fld ...Fielder) Wherer
	SQL() string
	Indent(lndelim, indelim string) string
	Selections() []string
	WhereString(ln, in string) string
	Name() string
}
```

Tabler is a sql table interface

#### type Where

```go
type Where struct {
}
```

Where is a where clause

#### func (*Where) And

```go
func (w *Where) And(fld Fielder) Wherer
```
And will return an AND wherer

#### func (*Where) Group

```go
func (w *Where) Group() Wherer
```
Group will return a grouped wherer

#### func (*Where) Indent

```go
func (w *Where) Indent(ln, in string) string
```
Indent returns formatted SQL

#### func (*Where) Or

```go
func (w *Where) Or(fld Fielder) Wherer
```
Or will return an OR wherer

#### func (*Where) SQL

```go
func (w *Where) SQL() string
```
SQL will return a sql string

#### type WhereType

```go
type WhereType int
```

WhereType is the type for a where clause

```go
const (
	WhereNone WhereType = iota
	WhereAnd
	WhereOr
	WhereGroup
)
```
where types

#### type Wherer

```go
type Wherer interface {
	And(fld Fielder) Wherer
	Or(fld Fielder) Wherer
	Group() Wherer
	SQL() string
	Indent(lndelim, indelim string) string
}
```

Wherer is a sql WHERE interface
