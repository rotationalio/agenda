# gRPC Mocking Tutorial

**Tutorial code for gRPC services and mocks**

## Agenda

During this continuous development session, we're going to do some live code to explore our current testing and mocking methodologies for gRPC servers. In it I will:

1. Create a simple gRPC service that depends on a database
2. Implement the server and a simple client on the CLI
3. Implement a database store utility
4. Implement a client-testing mock that demonstrates how to test the client code using "universal mocks"
5. Implement a mock for the database using sqlmock or just simple fixtures
6. Implement a mock server with a bufconn and show how to test the server RPCs.

This should take roughly an hour and a half and we will record the video if you can't make it.

Follows ons from this activity:

- Blog posts and Best Practices
- Identifying "hidden base knowledge" (e.g. about networking)
- Identifying "hidden advanced knowledge" (e.g. about how gRPC works)
- Considering the 5 key take aways
- Possible tutorials or other videos for future reference.

The goal of this activity is to ensure we can all get on the same page on our new gRPC testing and mocking practices.
