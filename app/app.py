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

	channel.queue_declare(queue=os.environ['SCRIBAM_QUEUE_NAME'], durable=True)
	
	message = json.dumps({'application': 'flask','user':'arthur','data': 'test with flask 2','dateTime': '2020-03-10','trackID':'123'})

	channel.basic_publish(exchange='', routing_key=os.environ['SCRIBAM_QUEUE_NAME'], body=message)
	print(" [x] Sent Message")
	connection.close()
	return 'Ok'

if __name__ == '__main__':
   app.run(host="0.0.0.0", debug=True)