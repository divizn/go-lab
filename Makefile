# Default target
.DEFAULT_GOAL := help

.PHONY: build clean help


web-server: # Run web-server with Air
	@echo "Starting web-server with Air..."
	./run_air.sh web-server


hex2dec: # Run hex2dec
	@echo "Starting hex2dec..."
	go run cmd/hex2dec/main.go

dec2bin: # Run dec2bin
	@echo "Starting dec2bin..."
	go run cmd/dec2bin/main.go


clean: # Clean built binaries
	@echo "Cleaning tmp directory..."
	@rm -rf ./tmp

help: # Show help for each command.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
