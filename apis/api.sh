while :
do
    res=`curl -i -X GET --url http://kong:8001/apis/ 2>&1 >/dev/null`
    if [ $? != 0 ]; then
      echo "start waiting..."
    else
      echo "starting confirm! Adding api!"
      curl -i -X POST --url http://kong:8001/apis/ --data 'name=google' --data 'uris=/google' --data 'upstream_url=https://www.google.co.jp' --data 'methods=GET'
      break
    fi
    sleep 5;
done

echo 'Added APIs!'
