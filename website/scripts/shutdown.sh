#!/usr/bin/env bash

air_pids=$(pgrep air)
if [ -n "$air_pids" ]; then
    echo "Shutting down..."
    kill $air_pids
else
    echo "No running server was found"
fi
