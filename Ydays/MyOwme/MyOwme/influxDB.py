
#pip influx
from influxdb import InfluxDBClient

from datetime import datetime
 #db
client = InfluxDBClient(host='localhost', port=8086, username='myuser', password='mypass', ssl=True, verify_ssl=True)
client.create_database('mydb')
client.get_list_database()
client.switch_database('mydb')


json_payload = []
data = {
    "measurement": "stocks",
    "tags" : {
        "ticker" : "TSLA"
    },
    "time" : datetime.now(),
    "fields" : {
        "open" : 688.37, # A changer avec le db final
        "close" : 667.93
    }
}

json_payload.append(data)

client.write_points(json_payload)

