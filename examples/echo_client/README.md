A simple echo client that echos any data received from EDDN to stdout.

## Building

    go build

## Running

    ./echo_client

## Instructions

    ./echo_client -h

Should be pretty self explanatory from the help output provided.  Please keep
in mind that this is **not** a common use-case for the library.  Generally
you'd want to be doing interesting things with the native go types provided
concurrently by the ChannelInterface.  However, this does show basic usage
and how easily moving to and from the native Go types and JSON is.
