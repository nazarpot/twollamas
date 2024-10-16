@echo off
REM Script for simplifying input

:print_cleaner
set str=%1
set type=%str:~0,1%
set val=%str:~1%
if "%type%"=="p" (
    echo Pick up: %val%
) else if "%type%"=="d" (
    echo Drop off: %val%
) else if "%type%"=="w" (
    echo Weight: %val%
) else if "%type%"=="r" (
    echo Rate: %val% pm
) else if "%type%"=="t" (
    echo Temp: %val%
)
exit /b

REM Loop through each argument
for %%i in (%*) do (
    call :print_cleaner %%i
)
