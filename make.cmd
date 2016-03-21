setlocal

for %%I in (amd64 386) do (
    set "GOARCH=%%I"
    mkdir "%%I" 2>nul
    go build -o "%%I/sponge.exe"
)

endlocal
