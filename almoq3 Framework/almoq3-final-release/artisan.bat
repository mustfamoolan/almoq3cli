@echo off
where almoq3 >nul 2>&1
if %errorlevel% neq 0 (
    if exist "..\almoq3.exe" (
        "..\almoq3.exe" %*
    ) else (
        echo Error: almoq3 CLI not found in PATH or parent directory.
        exit /b 1
    )
) else (
    almoq3 %*
)
