REGISTRY_DATA_PATH=/opt/registry
YOUR_IP=ip addr | grep 192.168  | awk 'NR==1{if ($1=="inet") print $2}' | awk -F/ 'NR==1{print $1}'
sudo docker pull registry.aliyuncs.com/gdk/registry 
sudo systemctl restart docker
docker tag registry.aliyuncs.com/gdk/registry registry
docker run -d --name registry -v $REGISTRY_DATA_PATH:/var/lib/registry -p 5000:5000 registry
echo -e "Now you can push and pull command to management images with registry.\nFor examples:\n  1.push images\ndocker pull daocloud.io/library/nginx:latest\ndocker tag daocloud.io/libary/nginx:latest $YOUR_IP:5000/nginx\ndocker push $YOUR_IP:5000/nginx\n  2.pull images \ndocker pull $YOUR_IP:5000/nginx"
