from selenium import webdriver
from time import sleep
from selenium.webdriver.common.keys import Keys
import urllib.request
import random
from manage_tables import db
import re
def scrap():
    bd = db("catalog")
    sql = "INSERT INTO netflix_en (name,year,time,genre,original) VALUES (%s ,%s, %s, %s, %s)"
    driver = webdriver.Chrome()
    driver.get("https://www.finder.com/netflix-movies")
    val = []
    sleep(4)
    table = driver.find_element_by_tag_name("table")
    tbody = table.find_element_by_tag_name("tbody")
    #for a in range(186):
    for row in table.find_elements_by_tag_name("tr"):
        tup = []
        cta = 0
        ban = 0
        td  = row.find_elements_by_tag_name("td")
        for a in td:
            if cta == 0 or cta == 3 :
                original = re.split(r'\n',a.text)
                tup.append(original[0].lower() ,)
                if len(original) == 2:
                    ban = 1
            else:
                if cta == 2 or cta == 1:
                    tup.append(int(a.text))
            cta+=1
        if(len(tup) != 0):
            tup.append('false' if ban == 1 else 'true')
            val.append(tuple(tup))
    print(val)
    bd.inser(val , sql)
    sleep(random.randint(5,10))
    driver.find_element_by_xpath('//button[text()="Next"]').click()
    driver.close()
scrap()

