
import pika


if __name__ == '__main__':

    auth = pika.PlainCredentials('guest', 'guest')
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost', port=5672))
    channel = connection.channel()
    # 声明一个
    channel.queue_declare(queue='TEST01')

    channel.basic_publish(exchange='',
                          routing_key='TEST01',
                          body='Hello World!')
    print(" [x] Sent 'Hello World!'")
    connection.close()




