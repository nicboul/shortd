#!/bin/bash

curl -v http://127.0.0.1:8080 \
-H 'Content-Type: application/json' \
-d \
'
	{"url": "http://example.com/a-very-long-url"}
'
