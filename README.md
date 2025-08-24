# 20250824_database_manager
# shiojicloud_golag2024

cd /usr/local/vpnclient/

ps -ef | grep vpnc

'''
sudo apt-get -y update
'''


## dockerの起動

### dockerの確認
'''
sudo docker ps
'''

### コンテナの起動コマンド
'''
sudo docker stop go-http2api
sudo docker rm go-http2api
sudo docker build -t go-http2api .
sudo docker run -d -p 8000:8000 --name go-http2api go-http2api
sudo docker ps
'''

## コンテナのログをみる
sudo docker logs go-http2api

### アクセス先
```
https://shioji.cloud:8000/data
http://tk2-123-61726.vs.sakura.ne.jp:8000/data
http://tk2-123-61726.vs.sakura.ne.jp:8000/senddata
```

## エラー確認
'''
sudo docker ps -a
sudo docker logs dee21ae12a81
sudo docker logs 34e61caa8489

'''

mysqlのインストール
'''
sudo apt install mysql-server
'''

## dockerからmysqlに接続

IPアドレス確認
'''
hostname -I
'''

'''
// dsn := "chartier:ss642644@tcp(127.0.0.1:3306)/my_database"
dsn := "chartier:ss642644@tcp(153.121.37.165:3306)/my_database"
'''


外部アクセスの有効化
'''
GRANT ALL PRIVILEGES ON my_database.* TO 'chartier'@'%' IDENTIFIED BY 'ss642644';
FLUSH PRIVILEGES;
'''

## mysqlにログインする

pw: ss642644

'''
sudo mysql -u root -p
'''

'''
CREATE USER 'chartier'@'localhost' IDENTIFIED BY 'ss642644';
GRANT ALL PRIVILEGES ON cloud_database.* TO 'chartier'@'localhost';
FLUSH PRIVILEGES;
'''

- アクセス権の付与
```
GRANT ALL PRIVILEGES ON shioji_database.* TO 'chartier'@'%';
FLUSH PRIVILEGES;
```

- 権限確認

```
SHOW GRANTS FOR 'chartier'@'%';
```

## databaseの作成

'''
CREATE DATABASE IF NOT EXISTS my_database;

USE my_database;

CREATE TABLE IF NOT EXISTS api_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    time VARCHAR(255),
    lat DOUBLE,
    lon DOUBLE,
    psi DOUBLE,
    rot DOUBLE,
    sog DOUBLE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
'''

- 多種類データベース

'''
CREATE DATABASE IF NOT EXISTS shioji_database;

USE shioji_database;

CREATE TABLE IF NOT EXISTS api_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    time VARCHAR(255),
    GPS_time VARCHAR(255),
    lat DOUBLE,
    lon DOUBLE,
    head DOUBLE,
    cog DOUBLE,
    yawrate DOUBLE,
    roll DOUBLE,
    pitch DOUBLE,
    sog DOUBLE,
    log DOUBLE,
    u DOUBLE,
    v DOUBLE,
    cpp DOUBLE,
    o_cpp DOUBLE,
    rud DOUBLE,
    o_rud DOUBLE,
    windir DOUBLE,
    winspd DOUBLE,
    curdir DOUBLE,
    curspd DOUBLE,
    shaft_output DOUBLE,
    fuel_inlet DOUBLE,
    fuel_outlet DOUBLE,
    governor_cmd DOUBLE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
'''


## mysqlに権限の付与

- 現在の状態を確認

'''
SELECT host, user FROM mysql.user WHERE user = 'chartier';
'''

- 権限の削除

'''
DROP USER 'chartier'@'localhost';
DROP USER 'chartier'@'127.0.0.1';
'''

- 新しい権限の付与

'''
SHOW GRANTS FOR 'chartier'@'%';
'''

## 証明書の発行

'''
sudo apt-get update
sudo apt-get install certbot
'''

- 証明書の取得

'''
sudo certbot certonly --standalone -d shioji.cloud
'''

- port 80の一時停止

'''
sudo systemctl stop apache2
'''

- 自動更新設定

'''
sudo crontab -e
'''

- 設定の追加

'''
0 0 * * * /usr/bin/certbot renew --quiet
'''


### 認証ファイルのコピー

'''
sudo cp /etc/letsencrypt/live/shioji.cloud-0001/fullchain.pem /home/ubuntu/workspace/go/app/
sudo cp /etc/letsencrypt/live/shioji.cloud-0001/privkey.pem /home/ubuntu/workspace/go/app/
'''

- 権限の付与

'''
sudo chmod 600 /home/ubuntu/workspace/go/app/fullchain.pem
sudo chmod 600 /home/ubuntu/workspace/go/app/privkey.pem
'''

## gitのインストール

```
git add .
git commit -m "commit message"
git push origin main
```

## AIS用のSQL文

```
CREATE TABLE IF NOT EXISTS ais_logdata_20250318 (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    ParentDataID INT, -- logdata_XXXX の DataIDとリレーション
    Timestamp_JST VARCHAR(26),
    MMSI VARCHAR(20),
    NavigationalStatus VARCHAR(64),
    ROT VARCHAR(20),
    SOG VARCHAR(20),
    Longitude VARCHAR(32),
    Latitude VARCHAR(32),
    COG VARCHAR(20),
    Heading VARCHAR(20),
    TimeStamp VARCHAR(20),
    AISRawData TEXT,
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

```