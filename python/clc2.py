#-*- coding: utf-8 -*-，
#coding = utf-8
import re
from urllib import request
from bs4 import BeautifulSoup
import urllib
import pymysql
import uuid
import _thread
import threading

global count

def set_value(value):
    # 告诉编译器我在这个方法中使用的a是刚才定义的全局变量a,而不是方法内部的局部变量.
    global a
    a = value

def get_value():
    # 同样告诉编译器我在这个方法中使用的a是刚才定义的全局变量a,并返回全局变量a,而不是方法内部的局部变量.
    global a
    return a

def resolve(url,begin):
    try:
        threadLock.acquire();
        count = get_value();
        proxys = [{'http':'114.112.104.223:80'},{'http':'119.28.50.37:82'},{'http':'39.134.10.27:8080'},{'http':'39.134.10.21:8080'},{'http':'39.134.10.2:8080'},{'http':'39.134.10.10:8080'},{'http':'39.134.10.16:8080'},{'http':'39.134.10.22:8080'},{'http':'39.134.10.11:8080'},{'http':'39.134.10.100:8080'},{'http':'39.134.10.18:8080'},{'http':'39.134.10.101:8080'},{'http':'39.134.10.98:8080'},{'http':'39.134.10.20:8080'},{'http':'39.134.10.3:8080'},{'http':'39.134.10.102:8080'},{'http':'39.134.10.99:8080'}]
        print(count)
        if(count>=len(proxys)):
            count = 0
        proxy = proxys[count]
        set_value(count+1);
        
        #创建ProxyHandler
        proxy_support = request.ProxyHandler(proxy)
        #创建Opener
        opener = request.build_opener(proxy_support)
        #添加User Angent
        opener.addheaders = [('User-Agent','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36')]
        #安装OPener
        request.install_opener(opener)
        threadLock.release();
        
        parent = ''
        urlArr= url.split('/');
        if(url[len(url)-1:]=='/'):
            parent = urlArr[len(urlArr)-2]
        else:
            parent = urlArr[len(urlArr)-1]
        data = request.urlopen(url)
        data = data.read()
        data = data.decode('utf-8')
        soup = BeautifulSoup(data,'lxml')
        for tr in soup.table.tbody.find_all('tr'):
            tr = BeautifulSoup(str(tr),'lxml')
            if(len(tr.find_all('td'))>1) :               
                clc = tr.find_all('td')[1]
                name = tr.find_all('td')[2].a
                href = name['href']
                hrefArr = href.split('/')
                nextParent = hrefArr[len(hrefArr)-2]
                if(parent=='category' and nextParent!=begin):
                    continue
    #            _thread.start_new_thread(write,(str.strip(clc.string),str.strip(name.string),parent))
                write(str.strip(clc.string),str.strip(name.string),parent)
                _thread.start_new_thread(resolve,('http://www.clcindex.com'+name['href'],begin))
                #resolve('http://www.clcindex.com'+name['href'],begin)
    except Exception as e :
        print('Reason:', e) 
        pass
            

def write(no,name,parent):
    db = pymysql.connect("localhost","root","aolong","eip")
    db.set_charset('utf8')
    cursor = db.cursor()
    cursor.execute('SET NAMES utf8;')
    cursor.execute('SET CHARACTER SET utf8;')
    cursor.execute('SET character_set_connection=utf8;')
    sql = "INSERT INTO tlib_bookCategory(bookCategoryId,bookCategoryNo,bookCategoryName,parentId) VALUES ('"+no+"','"+no+"','"+name+"','"+parent+"')";
    try:
        cursor.execute(sql)
        print(no+' ok')
        db.commit()
    except :
        #print(no+' shit')
        db.rollback()


threadLock = threading.Lock()
url = 'http://www.clcindex.com/category/'
set_value(0)
#_thread.start_new_thread(resolve,(url,'T'))
_thread.start_new_thread(resolve,(url,'X'))
#_thread.start_new_thread(resolve,(url,'V'))
_thread.start_new_thread(resolve,(url,'Z'))
#resolve(url,'T')
while 1:
    pass


