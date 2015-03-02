
go build
cp *.exe dist
rm -rf ./dist/templates
rm -rf ./dist/static
cp -r ./web/www/templates ./dist
cp -r ./web/www/static ./dist

cd dist 
devlibrary.exe