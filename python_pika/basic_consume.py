import pika


if __name__ == '__main__':

    auth = pika.PlainCredentials('guest', 'guest')
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost', port=5672))
    channel = connection.channel()

    channel.queue_declare(queue='TEST01')

    def callback(ch, method, properties, body):
        print(" [x] Received %r" % body)

    # basic_consume 定义以怎样的形式开始消费
    channel.basic_consume(on_message_callback=callback,
                          queue='TEST01',
                          auto_ack=True)
    print(' [*] Waiting for messages. To exit press CTRL+C')

    # 阻塞函数，每消费一条消息就执行一下回调函数 on_message_callback=callback
    channel.start_consuming()





