# study-microservices
マイクロサービス勉強用

# api登録(golangとkongの連携)

    curl -i -X POST --url http://localhost:8001/apis/ --data 'name=gorilla' --data 'uris=/gorilla' --data 'upstream_url=http://auth:8080/' --data 'methods=POST'

# 確認

    curl -X POST http://localhost:8000/gorilla


