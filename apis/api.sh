while :
do
    res=`curl -i -X GET --url http://kong:8001/apis/ 2>&1 >/dev/null`
    if [ $? != 0 ]; then
      echo "start waiting..."
    else
      echo "starting confirm! Adding api!"
      #test用google
      curl -i -X POST --url http://kong:8001/apis/ --data 'name=google' --data 'uris=/google' --data 'upstream_url=https://www.google.co.jp' --data 'methods=GET'

      #jwt認証用ユーザ
      curl -X POST http://kong:8001/consumers     --data "username=gaku"     --data "custom_id=gaku"

      #ユーザ登録
      curl -i -X POST --url http://kong:8001/apis/ --data 'name=signup' --data 'uris=/users/signup' --data 'upstream_url=http://auth:8080/users/signup' --data 'methods=POST'
      #ログイン
      curl -i -X POST --url http://kong:8001/apis/ --data 'name=login' --data 'uris=/users/login' --data 'upstream_url=http://auth:8080/users/login' --data 'methods=POST'



      #パスワード変更
      curl -i -X POST --url http://kong:8001/apis/ --data 'name=changepw' --data 'uris=/change_password' --data 'upstream_url=http://auth:8080/users/change_password' --data 'methods=POST'
      curl -X POST http://kong:8001/apis/changepw/plugins     --data "name=jwt"
      break
    fi
    sleep 5;
done

echo 'Added APIs!'
