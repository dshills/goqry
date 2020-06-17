package goqry

// JoinType is the type of a join
type JoinType int

// Join types
const (
	JoinInner JoinType = iota
	JoinFull
	JoinLeft
	JoinRight
)

// WhereType is the type for a where clause
type WhereType int

// where types
const (
	WhereNone WhereType = iota
	WhereAnd
	WhereOr
	WhereGroup
)
