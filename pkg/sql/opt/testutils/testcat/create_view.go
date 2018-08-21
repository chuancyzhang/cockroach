// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package testcat

import (
	"bytes"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

// CreateView creates a test view from a parsed DDL statement and adds it to the
// catalog.
func (tc *Catalog) CreateView(stmt *tree.CreateView) *View {
	tn, err := stmt.Name.Normalize()
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	// Update the view name to include catalog and schema if not provided.
	tc.qualifyTableName(tn)

	var buf bytes.Buffer
	fmtCtx := tree.MakeFmtCtx(&buf, tree.FmtParsable)
	stmt.AsSource.Format(&fmtCtx)

	view := &View{ViewName: *tn, QueryText: buf.String(), ColumnNames: stmt.ColumnNames}

	// Add the new view to the catalog.
	tc.AddView(view)

	return view
}