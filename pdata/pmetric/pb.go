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

package pmetric // import "go.opentelemetry.io/collector/pdata/pmetric"

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

// NewProtoMarshaler returns a MarshalerSizer.
// Marshals to OTLP binary protobuf bytes and calculates the size of the marshaled Metrics.
func NewProtoMarshaler() MarshalerSizer {
	return newPbMarshaler()
}

type pbMarshaler struct{}

func newPbMarshaler() *pbMarshaler {
	return &pbMarshaler{}
}

var _ Sizer = (*pbMarshaler)(nil)

func (e *pbMarshaler) MarshalMetrics(md Metrics) ([]byte, error) {
	pb := internal.MetricsToProto(internal.Metrics(md))
	return pb.Marshal()
}

func (e *pbMarshaler) MetricsSize(md Metrics) int {
	pb := internal.MetricsToProto(internal.Metrics(md))
	return pb.Size()
}

type pbUnmarshaler struct{}

// NewProtoUnmarshaler returns a model.Unmarshaler. Unmarshals from OTLP binary protobuf bytes.
func NewProtoUnmarshaler() Unmarshaler {
	return newPbUnmarshaler()
}

func newPbUnmarshaler() *pbUnmarshaler {
	return &pbUnmarshaler{}
}

func (d *pbUnmarshaler) UnmarshalMetrics(buf []byte) (Metrics, error) {
	pb := otlpmetrics.MetricsData{}
	err := pb.Unmarshal(buf)
	return Metrics(internal.MetricsFromProto(pb)), err
}
