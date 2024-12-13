import mysql.connector as con 
from dotenv import load_dotenv
import os
import json
from pathlib import Path

load_dotenv()

db = con.connect(host = "localhost" ,user = "root" , password = os.getenv("mysqlPassword") )
cursor = db.cursor()
cursor.execute("use kumar;")

cwd = Path(__file__).parent

backlog = {}
addr = cwd / "../../webScraping/scrape2Backlog.json"
with open(addr) as filer:
    backlog = json.load(filer)

content = {}
with open(cwd / "../../webScraping/main_scraped_2.json") as filer:
    content = json.load(filer)

cursor.execute("create table backlog2 ( id int auto_increment primary key , category varchar(50) , topic varchar(250) , link varchar (250) )")

for category , topics in backlog.items():
    for topic,link in topics.items():
        query = "insert into backlog2 (category,topic,link) values (%s,%s,%s);"
        data = (category,topic,link)
        cursor.execute(query,data)


db.commit()

cursor.execute("create table scraped_data_2_content (id int auto_increment  primary key,category varchar(50) , topic varchar(250) , content mediumtext);")
cursor.execute("create table scraped_data_2_comment (id int auto_increment  primary key, topic varchar(250) , comment mediumtext);")

for category , topics in content.items():
    for topic,posts in topics.items():
        query = "insert into scraped_data_2_content (category,topic,content) values (%s,%s,%s);"
        data = (category,topic,posts["content"])
        cursor.execute(query,data)
        for comment in posts["comments"]:
            query = "insert into scraped_data_2_comment (topic,comment) values (%s,%s);"
            data = (topic,comment)
            cursor.execute(query,data)
db.commit()


