srcbasepath=/opt/eip-sourcecode
sourcepath=${srcbasepath}/aolong-eip
tomcatpath=/opt/eip-package/dockerServer
projectname=gzievs
projectsourcepath=${sourcepath}/project/${projectname}/trunk

svn update ${sourcepath}
svn update ${srcbasepath}/aolong-lessonClass
svn update ${srcbasepath}/aolong-cms

chmod 755 ${sourcepath}/*.sh
chmod 755 ${srcbasepath}/aolong-lessonClass/*.sh
chmod 755 ${srcbasepath}/aolong-cms/*.sh

cd ${sourcepath}
${sourcepath}/install-module.sh

cd ${srcbasepath}/aolong-lessonClass/code/trunk
mvn clean install -Pinstall-module -Dmaven.test.skip=true

cd ${srcbasepath}/aolong-cms/code/trunk
mvn clean install -Pinstall-module -Dmaven.test.skip=true

cd ${projectsourcepath}
mvn clean package -Pal25 -Dmaven.test.skip=true

cp ${projectsourcepath}/target/aolong-eip-${projectname}.war ${tomcatpath}/tomcat1/webapps/ROOT/${projectname}.war
cd ${tomcatpath}/tomcat1/webapps/ROOT
jar -xf ${projectname}.war 

cp ${projectsourcepath}/target/aolong-eip-${projectname}.war ${tomcatpath}/tomcat2/webapps/ROOT/${projectname}.war
cd ${tomcatpath}/tomcat2/webapps/ROOT
jar -xf ${projectname}.war 

cp ${projectsourcepath}/target/aolong-eip-${projectname}.war ${tomcatpath}/tomcat3/webapps/ROOT/${projectname}.war
cd ${tomcatpath}/tomcat3/webapps/ROOT
jar -xf ${projectname}.war 

#cp ${projectsourcepath}/target/aolong-eip-${projectname}.war ${tomcatpath}/tomcat4/webapps/ROOT/${projectname}.war
#cd ${tomcatpath}/tomcat4/webapps/ROOT
#jar -xf ${projectname}.war 

docker restart tomcat1 tomcat2
#sh ${tomcatpath}/tomcat3/bin/shutdown.sh
#sh ${tomcatpath}/tomcat3/bin/start.sh

echo ----------------------------------
echo system restart done !!!
echo ----------------------------------

