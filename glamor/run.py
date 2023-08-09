from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.chrome.options import Options
import sys
import time
import json
from bs4 import BeautifulSoup

route = sys.argv[1]

options = Options()
browser = webdriver.Chrome(options=options)
browser.get(route)

time.sleep(6)

rendered_html = browser.page_source
browser.quit()

soup = BeautifulSoup(rendered_html, 'html.parser')

for script in soup.find_all('script'):
    script.extract()

soup = BeautifulSoup(str(soup), 'html.parser')
formatted_html = soup.prettify()
with open("glamor_layout.html", "w", encoding="utf-8") as f:
    f.write(formatted_html)

