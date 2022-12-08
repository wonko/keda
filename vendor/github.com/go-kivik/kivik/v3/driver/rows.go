package driver

import (
	"encoding/json"
)

// Row is a generic view result row.
type Row struct {
	// ID is the document ID of the result.
	ID string `json:"id"`
	// Key is the view key of the result. For built-in views, this is the same
	// as ID.
	Key json.RawMessage `json:"key"`
	// Value is the raw, un-decoded JSON value. For most built-in views (such as
	// /_all_docs), this is `{"rev":"X-xxx"}`.
	Value json.RawMessage `json:"value"`
	// Doc is the raw, un-decoded JSON document. This is only populated by views
	// which return docs, such as /_all_docs?include_docs=true.
	Doc json.RawMessage `json:"doc"`
	// Error represents the error for any row not fetched. Usually just
	// 'not_found'.
	Error error `json:"-"`
}

// Rows is an iterator over a view's results.
type Rows interface {
	// Next is called to populate row with the next row in the result set.
	//
	// Next should return io.EOF when there are no more rows.
	Next(row *Row) error
	// Close closes the rows iterator.
	Close() error
	// UpdateSeq is the update sequence of the database, if requested in the
	// result set.
	UpdateSeq() string
	// Offset is the offset where the result set starts.
	Offset() int64
	// TotalRows is the number of documents in the database/view.
	TotalRows() int64
}

// RowsWarner is an optional interface that may be implemented by a Rows, which
// allows a rows iterator to return a non-fatal warning. This is intended for
// use by the /_find endpoint, which generates warnings when indexes don't
// exist.
type RowsWarner interface {
	// Warning returns the warning generated by the query, if any.
	Warning() string
}

// Bookmarker is an optional interface that may be implemented by a Rows for
// returning a paging bookmark.
type Bookmarker interface {
	// Bookmark returns an opaque bookmark string used for paging, added to
	// the /_find endpoint in CouchDB 2.1.1.  See the CouchDB documentation for
	// usage: http://docs.couchdb.org/en/2.1.1/api/database/find.html#pagination
	Bookmark() string
}

// QueryIndexer is an optional interface that may be implemented by a Rows,
// which allows a rows iterator to return a query index value. This is intended
// for use by multi-query queries to views.
type QueryIndexer interface {
	QueryIndex() int
}
