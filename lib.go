// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlinterpolate

import (
	"errors"
)

// Supported flavors.
const (
	invalidFlavor Flavor = iota

	MySQL
	PostgreSQL
	SQLite
	SQLServer
	CQL
	ClickHouse
	Presto
	Oracle
	Informix
)

// DefaultFlavor is the default flavor for all builders.
var DefaultFlavor = MySQL

var (
	// ErrInterpolateNotImplemented means the method or feature is not implemented right now.
	ErrInterpolateNotImplemented = errors.New("go-sqlbuilder: interpolation for this flavor is not implemented")

	// ErrInterpolateMissingArgs means there are some args missing in query, so it's not possible to
	// prepare a query with such args.
	ErrInterpolateMissingArgs = errors.New("go-sqlbuilder: not enough args when interpolating")

	// ErrInterpolateUnsupportedArgs means that some types of the args are not supported.
	ErrInterpolateUnsupportedArgs = errors.New("go-sqlbuilder: unsupported args when interpolating")
)

// Flavor is the flag to control the format of compiled sql.
type Flavor int

// String returns the name of f.
func (f Flavor) String() string {
	switch f {
	case MySQL:
		return "MySQL"
	case PostgreSQL:
		return "PostgreSQL"
	case SQLite:
		return "SQLite"
	case SQLServer:
		return "SQLServer"
	case CQL:
		return "CQL"
	case ClickHouse:
		return "ClickHouse"
	case Presto:
		return "Presto"
	case Oracle:
		return "Oracle"
	case Informix:
		return "Informix"
	}

	return "<invalid>"
}

// Interpolate parses sql returned by `Args#Compile` or `Builder`,
// and interpolate args to replace placeholders in the sql.
//
// If there are some args missing in sql, e.g. the number of placeholders are larger than len(args),
// returns ErrMissingArgs error.
func (f Flavor) Interpolate(sql string, args []interface{}) (string, error) {
	switch f {
	case MySQL:
		return mysqlInterpolate(sql, args...)
	case PostgreSQL:
		return postgresqlInterpolate(sql, args...)
	case SQLite:
		return sqliteInterpolate(sql, args...)
	case SQLServer:
		return sqlserverInterpolate(sql, args...)
	case CQL:
		return cqlInterpolate(sql, args...)
	case ClickHouse:
		return clickhouseInterpolate(sql, args...)
	case Presto:
		return prestoInterpolate(sql, args...)
	case Oracle:
		return oracleInterpolate(sql, args...)
	case Informix:
		return informixInterpolate(sql, args...)
	}

	return "", ErrInterpolateNotImplemented
}
