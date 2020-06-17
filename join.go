package goqry

import (
	"fmt"
	"strings"
)

// Join is a join
type Join struct {
	typ    JoinType
	from   string
	to     Tabler
	fields []Fielder
}

// Field returns a field for a join
func (j *Join) Field(fld string) Fielder {
	f := &Field{name: fld, tblname: j.from}
	j.fields = append(j.fields, f)
	return f
}

// To will return the to table
func (j *Join) To() Tabler {
	return j.to
}

// SQL will return a sql string
func (j *Join) SQL() string {
	return j.str("", "", false)
}

// Indent returns formatted SQL
func (j *Join) Indent(ln, in string) string {
	return j.str(ln, in, true)
}

// Indent returns a formatted string
// ln = line delimeter
// in = indention delimeter
func (j *Join) str(ln, in string, ind bool) string {
	if ln == "" {
		ln = " "
	}
	builder := strings.Builder{}
	jn := ""
	switch j.typ {
	case JoinInner:
		jn = "INNER JOIN"
	case JoinFull:
		jn = "FULL JOIN"
	case JoinLeft:
		jn = "LEFT JOIN"
	case JoinRight:
		jn = "RIGHT JOIN"
	}
	builder.WriteString(fmt.Sprintf("%s `%s`%s", jn, j.to.Name(), ln))
	if len(j.fields) == 0 {
		return builder.String()
	}
	flds := []string{}
	for _, fld := range j.fields {
		if ind {
			flds = append(flds, fld.Indent(ln, in))
		} else {
			flds = append(flds, fld.SQL())
		}
	}
	builder.WriteString(fmt.Sprintf("%sON %v", in, strings.Join(flds, " AND ")))
	return builder.String()
}
