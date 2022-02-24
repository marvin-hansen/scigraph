@REM Define the federation name.
@set FDNAME=example2

@REM Define the federation data directory.
@set DATA_DIR=data

@REM Define the name and path for the federation's configuration file.
@set BOOT_FILE=%FDNAME%.boot
@set BOOT_PATH=%DATA_DIR%\%BOOT_FILE%

echo Checking if Lock Server is running.
objy checkls -notitle
@if %errorlevel% neq 0 (
  @echo Please start the lock server before running this script.
  @echo To start the lockserver run the following command from a user account that has admin privileges:
  @echo objy startlockserver
  @exit /b 1
)

@REM Create a Data directory
@if not exist "%DATA_DIR%" (
  mkdir "%DATA_DIR%"
)

@REM Cleanly remove a previous run of this script.
@if exist %BOOT_PATH% (
    objy installfd -notitle -bootfile %BOOT_PATH%
    objy cleanupfd -notitle -bootfile %BOOT_PATH% -local
    objy deletefd  -notitle -bootfile %BOOT_PATH%
  )


@REM Create the federation.
@REM A federation is composed of a configuration (*.boot) and data (*.fdb) file.
objy createFD -notitle -fdName %FDNAME% -fddirp .\%DATA_DIR%
@if %errorlevel% neq 0 exit /b %errorlevel%

@REM Execute a Do script to create the federation's schema.
objy do -notitle -bootfile %BOOT_PATH% -infile createSchema.do
@if %errorlevel% neq 0 exit /b %errorlevel%

@REM Create an index named Person_id for the person id attribute.
objy addindex -notitle -bootfile %BOOT_PATH% -name Person_id -class Person -attribute id
@if %errorlevel% neq 0 exit /b %errorlevel%

@REM Create an index named Phone_number for the phone number attribute.
objy addIndex -notitle -bootfile %BOOT_PATH% -name Phone_number -class Phone -attribute phoneNumber
@if %errorlevel% neq 0 exit /b %errorlevel%

@REM Populate the database with example data via a DO script.
@REM The script reads in data from CSV files.
objy do -notitle -bootfile %BOOT_PATH% -infile createData.do
@if %errorlevel% neq 0 exit /b %errorlevel%

@REM Query the database via a Do script and store the results in a file.
objy do -notitle -bootfile %BOOT_PATH% -infile runQuery.do -outfile queryResults.txt
@if %errorlevel% neq 0 exit /b 1

@type queryResults.txt
