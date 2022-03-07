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
go build -ldflags="-s -w" -trimpath -o runb2bqueries

Running Application
--------------------
#### Please fix b2bqueries.conf as desired and run the application
# sfghome variable has the base path where the SFG is installed and needs to be changed in b2bqueries.conf before running the below command in your b2b environment.
# These queries have been created for DB2, so some queries will not work in Oracle.

./runb2bqueries -conf b2bqueries.conf




