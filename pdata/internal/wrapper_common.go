// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal // import "go.opentelemetry.io/collector/pdata/internal"

import (
	otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"
)

type Value struct {
	orig *otlpcommon.AnyValue
}

func GetOrigValue(ms Value) *otlpcommon.AnyValue {
	return ms.orig
}

func NewValue(orig *otlpcommon.AnyValue) Value {
	return Value{orig: orig}
}

type Map struct {
	orig *[]otlpcommon.KeyValue
}

func GetOrigMap(ms Map) *[]otlpcommon.KeyValue {
	return ms.orig
}

func NewMap(orig *[]otlpcommon.KeyValue) Map {
	return Map{orig: orig}
}

func FillTestValue(dest Value) {
	dest.orig.Value = &otlpcommon.AnyValue_StringValue{StringValue: "v"}
}

func GenerateTestValue() Value {
	var orig otlpcommon.AnyValue
	av := NewValue(&orig)
	FillTestValue(av)
	return av
}

func GenerateTestMap() Map {
	var orig []otlpcommon.KeyValue
	am := NewMap(&orig)
	FillTestMap(am)
	return am
}

func FillTestMap(dest Map) {
	*dest.orig = nil
	*dest.orig = append(*dest.orig, otlpcommon.KeyValue{Key: "k", Value: otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "v"}}})
}
