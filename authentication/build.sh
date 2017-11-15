cd /go/src/authentication
go build .

if ! psql -q -U postgres -h auth-db -lqt | cut -d \| -f 1 | grep -wq auth_db; then
  echo "Databaseが存在しません。作成を行います。"
  psql -q -U postgres -h auth-db -c "create database auth_db;"
fi

echo "Database OK"

#起動
./authentication
