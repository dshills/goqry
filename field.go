package goqry

import (
	"fmt"
	"strings"
)

type filter struct {
	v   string
	op  string
	fld Fielder
}

// Field is a field
type Field struct {
	tblname string
	isnum   bool
	name    string
	filters []filter
}

// SQL will return the sql for the field
func (f *Field) SQL() string {
	strs := []string{}
	for _, fil := range f.filters {
		switch {
		case fil.fld == nil && f.isnum:
			strs = append(strs, fmt.Sprintf("%v %v %v", f.Name(), fil.op, fil.v))
		case fil.fld == nil:
			strs = append(strs, fmt.Sprintf("%v %v '%v'", f.Name(), fil.op, fil.v))
		default:
			strs = append(strs, fmt.Sprintf("%v %v %v", f.Name(), fil.op, fil.fld.Name()))
		}
	}
	return strings.Join(strs, " AND ")
}

// Indent returns a formatted string
// ln = line delimeter
// in = indention delimeter
func (f *Field) Indent(ln, in string) string {
	strs := []string{}
	for _, fil := range f.filters {
		switch {
		case fil.fld == nil && f.isnum:
			strs = append(strs, fmt.Sprintf("%s%s %s %s", in, f.Name(), fil.op, fil.v))
		case fil.fld == nil:
			strs = append(strs, fmt.Sprintf("%s%s %s '%s'", in, f.Name(), fil.op, fil.v))
		default:
			strs = append(strs, fmt.Sprintf("%s%s %s %s", in, f.Name(), fil.op, fil.fld.Name()))
		}
	}
	return strings.Join(strs, " AND ")
}

// Eq sets the field filter
func (f *Field) Eq(v string) Fielder {
	fil := filter{op: "=", v: v}
	f.filters = append(f.filters, fil)
	return f
}

// NotEq sets the field filter
func (f *Field) NotEq(v string) Fielder {
	fil := filter{op: "!=", v: v}
	f.filters = append(f.filters, fil)
	return f
}

// Gr sets the field filter
func (f *Field) Gr(v string) Fielder {
	fil := filter{op: ">", v: v}
	f.filters = append(f.filters, fil)
	return f
}

// Ls sets the field filter
func (f *Field) Ls(v string) Fielder {
	fil := filter{op: "<", v: v}
	f.filters = append(f.filters, fil)
	return f
}

// GrEq sets the field filter
func (f *Field) GrEq(v string) Fielder {
	fil := filter{op: ">=", v: v}
	f.filters = append(f.filters, fil)
	return f
}

// LsEq sets the field filter
func (f *Field) LsEq(v string) Fielder {
	fil := filter{op: "<=", v: v}
	f.filters = append(f.filters, fil)
	return f
}

// EqField sets the field filter
func (f *Field) EqField(fld Fielder) Fielder {
	fil := filter{op: "=", fld: fld}
	f.filters = append(f.filters, fil)
	return f
}

// NotEqField sets the field filter
func (f *Field) NotEqField(fld Fielder) Fielder {
	fil := filter{op: "!=", fld: fld}
	f.filters = append(f.filters, fil)
	return f
}

// GrField sets the field filter
func (f *Field) GrField(fld Fielder) Fielder {
	fil := filter{op: ">", fld: fld}
	f.filters = append(f.filters, fil)
	return f
}

// LsField sets the field filter
func (f *Field) LsField(fld Fielder) Fielder {
	fil := filter{op: "<", fld: fld}
	f.filters = append(f.filters, fil)
	return f
}

// GrEqField sets the field filter
func (f *Field) GrEqField(fld Fielder) Fielder {
	fil := filter{op: ">=", fld: fld}
	f.filters = append(f.filters, fil)
	return f
}

// LsEqField sets the field filter
func (f *Field) LsEqField(fld Fielder) Fielder {
	fil := filter{op: "<=", fld: fld}
	f.filters = append(f.filters, fil)
	return f
}

// Null sets the field filter
func (f *Field) Null() Fielder {
	fil := filter{op: "IS NULL"}
	f.filters = append(f.filters, fil)
	return f
}

// NotNull sets the field filter
func (f *Field) NotNull() Fielder {
	fil := filter{op: "IS NOT NULL"}
	f.filters = append(f.filters, fil)
	return f
}

// Name will return the qualified name of the field
func (f *Field) Name() string {
	return fmt.Sprintf("%v.%v", f.tblname, f.name)
}

// MakeNum tells the field to behave as numeric
func (f *Field) MakeNum() Fielder {
	f.isnum = true
	return f
}

// IsNumeric returns true if field is numeric
func (f *Field) IsNumeric() bool {
	return f.isnum
}
