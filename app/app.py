from flask import Flask
import pika
import os
import json

app = Flask(__name__)

@app.route('/')
def create_message():
	connection = pika.BlockingConnection(
    pika.URLParameters(os.environ['SCRIBAM_QUEUE']))
	channel = connection.channel()

	channel.queue_declare(queue='log_message', durable=True)
	
	message = json.dumps({'source': 'flask','message': 'test with flask 2','dateTime': '2020-03-10','type':'info'})

	channel.basic_publish(exchange='', routing_key='log_message', body=message)
	print(" [x] Sent Message")
	connection.close()
	return 'Ok'

if __name__ == '__main__':
   app.run(host="0.0.0.0", debug=True)