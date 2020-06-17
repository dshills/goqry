package goqry

import (
	"fmt"
	"strings"
)

// Table is a table
type Table struct {
	name  string
	sel   []string
	joins []Joiner
	where []Wherer
}

// Field returns a field for a join
func (t *Table) Field(fld string) Fielder {
	return &Field{name: fld, tblname: t.name}
}

// Name returns the table name
func (t *Table) Name() string { return t.name }

// SQL returns a sql query
func (t *Table) SQL() string {
	return t.str("", "", false)
}

// Indent return formatted SQL
func (t *Table) Indent(ln, in string) string {
	return t.str(ln, in, true)
}

func (t *Table) str(ln, in string, ind bool) string {
	builder := strings.Builder{}

	builder.WriteString("SELECT" + ln)
	builder.WriteString(strings.Join(t.Selections(), ","+ln))

	builder.WriteString(fmt.Sprintf("%sFROM `%s`%s", ln, t.name, ln))
	for _, j := range t.joins {
		if ind {
			builder.WriteString(j.Indent(ln, in))
		} else {
			builder.WriteString(j.SQL())
		}
	}

	wbuilder := strings.Builder{}
	wbuilder.WriteString(t.WhereString(ln, in))
	for _, j := range t.joins {
		if tbl := j.To(); tbl != nil {
			wbuilder.WriteString(tbl.WhereString(ln, in))
		}
	}
	where := wbuilder.String()
	if len(where) > 0 {
		builder.WriteString("WHERE" + ln)
		builder.WriteString(where)
	}

	return builder.String()
}

// Selections returns the SELECT portion of the table
func (t *Table) Selections() []string {
	selection := []string{}
	selection = append(selection, t.sel...)
	for _, j := range t.joins {
		if tbl := j.To(); tbl != nil {
			selection = append(selection, tbl.Selections()...)
		}
	}
	return selection
}

// WhereString returns the WHERE portion of the table
func (t *Table) WhereString(ln, in string) string {
	if len(t.where) == 0 {
		return ""
	}
	builder := strings.Builder{}
	for _, w := range t.where {
		if ln != "" || in != "" {
			builder.WriteString(w.Indent(ln, in))
		} else {
			builder.WriteString(w.SQL())
		}
	}
	return builder.String()
}

// Select will return fields in the query for the table
func (t *Table) Select(flds ...string) Tabler {
	for _, f := range flds {
		t.sel = append(t.sel, fmt.Sprintf("%s.%s", t.name, f))
	}
	return t
}

// FullJoin will return a joiner to another table
func (t *Table) FullJoin(tbl Tabler) Joiner {
	return t.join(tbl, JoinFull)
}

// InnerJoin will return a inner joiner to another table
func (t *Table) InnerJoin(tbl Tabler) Joiner {
	return t.join(tbl, JoinInner)
}

// LeftJoin will return a left joiner to another table
func (t *Table) LeftJoin(tbl Tabler) Joiner {
	return t.join(tbl, JoinLeft)
}

// RightJoin will return a table join
func (t *Table) RightJoin(tbl Tabler) Joiner {
	return t.join(tbl, JoinRight)
}

func (t *Table) join(tbl Tabler, jt JoinType) Joiner {
	j := &Join{typ: jt, to: tbl, from: t.name}
	t.joins = append(t.joins, j)
	return j
}

// Where returns a where clause
func (t *Table) Where(flds ...Fielder) Wherer {
	w := &Where{}
	t.where = append(t.where, w)
	if len(flds) == 0 {
		return w
	}
	w.fld = flds[0]
	if len(flds) > 1 {
		for i := 1; i < len(flds); i++ {
			w.And(flds[i])
		}
	}
	return w
}

// NewTable will return a table
func NewTable(name string) *Table {
	return &Table{name: name}
}
