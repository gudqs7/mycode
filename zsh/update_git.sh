mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

yum install curl-devel expat-devel gettext-devel openssl-devel zlib-devel asciidoc
yum install  gcc perl-ExtUtils-MakeMaker

git --version

yum remove git

mkdir -p /usr/local/src
cd /usr/local/src/
curl -o git-2.15.1.tar.xz https://www.kernel.org/pub/software/scm/git/git-2.15.1.tar.xz
tar -vxf git-2.15.1.tar.xz
cd git-2.15.1

make prefix=/usr/local/git all
make prefix=/usr/local/git install

echo "export PATH=$PATH:/usr/local/git/bin" >> /etc/profile
source /etc/profile

git --version