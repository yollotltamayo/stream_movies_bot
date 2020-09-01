import psycopg2

class db:
    def __init__(self,db):
        self.mydb = psycopg2.connect(
            host="localhost",
            password="yollotl",
            user="yollotl",
            database=db
        )

    def inser(self,val,sql):
        mycursor = self.mydb.cursor()
        mycursor.executemany(sql, val)
        self.mydb.commit()
        print(mycursor.rowcount, "was inserted.")

