setlocal
set "ARCHS=amd64 386"

if not "%1" == "" goto %1

    for %%I in (%ARCHS%) do (
        set "GOARCH=%%I"
        mkdir "%%I" 2>nul
        go build -o "%%I/sponge.exe"
    )
    goto end

:package
    for %%I in (%ARCHS%) do (
        zip -j sponge-%DATE:/=%-%%I.zip readme.md %%I\sponge.exe 
    )
    goto end

:end

endlocal
