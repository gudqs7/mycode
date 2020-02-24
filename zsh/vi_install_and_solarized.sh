git clone https://github.com/altercation/vim-colors-solarized.git
if [ ! -d ~/.vim/colors ]; then
  mkdir -p ~/.vim/colors
fi
cp vim-colors-solarized/colors/solarized.vim ~/.vim/colors
tee ~/.vimrc <<-'EOF'
set nu
syntax enable
set background=dark
colorscheme solarized
EOF

sudo mv /bin/vi /bin/vibak
sudo ln -s /bin/vim /bin/vi

rm -rf vim-colors-solarized