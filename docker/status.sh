MY_DOCKER_PATH=/opt/eip-package/dockerServer
echo "tomocat1: "
ls $MY_DOCKER_PATH/tomcat1/webapps/ROOT

echo -e "\ntomcat2: "
ls $MY_DOCKER_PATH/tomcat2/webapps/ROOT

echo -e "\ncontainer info: "
docker ps -a | grep tomcat
