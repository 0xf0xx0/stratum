# A Golang library for the Bitcoin Stratum Protocol

Stratum is how Bitcoin miners connect with mining pools.

See [here for a description of Stratum](https://web.archive.org/web/20210224235216/https://braiins.com/stratum-v1/docs)
and
[here for Stratum extensions](https://github.com/slushpool/stratumprotocol/blob/master/stratum-extensions.mediawiki).
Extensions are necessary to support ASIC Boost.

forked from https://github.com/kbnchk/go-Stratum

## Importing

use master for now, eventually ill do a 1.0.0 release

```
go get git.0xf0xx0.eth.limo/0xf0xx0/stratum@master
```

## TODO

- Implement everything to get up to stratum 1.1 support
- work out all the bugs
- streaming support?

# Supported BIPs

- 310

## Supported Methods

- client.show_message
- mining.authorize
- mining.configure
- mining.notify
- mining.ping
- mining.set_difficulty
- mining.set_version_mask
- mining.submit
- mining.subscribe
- mining.suggest_difficulty

## Unsupported methods

- client.get_version
- client.reconnect
- mining.extranonce.subscribe
- mining.get_transactions
- mining.set_extranonce
- mining.suggest_target

## Method types

Some methods are client-to-server, others are server-to-client. Some methods
require a response, others do not.

### Client-to-server

| method                      | type               |
| --------------------------- | ------------------ |
| mining.authorize            | request / response |
| mining.configure            | request / response |
| mining.extranonce.subscribe | request / response |
| mining.get_transactions     | request / response |
| mining.ping                 | request / response |
| mining.submit               | request / response |
| mining.subscribe            | request / response |
| mining.suggest_difficulty   | notification       |
| mining.suggest_target       | notification       |

### Server-to-client

| method                  | type               |
| ----------------------- | ------------------ |
| client.get_version      | request / response |
| client.reconnect        | notification       |
| client.show_message     | notification       |
| mining.notify           | notification       |
| mining.ping             | request / response |
| mining.set_difficulty   | notification       |
| mining.set_extranonce   | notification       |
| mining.set_version_mask | notification       |

```
TODO: replace all this with a link to bip-41, mayb
```

## Message Formats

Stratum uses json. There are three message types: notification, request, and response.

### Notification

Notification is for methods that don't require a response.

```
{
  method: string,        // one of the methods above
  params: [json...]      // array of json values
}
```

### Request

Request is for methods that require a response.

```
{
  "id": uint or string, // a unique id, typically a hex-encoded number
  "method": string,        // one of the methods above
  "params": [json...]      // array of json values
}
```

### Response

Response is the response to requests.

```
{
  "id": integer or string,   // a unique id, must be the same as on the request
  "result": json,            // typically a boolean response
  "error": null or [
    uint,                    // error code
    string                   // error message
  ]
}
```

## Methods

### mining.authorize

The first message that is sent in classic Stratum. For extended Stratum,
`mining.configure` comes first and then `mining.authorize`.

### mining.subscribe

Sent by the client after `mining.authorize`.

### mining.set_difficulty

Sent by the server after responding to `mining.subscribe` and every time
the difficulty changes.

### mining.notify

Sent by the server whenever the block is updated. This happens periodically
as the block is being built and when a new block is discovered.

### mining.configure

The first message that is sent in extended Stratum. Necessary for ASICBoost.
The client sends this to tell the server what extensions it supports.

### mining.set_version_mask

Sent by the server to notify of a change in version mask. Requires the
`version-rolling` extension.

### mining.submit

Sent by the client when a new share is mined. Modified by `version-rolling`.

### mining.ping

Sent by either the server or client at any time. Expects a boolean response.
