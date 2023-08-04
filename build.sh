
appName="chrome-tabs"
zipName=$appName

version=$(git describe --abbrev=0 --tags)
if test -z $version
then
  echo "version empty"
  version="dev"
fi

echo "build darwin amd64 package"
GOOS=darwin GOARCH=amd64 go build -o "$appName"-darwin-amd64 .

echo "build darwin arm64 package"
GOOS=darwin GOARCH=arm64 go build -o "$appName"-darwin-arm64 .

echo "merge darwin amd64 and arm64"

lipo -create -output "$appName"-darwin \
 "$appName"-darwin-amd64 \
 "$appName"-darwin-arm64

rm -r "$appName"-darwin-amd64
rm -r "$appName"-darwin-arm64

test ! -d build/dist && mkdir build/dist

test -d build/dist/$version && rm -rf "build/dist/$version"
mkdir "build/dist/$version"

mv "$appName"-darwin "build/dist/$version/"

cp alfred/chrome-tabs.alfredworkflow "build/dist/$version/chrome-tabs-$version.alfredworkflow"

cd "build/dist/$version/"
zip "$zipName-$version.zip" *

