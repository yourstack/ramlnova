#!/bin/sh
go build
mv ramlnova bin/linux64/ramlnova
cp -rf template bin/linux64/template
# rsync -azP bin/ramlnova root@api.anasit.com:/opt/anasit/ramlnova
git add .
git commit -m "update"
git push -u origin master

