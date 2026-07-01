# Run the full check/build pipeline by default.
# .DEFAULT_GOAL is a special target that specifies the default goal for the Makefile.
# Incase not specified, the first target in the Makefile is used as the default goal.
.DEFAULT_GOAL := build 

# .PHONY line keeps Make from getting confused if a file named fmt, vet, or build exists in the directory.
#  Define target definitions for fmt, vet, and build as phony targets to avoid conflicts with files of the same name.
.PHONY: fmt vet build

# Format the Go files in this module.
fmt:
	go fmt .

# Run vet after formatting so static checks see cleaned-up code.
vet: fmt
	go vet .

# Build the module after formatting and vetting.
build: vet
	go build

