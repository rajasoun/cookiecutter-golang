#!/usr/bin/env bash 

MODE=${1:-"send"}
CLUSTER_NAME=${2:-"dev"}
TOPIC_NAME=${3:-"my"}

POD_PRODUCER="kafka-producer"
POD_CONSUMER="kafka-consumer"

# Delete pod if it exists
function delete_pod_if_exists(){
    local POD_NAME=$1
    if kubectl -n kafka get pod $POD_NAME &> /dev/null; then
        kubectl -n kafka delete pod $POD_NAME
    fi
}

# Send a message to the topic
function send_msg(){
    delete_pod_if_exists $POD_PRODUCER
    kubectl -n kafka run kafka-producer -ti --image=quay.io/strimzi/kafka:0.36.1-kafka-3.5.1 --rm=true --restart=Never -- bin/kafka-console-producer.sh --bootstrap-server $CLUSTER_NAME-cluster-kafka-bootstrap:9092 --topic $TOPIC_NAME-topic
}

# Consume messages from the topic
function consume_msg(){
    delete_pod_if_exists $POD_CONSUMER
    kubectl -n kafka run kafka-consumer -ti --image=quay.io/strimzi/kafka:0.36.1-kafka-3.5.1 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server $CLUSTER_NAME-cluster-kafka-bootstrap:9092 --topic $TOPIC_NAME-topic --from-beginning
}

# Depending on the mode, run one of the two functions
case "$MODE" in
  send)
    send_msg
    ;;
  consume)
    consume_msg
    ;;
  *)
    echo "Specify either 'send' or 'consume' as the first argument."
    exit 1
    ;;
esac
