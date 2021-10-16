A="$1"
B="$2"

echo "输入的原始值：A=$A,B=$B"

#判断字符串是否相等
if [ x"$A" = x"$B" ];then
echo "[ = ]"
fi

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

if [[ $release == "centos" ]]
then
    echo '-----------------Update Finished. Start Install-----------------'
elif [[ $release == "debian" || $release == "ubuntu" ]]
then     
    echo '-----------------Update Finished. Start Install-----------------'
else
    v=`cat /proc/version`
    echo "Unknown Linux Release: ${v}"
fi