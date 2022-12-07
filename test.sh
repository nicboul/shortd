#!/bin/bash

curl http://127.0.0.1:8080 \
-H 'Content-Type: application/json' \
-d \
'
	{"url": "htp://example.com/a-very-long-url"}
'
