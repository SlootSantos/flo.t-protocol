# flo.t-server
## flo.t Protocol
A naïve implementation of a custom communciation protocol to be send over TCP or UDP.
Taking experience from the HTTP specification _(see [RFC spec](https://www.rfc-editor.org/rfc/rfc2616))_

### Message
Each message consists of
- a start line [1] to define e.g. version and the requested resource
- line-separated request-headers [2]
- body [3]

### [1] Start line
- first location: request resource, e.g. "/resource"
- second location: protocol version, e.g. flo.t/1.0

### [2] Request Headers
- max length = 1KB
- ASCII only
- line-delimited per header
- key-value pair separated by ":"
- e.g. "Example-Header: Value1"

### [3] Body
- can contain multiple byte buffers

### Example flo.t message
```
/resource flo.t/1.0
Example-Header: Value1
Content-Type: plain/text

Hello World
``´