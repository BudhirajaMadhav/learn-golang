##Why Close() response.body?

1. If the Body is not both read to EOF and closed, the Client’s underlying RoundTripper (typically Transport) may not be able to re-use a persistent TCP connection to the server for a subsequent “keep-alive” request.
    - What's a roundTripper: "round-trip": ghoom ke waps aana:          RoundTripper is an interface representing the ability to execute a single HTTP transaction, obtaining the Response for a given Request." It sits in between the low level stuff like dialing, tcp, etc. and the high level details of HTTP (redirects, etc.) RoundTrip is the method do do a single round trip of request sent to server, server answers with response

2. So essentially, transport may not reuse HTTP/1.x “keep-alive” TCP connections if the Body is not read to completion and closed.
3. Also to mention that it is the caller’s responsibility to close the response body


Some important guidelines to follow 

- Use the defer method to close the response body. This is done to make sure that the response body gets closed even in case of runtime error during reading and parsing of response.

- Check for error first always. Only proceed to close the response body using defer if the error is nil. A nil error always indicates a non-nil response body. If the error is non-nil then handle that error and return. **Also, note that closing a nil response body will cause a panic.**
