yum install zsh git -y
wget https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O - | sh
chsh -s /bin/zsh
home=`echo ~`
sed -i "" 's/robbyrussell/af-magic/g' $home/.zshrc 
