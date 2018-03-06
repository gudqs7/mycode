#-*- coding: utf-8 -*-，
#coding = utf-8
import re
from urllib.request import urlopen
import urllib
import pymysql
import uuid

def unescape(text):
    def fixup(m):
        text = m.group(0)
        if text[:2] == "&#":
            # character reference
            try:
                if text[:3] == "&#x":
                    return chr(int(text[3:-1], 16))
                else:
                    return chr(int(text[2:-1]))
            except ValueError:
                print("erreur de valeur")
                pass
        else:
            # named entity
            try:
                if text[1:-1] == "amp":
                    text = "&amp;amp;"
                elif text[1:-1] == "gt":
                    text = "&amp;gt;"
                elif text[1:-1] == "lt":
                    text = "&amp;lt;"
                else:
                    print(text[1:-1])
            except KeyError:
                print("keyerror")
                pass
        return text  # leave as is
    return re.sub("&#?\w+;", fixup, text)

def getinfo(content):
    array = content.split('\n')
    one = ''
    two = ''
    three=''
    parent =''
    for line in array :
        enil = line
        bc = enil.lstrip().split(' ')
        if len(bc)<=1:
            continue
        if line[:8] =='        ':
            parent = three
        elif line[:6] =='      ':
            three = bc[0]
            parent = two;
        elif line[:4] =='    ':
            two = bc[0]
            parent = one
        else :
            one = bc[0]
            parent = '';
        write(bc[0],bc[1],parent)


def resolve(url):
    url = 'http://www.tyut.edu.cn/xbsk/tougao/ztflh/'+url
    for data in urlopen(url):
        data = data.decode('gbk')
        clc = re.search('<pre class="style3">(.*?)</pre>', data,re.I|re.M|re.S)
        if clc!=None:
            content = clc.group(1)
            content = re.sub("<br>",'\n',content);
            getinfo(content)

def write(no,name,parent):
    db = pymysql.connect("localhost","root","guddqs","eip")
    db.set_charset('utf8')
    cursor = db.cursor()
    cursor.execute('SET NAMES utf8;')
    cursor.execute('SET CHARACTER SET utf8;')
    cursor.execute('SET character_set_connection=utf8;')
    sql = "INSERT INTO tlib_bookCategory(bookCategoryId,bookCategoryNo,bookCategoryName,parentId) VALUES ('"+uuid.uuid1().urn[9:]+"','"+no+"','"+name+"','"+parent+"')";
    try:
        cursor.execute(sql)
        print(no+' ok')
        db.commit()
    except :
        print(no+' shit')
        db.rollback()


for line in urlopen('http://www.tyut.edu.cn/xbsk/tougao/ztflh/'):
    line = line.decode('gb2312')
    back = re.search('<a href="(.*?)">',line,re.I|re.M|re.S)
    if back!=None:
        lujing = back.group(1)
        url = unescape(lujing)
        url = urllib.parse.quote(url)
        resolve(url)


