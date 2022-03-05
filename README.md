# capturesql

git clone https://github.com/mft-labs/capturesql

### Fix (2022/03/05) if flag.Parse() added manually
git checkout main.go

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
go build -ldflags="-s -w" -trimpath -o RunDbReports

Running Application
--------------------
./RunDbReports -conf app.conf




