#echo "create mysql"
MY_DOCKER_PATH=/opt/eip-package/dockerServer
TOMCAT1_PORT=8081
TOMCAT2_PORT=8082
NGINX_PORT=80
REDIS_PORT=6379
MYSQL_PORT=3306
MYSQL_PASSWORD=aolong123
#docker run -p $MYSQL_PORT:3306 --privileged=true --name mysql -v $MY_DOCKER_PATH/mysql/conf/my.cnf:/etc/mysql/my.cnf -v $MY_DOCKER_PATH/mysql/data:/mysql_data -e MYSQL_ROOT_PASSWORD=$MYSQL_PASSWORD -d mysql

echo "create redis: "

docker run -p $REDIS_PORT:6379 --privileged=true --name redis -v $MY_DOCKER_PATH/redis/data:/data  -d redis redis-server --appendonly yes

echo "create tomcat1: "

docker run -d --name tomcat1 --privileged=true -p $TOMCAT1_PORT:8080 -v $MY_DOCKER_PATH/tomcat1/conf/:/opt/tomcat/conf -v $MY_DOCKER_PATH/tomcat1/logs/:/opt/tomcat/logs -v $MY_DOCKER_PATH/tomcat1/webapps:/opt/tomcat/webapps  mytomcat

echo "create tomcat2: "

docker run -d --name tomcat2 --privileged=true -p $TOMCAT2_PORT:8080 -v $MY_DOCKER_PATH/tomcat2/conf/:/opt/tomcat/conf -v $MY_DOCKER_PATH/tomcat2/logs/:/opt/tomcat/logs -v $MY_DOCKER_PATH/tomcat2/webapps:/opt/tomcat/webapps  mytomcat

echo "create nginx: "

docker run -p $NGINX_PORT:80 --privileged=true --name nginx -v $MY_DOCKER_PATH/nginx/www:/etc/nignx/www -v $MY_DOCKER_PATH/nginx/nginx.conf:/etc/nginx/nginx.conf -v $MY_DOCKER_PATH/nginx/logs:/wwwlogs  -d nginx

echo "All done"
