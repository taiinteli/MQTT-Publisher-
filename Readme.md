# Simulador de dispositivos IoT

Este script Python utiliza a biblioteca Paho MQTT para conectar-se a um broker MQTT local e publicar mensagens JSON aleatórias. Os dados são carregados de um arquivo `dados.json` e, a cada iteração, um campo específico dos dados é alterado aleatoriamente antes da publicação.

- Vídeo demonstrativo de funcionamento: 
- Python 3.x
- Paho MQTT (`pip install paho-mqtt`)
- Arquivo `dados.json` no mesmo diretório do script

## Funcionamento

1. O script carrega o JSON do arquivo `dados.json`.
2. Link para os dados utilizados: https://sigmasensors.com.br/produtos/sensor-de-radiacao-solar-sem-fio-hobonet-rxw-lib-900 
3. Em um loop infinito, seleciona aleatoriamente um campo para modificar e publica os dados no tópico "test/topic".
4. A publicação ocorre a cada 2 segundos, até que o script seja interrompido manualmente.

## Uso

1. Certifique-se de que o broker MQTT esteja rodando na porta 1891.
2. Coloque o arquivo `dados.json` no mesmo diretório do script.
3. Execute o script: `python script.py`.
4. Interrompa com `Ctrl+C` para encerrar a publicação.

