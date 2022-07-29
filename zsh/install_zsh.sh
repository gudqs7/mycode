#!/bin/bash

release="centos"

if [ -f /etc/redhat-release ]; then
    release="centos"
elif cat /etc/issue | grep -Eqi "debian"; then
    release="debian"
elif cat /etc/issue | grep -Eqi "ubuntu"; then
    release="ubuntu"
elif cat /etc/issue | grep -Eqi "centos|red hat|redhat"; then
    release="centos"
elif cat /proc/version | grep -Eqi "debian"; then
    release="debian"
elif cat /proc/version | grep -Eqi "ubuntu"; then
    release="ubuntu"
elif cat /proc/version | grep -Eqi "centos|red hat|redhat"; then
    release="centos"
else
    release=""
fi


if [[ x"${release}" == x"centos" ]]; then
    mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
    curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

    sudo yum update -y
    echo '-----------------Update Finished. Start Install-----------------'

    yum install -y zsh git vim
    
elif [[ x"${release}" == x"debian" || x"${release}" == x"ubuntu" ]]; then
    sudo apt-get update -y
    echo '-----------------Update Finished. Start Install-----------------'

    sudo apt-get install -y zsh git vim
    
else
    v=`cat /proc/version`
    echo "Unknown Linux Release: ${release}"
    exit 1
fi

curl -L https://gitcode.net/mirrors/ohmyzsh/ohmyzsh/-/raw/master/tools/install.sh | sh

chsh -s /bin/zsh
zsh

home=`echo ~`
if [ -f /etc/issue ]; then
  sed -i 's/robbyrussell/af-magic/g' $home/.zshrc
else
  sed -i "" 's/robbyrussell/af-magic/g' $home/.zshrc
fi

source $home/.zshrc
