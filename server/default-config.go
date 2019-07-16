package server

const defaultConfig = `
[Server]
ListenAddress = "127.0.0.1:8088"
BasePath = "/files"
# BasePath = "https://example.com/files" # external URL for use behind reverse proxy
# CorsOrigins = [ "http://example.com" , "https://example.org" ]

# Requests from these networks will have their X-Forwarded-For headers trusted
TrustedReverseProxyRanges = [
	"10.0.0.0/8",
	"172.16.0.0/12",
	"192.168.0.0/16",
	"fc00::/7",
	"127.0.0.0/8",
	"::1/128",
]

[Storage]
Path = "./uploads"
ShardLayers = 6
MaximumUploadSize = "10 MB" # accepts units such as: MB, g, tB, peta, kilobytes, gigabyte

[Database]
Type = "sqlite3" # sqlite3 | mysql

# for sqlite3: a filesystem path
# for mysql: a DSN like "user:password@tcp(127.0.0.1:3306)/database". see https://github.com/go-sql-driver/mysql#dsn-data-source-name
Path = "./uploads.db"

[Expiration]
# Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
MaxAge = "24h" # 1 day
IdentifiedMaxAge = "168h" # 1 week
CheckInterval = "5m"

# If EXTJWT is supported by the gateway or network, a validated token with an account present (when
# the user is authenticated to an irc services account) will use the IdentifiedMaxAge setting above
# instead of the base MaxAge.
#
# The HMAC secret used to sign the token is needed here to be able to validate the token.
#
# When using a webircgateway, the issuer will be the network_common_address of the upstream server
# if set. Otherwise it will be the hostname used to connect to the network.
[JwtSecretsByIssuer]
# "example.com" = "examplesecret"
# "169.254.0.0" = "anothersecret"

[[Loggers]]
Level = "info" # debug | info | warn | error | fatal | panic
Format = "json" # pretty | json
Output = "stderr:" # stderr: | stdout: | file:/path | udp:ip:port | unix:/path

# [[Loggers]]
# Level = "debug"
# Format = "json"
# Output = "file:./debug.log"

# [[Loggers]]
# # example receiver for testing: socat UDP-RECV:10101,bind=127.0.0.1 STDOUT
# Level = "info"
# Format = "json"
# Output = "udp:127.0.0.1:10101"

# [[Loggers]]
# # example receiver for testing: socat UNIX-LISTEN:./log.sock,fork STDOUT
# Level = "info"
# Format = "json"
# Output = "unix:./log.sock" # filesystem path to unix socket

`
