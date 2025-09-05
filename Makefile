# Default target
.DEFAULT_GOAL := help

.PHONY: build clean help


web-server: # Run web-server with Air
	@echo "Starting web-server with Air..."
	./run_air.sh web-server


hex2dec: # Run hex2dec with Air
	@echo "Starting hex2dec tool with Air..."
	./run_air.sh hex2dec

dec2bin: # Run dec2bin with Air
	@echo "Starting dec2bin with Air..."
	./run_air.sh dec2bin


clean: # Clean built binaries
	@echo "Cleaning tmp directory..."
	@rm -rf ./tmp

help: # Show help for each command.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
