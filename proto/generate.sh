cd proto
buf mod update
buf generate
cd ..

cp -r github.com/QuadrateOrg/core/* ./
rm -rf github.com