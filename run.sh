#!/usr/bin/env bash

cd ~/projects/wifi-file-transfer

source venv/bin/activate

hostname -I | awk '{print "http://"$1":8000"}'

uvicorn main:app --host 0.0.0.0 --port 8000