# runb2bqueries

git clone https://github.com/mft-labs/runb2bqueries

Pull latest changes
-------------------
git pull origin main

Compile Application
--------------------

go mod tidy

### Build for linux on windows
set GOOS=linux

Compile
---------
go build -ldflags="-s -w" -trimpath

Running Application
--------------------
./runb2bqueries -conf app.conf




