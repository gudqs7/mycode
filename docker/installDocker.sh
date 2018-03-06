yum -y install docker
tee /etc/docker/daemon.json <<-'EOF'
{
  "insecure-registries":["192.168.1.252:5000"],
  "registry-mirrors":["https://apg1cbea.mirror.aliyuncs.com"],
  "graph":"/opt/eip-package/docker"
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker

