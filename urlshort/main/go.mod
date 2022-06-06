module example.com/main

go 1.15

replace example.com/urlshort => ../urlshort

require (
	example.com/urlshort v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.10.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
)
