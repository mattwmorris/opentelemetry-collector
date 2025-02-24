// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal // import "go.opentelemetry.io/collector/pdata/internal/cmd/pdatagen/internal"

var commonFile = &File{
	Name:        "common",
	PackageName: "pcommon",
	imports: []string{
		`otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"`,
	},
	testImports: []string{
		`"testing"`,
		``,
		`"github.com/stretchr/testify/assert"`,
		``,
		`"go.opentelemetry.io/collector/pdata/internal"`,
		`otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"`,
	},
	structs: []baseStruct{
		scope,
		attributeValueSlice,
	},
}

var scope = &messageValueStruct{
	structName:     "InstrumentationScope",
	packageName:    "pcommon",
	description:    "// InstrumentationScope is a message representing the instrumentation scope information.",
	originFullName: "otlpcommon.InstrumentationScope",
	fields: []baseField{
		nameField,
		&primitiveField{
			fieldName:       "Version",
			originFieldName: "Version",
			returnType:      "string",
			defaultVal:      `""`,
			testVal:         `"test_version"`,
		},
		attributes,
		droppedAttributesCount,
	},
}

// This will not be generated by this class.
// Defined here just to be available as returned message for the fields.
var mapStruct = &sliceOfPtrs{
	structName:  "Map",
	packageName: "pcommon",
	element:     attributeKeyValue,
}

var attributeKeyValue = &messageValueStruct{
	structName:  "Map",
	packageName: "pcommon",
}

var scopeField = &messageValueField{
	fieldName:       "Scope",
	originFieldName: "Scope",
	returnMessage:   scope,
}

var timestampType = &primitiveType{
	structName:  "Timestamp",
	packageName: "pcommon",
	rawType:     "uint64",
	defaultVal:  "0",
	testVal:     "1234567890",
}

var startTimeField = &primitiveTypedField{
	fieldName:       "StartTimestamp",
	originFieldName: "StartTimeUnixNano",
	returnType:      timestampType,
}

var timeField = &primitiveTypedField{
	fieldName:       "Timestamp",
	originFieldName: "TimeUnixNano",
	returnType:      timestampType,
}

var endTimeField = &primitiveTypedField{
	fieldName:       "EndTimestamp",
	originFieldName: "EndTimeUnixNano",
	returnType:      timestampType,
}

var attributes = &sliceField{
	fieldName:       "Attributes",
	originFieldName: "Attributes",
	returnSlice:     mapStruct,
}

var nameField = &primitiveField{
	fieldName:       "Name",
	originFieldName: "Name",
	returnType:      "string",
	defaultVal:      `""`,
	testVal:         `"test_name"`,
}

var anyValue = &messageValueStruct{
	structName:     "Value",
	packageName:    "pcommon",
	originFullName: "otlpcommon.AnyValue",
}

var attributeValueSlice = &sliceOfValues{
	structName: "Slice",
	element:    anyValue,
}

var traceIDField = &primitiveTypedField{
	fieldName:       "TraceID",
	originFieldName: "TraceId",
	returnType:      traceIDType,
}

var traceIDType = &primitiveType{
	structName:  "TraceID",
	packageName: "pcommon",
	rawType:     "data.TraceID",
	defaultVal:  "data.NewTraceID([16]byte{})",
	testVal:     "data.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})",
}

var spanIDField = &primitiveTypedField{
	fieldName:       "SpanID",
	originFieldName: "SpanId",
	returnType:      spanIDType,
}

var parentSpanIDField = &primitiveTypedField{
	fieldName:       "ParentSpanID",
	originFieldName: "ParentSpanId",
	returnType:      spanIDType,
}

var spanIDType = &primitiveType{
	structName:  "SpanID",
	packageName: "pcommon",
	rawType:     "data.SpanID",
	defaultVal:  "data.NewSpanID([8]byte{})",
	testVal:     "data.NewSpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})",
}

var schemaURLField = &primitiveField{
	fieldName:       "SchemaUrl",
	originFieldName: "SchemaUrl",
	returnType:      "string",
	defaultVal:      `""`,
	testVal:         `"https://opentelemetry.io/schemas/1.5.0"`,
}
