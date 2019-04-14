import sys
import time

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait


def navigate_to_top(driver):
    driver.get('https://note.mu/')


def get_posts(driver):
    posts = []
    for card in driver.find_elements_by_css_selector('.p-cardItem'):
        posts.append({
            'url': card.find_element_by_tag_name('a').get_attribute('href'),
            'title': card.find_element_by_tag_name('h3').text,
        })
    return posts


def main():
    driver = webdriver.PhantomJS()
    driver.set_window_size(1024, 1024)
    navigate_to_top(driver)
    posts = get_posts(driver)
    # スクロールは JS を Inject するが、うまく画面が動いていないので、、、
    # Angular のページをそのまま selenium で扱うのは難しい、、
    for post in posts:
        print(post)


if __name__ == '__main__':
    main()
