@echo off

protoc-3.11.1-win64\bin\protoc.exe --go_out=src/game ./msg/*.proto

pause