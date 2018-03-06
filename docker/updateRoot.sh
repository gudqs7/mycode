MY_DOCKER_PATH=/opt
NOW_PATH=pwd
echo $1
cp $1 $MY_DOCKER_PATH/tomcat1/webapps/ROOT/ROOT.war
cd /$MY_DOCKER_PATH/tomcat1/webapps/ROOT
jar xvf ROOT.war
rm -rf ROOT.war

cd $NOW_PATH

cp $1 $MY_DOCKER_PATH/tomcat2/webapps/ROOT/ROOT.war
cd $MY_DOCKER_PATH/tomcat2/webapps/ROOT
jar xvf ROOT.war
rm -rf ROOT.war

docker restart tomcat1 tomcat2
