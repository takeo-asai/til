from bs4 import BeautifulSoup

# ファイルから
with open("./dp.html") as f:
    soup = BeautifulSoup(f, "html.parser")
    print(soup.find_all("h1"))

# 直HTMLから
soup = BeautifulSoup('''
<html>
<head><title>this is from html</title></head>
</html>
''', "html.parser")
print(soup.find_all("title"))
