#!/bin/sh
chmod x down.sh
migrate -path ../../migrations -database mysql://reihan:reihan@/ktp -verbose down
migrate -path ../../migrations -database mysql://reihan:reihan@/ktp -verbose up