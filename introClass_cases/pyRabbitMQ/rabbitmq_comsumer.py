import time

import pika
import threading
from grade.gen_cases import *


def consumer_callback(channel, method, properties, body):
    # projectName = body.split(",")[0]
    # message = body.split(",")[1]
    print("py recv msg:",body)
    # time.sleep(4)
    # py_mq_send("{}".format(time.time()))
    mission = threading.Thread(target=gen_loop,args=(chan,))
    mission.start()
    mission.join()


def py_mq_send(message):
    chan.basic_publish(exchange='',
                       routing_key='goQueue',
                       body=message
                       )
    print(f" python [x] Sent{message}")



credentials = pika.PlainCredentials("admin", "123456")
parameters = pika.ConnectionParameters(
    host="127.0.0.1",
    virtual_host='/',
    credentials=credentials
)

conn = pika.BlockingConnection(parameters)

chan = conn.channel()

chan.queue_declare(
    queue="goQueue",
    durable=True,
    auto_delete=False,
    exclusive=False,

)


chan.queue_declare(
    queue="pyQueue",
    durable=True,
    auto_delete=False,
    exclusive=False,

)

chan.basic_consume(
    queue="pyQueue",
    on_message_callback=consumer_callback,
    auto_ack=True)
print(' [*] Waiting for messages. To exit press CTRL+C')
chan.start_consuming()
