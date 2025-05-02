from datetime import datetime


from influxdb_client import InfluxDBClient, Point, WritePrecision
from influxdb_client.client.write_api import SYNCHRONOUS

# You can generate an API token from the "API Tokens Tab" in the UI
#token = #a remplir avec le token une fois creer
org = "deoko@outlook.fr"
bucket = "bucktest"

with InfluxDBClient(url="https://us-east-1-1.aws.cloud2.influxdata.com", token=token, org=org) as client:
    write_api = client.write_api(write_options=SYNCHRONOUS)

data = "mem,host=host1 used_percent=23.43234543"
write_api.write(bucket, org, data)

point = Point("mem") \
  .tag("host", "host1") \
  .field("used_percent", 23.43234543) \
  .time(datetime.utcnow(), WritePrecision.NS)

write_api.write(bucket, org, point)

sequence = ["mem,host=host1 used_percent=23.43234543","mem,host=host1 available_percent=15.856523"]
write_api.write(bucket, org, sequence)

query = 'from(bucket: "bucktest") |> range(start: -1h)'
tables = client.query_api().query(query, org=org)
for table in tables:
    for record in table.records:
        print(record)
        client.close()