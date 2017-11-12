FROM kong:0.11.1

# psqlコマンドインストール
RUN yum install -y postgresql

ADD ./init.sh /init.sh
CMD bash init.sh
