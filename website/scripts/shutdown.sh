#!/usr/bin/env bash

if [ -f tmp/air.pid ]; then
    echo "Shutting down..."
    kill $(cat tmp/air.pid)
    rm -f tmp/air.pid tmp/build-errors.log
fi
