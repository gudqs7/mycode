# 参考自博客: https://www.cnblogs.com/iois/p/11665825.html
# 效果为 红色包围的当前目录 + 蓝色包围的 git 当前的分支名称
branch=`git branch 2>/dev/null | grep '*' | sed 's/* \(.*\)/(\1)/'`
BLUE="\[\e[0;36m\]"
RED="\[\e[0;33m\]"
NC="\[\e[0m\]"
export PS1="${RED}\w${NC} ${BLUE}${branch}${NC} \$ "
# \u 就是当前工作目录的意思  \$ 是根据用户情况显示 $(普通用户) 或 #(root时)
export PS1='\[\e[0;31;07m\]\u\[\e[0;07m\]@\[\e[0;32;07m\]\h\[\e[0m\]:\[\e[0;36m\]\w\[\e[0m\]\$ '
