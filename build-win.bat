del /F /S /Q govcl-template.exe
del /F /S /Q .\win64\govcl-template.exe
go build -buildmode=exe -i -ldflags="-H windowsgui"
powershell Copy-Item govcl-template.exe ./win64/
cd win64
govcl-template.exe

