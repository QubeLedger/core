cd proto
buf mod update
buf generate
cd ..

cp -r github.com/QubeLedger/core/* ./
rm -rf github.com