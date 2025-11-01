@echo off
REM Ensure using Go environment variables
cd /d %~dp0

REM Execute main.go
go run main.go

pause
