until psql -h "kong-database" -U "postgres" -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
>&2 echo "Postgres is up - executing command"

echo "DB OK"

#migration
kong migrations up

#start
kong start
