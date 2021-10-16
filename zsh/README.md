# 安装 zsh 并安装 oh-my-zsh , mac自带 zsh 则仅安装 oh-my-zsh; 并切换默认主题为 af-magic
```
curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/zsh/install_zsh.sh | bash
```

# 安装 zsh-high-light-syntax; 并尝试自动配置到 .zshrc
```
curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/zsh/zsh_syntax_highlighting.sh | bash
```

# 安装 vim 配色, 注意是 vim , centos 自带的 vi 似乎不管用?
```
curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/zsh/vi_install_and_solarized.sh | bash
```

# 一键上面三个
```
sudo ls
curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/zsh/one_zsh.sh | bash
```

# GIT 缩写别名
```
curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/git/git-alias.sh | bash
```


# GIT 升级版本
```
curl -L https://raw.githubusercontent.com/gudqs7/mycode/master/zsh/update_git.sh | bash
```

# 设置 PROXY
```bash
alias proxy="export ALL_PROXY=http://127.0.0.1:10809"
alias unproxy="unset ALL_PROXY"
alias ip="curl cip.cc"
```
