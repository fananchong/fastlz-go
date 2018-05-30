set CURDIR=%~dp0
set BASEDIR=%CURDIR%\..\..\..\..\
set GOPATH=%BASEDIR%
echo %GOPATH%

cd tests\c\bin
call ctest.exe rand
call cbenchmark.exe

copy /y %CURDIR%\tests\rand.bin %CURDIR%\benchmarks

cd %CURDIR%\benchmarks
call go test -v -test.bench=".*" -count=1

cd %CURDIR%