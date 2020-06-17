package goqry

import (
	"strings"
)

// Where is a where clause
type Where struct {
	typ     WhereType
	filters []Wherer
	fld     Fielder
}

// And will return an AND wherer
func (w *Where) And(fld Fielder) Wherer {
	nw := &Where{typ: WhereAnd, fld: fld}
	w.filters = append(w.filters, nw)
	return nw
}

// Or will return an OR wherer
func (w *Where) Or(fld Fielder) Wherer {
	nw := &Where{typ: WhereOr, fld: fld}
	w.filters = append(w.filters, nw)
	return nw
}

// Group will return a grouped wherer
func (w *Where) Group() Wherer {
	nw := &Where{typ: WhereGroup}
	w.filters = append(w.filters, nw)
	return nw
}

// SQL will return a sql string
func (w *Where) SQL() string {
	return w.str("", "", false)
}

// Indent returns formatted SQL
func (w *Where) Indent(ln, in string) string {
	return w.str(ln, in, true)
}

func (w *Where) str(ln, in string, ind bool) string {
	if ln == "" {
		ln = " "
	}
	builder := strings.Builder{}
	if w.typ == WhereGroup {
		builder.WriteRune('(')
		for _, flt := range w.filters {
			if ind {
				builder.WriteString(flt.Indent(ln, in))
			} else {
				builder.WriteString(flt.SQL())
			}
		}
		builder.WriteRune(')')
		return builder.String()
	}

	if w.fld != nil {
		switch w.typ {
		case WhereAnd:
			builder.WriteString(" AND ")
		case WhereOr:
			builder.WriteString(" OR ")
		}
		if ind {
			builder.WriteString(w.fld.Indent(ln, in))
		} else {
			builder.WriteString(w.fld.SQL())
		}
	}
	for _, flt := range w.filters {
		if ind {
			builder.WriteString(flt.Indent(ln, in))
		} else {
			builder.WriteString(flt.SQL())
		}
	}

	return builder.String()
}
