Missing Features:
- Index Files
- Constraints
- Nested Messages Support https://protobuf.dev/programming-guides/proto3/#nested
- Deep Recursion Support
- **FULL** Spectrum of Well Known Types https://protobuf.dev/reference/protobuf/google.protobuf/#index
- Dependency management when needed
- map<string,ANYTHING>, not only primitives
- Parameter ApiBaseURL, default is "/api"
- Additional Bindings

```
service Messaging {
  rpc GetMessage(GetMessageRequest) returns (Message) {
    option (google.api.http) = {
      get: "/v1/messages/{message_id}"
      additional_bindings {
        get: "/v1/users/{user_id}/messages/{message_id}"
      }
    };
  }
} 
```