---
weight: 200
title: "Example"
---

# Example setup for grpc-gateway
If you need newline delimited  messages and raw streams, define a marhalerOption for the corresponding mimetype.


**Example**
```go
    // import 	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

        // Do not send  x-ndjson for gwruntime.MIMEWildcard (*)
		gwruntime.WithMarshalerOption(gwruntime.MIMEWildcard, &marshaller.HTTPBodyMarshaler{
			Marshaler: &gwruntime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
        
        // send newline delimeted json if client sends the corresponding accept type
		gwruntime.WithMarshalerOption("application/x-ndjson", &gwruntime.HTTPBodyMarshaler{
			Marshaler: &gwruntime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	}
 

```


```go

package marshaller

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

// HTTPBodyMarshaler is a Marshaler which supports marshaling of a
// google.api.HttpBody message as the full response body if it is
// the actual message used as the response. If not, then this will
// simply fallback to the Marshaler specified as its default Marshaler.
type HTTPBodyMarshaler struct {
	runtime.Marshaler
}

// ContentType returns its specified content type in case v is a
// google.api.HttpBody message, otherwise it will fall back to the default Marshalers
// content type.
func (h *HTTPBodyMarshaler) ContentType(v interface{}) string {
	if httpBody, ok := v.(*httpbody.HttpBody); ok {
		return httpBody.GetContentType()
	}
	return h.Marshaler.ContentType(v)
}

// Marshal marshals "v" by returning the body bytes if v is a
// google.api.HttpBody message, otherwise it falls back to the default Marshaler.
func (h *HTTPBodyMarshaler) Marshal(v interface{}) ([]byte, error) {
	if httpBody, ok := v.(*httpbody.HttpBody); ok {
		return httpBody.Data, nil
	}
	return h.Marshaler.Marshal(v)
}

// no Delimiter for httpbody
func (j *HTTPBodyMarshaler) Delimiter() []byte {
	return []byte{}
}

```