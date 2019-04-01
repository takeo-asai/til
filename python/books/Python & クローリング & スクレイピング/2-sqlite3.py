import sqlite3

conn = sqlite3.connect('top_cities.db')

c = conn.cursor()

# create empty table
c.execute('Drop table if exists cities')
c.execute('''
Create table cities (
    rank integer,
    city text,
    population integer
)
''')

# insert values
c.execute('Insert into cities values (?, ?, ?)', (1, '上海', 24150000))
c.execute('Insert into cities values (:rank, :city, :population)',
          {'rank': 2, 'city': 'カラチ', 'population': 23500000})
# ラベル付きで複数の insert は便利かもしれない
c.executemany('Insert into cities values (:rank, :city, :population)', [
    {'rank': 3, 'city': '北京', 'population': 21516000},
    {'rank': 4, 'city': '天津', 'population': 14722100},
    {'rank': 5, 'city': 'イスタンブール', 'population': 14160467}
])
conn.commit()  # autocommit ではない

# select
c.execute('select * from cities order by rank desc')
for row in c.fetchall():  # execute の返り値で受け取るのではなく、fetchall で受け取る
    print(row)

conn.close()
