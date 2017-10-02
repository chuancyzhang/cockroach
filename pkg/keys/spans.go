// Copyright 2015 The Cockroach Authors.
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

package keys

import "github.com/cockroachdb/cockroach/pkg/roachpb"

//go:generate go run gen_cpp_keys.go

var (
	// MetaSpan holds all the addressing records.
	//
	// TODO(peter): Allow splitting of meta2 records. This requires enhancement
	// to range lookup. See #16266.
	//
	//  Meta1Span = roachpb.Span{Key: roachpb.KeyMin, EndKey: Meta2Prefix}
	MetaSpan = roachpb.Span{Key: roachpb.KeyMin, EndKey: MetaMax}

	// NodeLivenessSpan holds the liveness records for nodes in the cluster.
	NodeLivenessSpan = roachpb.Span{Key: NodeLivenessPrefix, EndKey: NodeLivenessKeyMax}

	// SystemConfigSpan is the range of system objects which will be gossiped.
	SystemConfigSpan = roachpb.Span{Key: SystemConfigSplitKey, EndKey: SystemConfigTableDataMax}

	// NoSplitSpans describes the ranges that should never be split.
	// MetaSpan: needed to find other ranges.
	// NodeLivenessSpan: liveness information on nodes in the cluster.
	// SystemConfigSpan: system objects which will be gossiped.
	NoSplitSpans = []roachpb.Span{MetaSpan, NodeLivenessSpan, SystemConfigSpan}
)

// Silence unused warning. This variable is actually used by gen_cpp_keys.go.
var _ = NoSplitSpans
