package goqry

// Tabler is a sql table interface
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

// Wherer is a sql WHERE interface
type Wherer interface {
	And(fld Fielder) Wherer
	Or(fld Fielder) Wherer
	Group() Wherer
	SQL() string
	Indent(lndelim, indelim string) string
}

// Fielder is a sql FIELD interface
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

// Joiner is a sql JOIN interface
type Joiner interface {
	Field(fld string) Fielder
	To() Tabler
	SQL() string
	Indent(lndelim, indelim string) string
}
