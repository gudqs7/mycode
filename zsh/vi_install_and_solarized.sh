sudo cd ~
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

