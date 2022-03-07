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
#### Please fix b2bqueries.conf as desired and run the application


./runb2bqueries -conf b2bqueries.conf




