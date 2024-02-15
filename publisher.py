import paho.mqtt.client as mqtt
import time
import json
import random

# Configuração do cliente MQTT
client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION1, client_id="")

# Conexão ao broker MQTT local na porta 1891
client.connect("localhost", 1891, 60)

# Carrega os dados JSON do arquivo
with open('dados.json', 'r') as file:
    data = json.load(file)

# Loop para publicar mensagens continuamente
try:
    while True:
        # modified_data = modify_random_value(data.copy())
        keys = list(data.keys())
        amount = len(data.keys())
        index = random.randint(0, amount - 2 + 1 )
        # Serializa os dados JSON modificados para uma string
        message = json.dumps(data[keys[index]])
        # Publica a mensagem no tópico "test/topic"
        client.publish("test/topic", message)
        print(f"Publicado: {message}")
        # Intervalo de 2 segundos entre publicações
        time.sleep(2)
except KeyboardInterrupt:
    # Mensagem exibida ao interromper o script manualmente
    print("Publicação encerrada")

# Desconecta do broker MQTT
client.disconnect()
