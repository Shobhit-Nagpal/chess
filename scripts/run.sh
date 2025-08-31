#!/bin/bash

cd apps/server && air &
SERVER_PID=$!

cd apps/web && npm run dev &
WEB_PID=$!

trap "kill $SERVER_PID $WEB_PID" EXIT
wait
