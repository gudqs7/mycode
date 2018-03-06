registryIp=192.168.1.252:5000
docker tag $registryIp/nginx nginx 
docker tag $registryIp/redis redis
docker tag $registryIp/tomcat:v2 mytomcat
#docker tag $registryIp/mysql mysql 

