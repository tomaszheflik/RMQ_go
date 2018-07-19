#!/bin/bash

function get_cluster_info() {
    sleep 30
    FP=$(kubectl get pods -n test1 -l 'app=rabbitmq' -o jsonpath='{.items[0].metadata.name}')
    echo "[INFO] - Cluster Info"
    kubectl -n test1 exec -it $FP rabbitmqctl cluster_status
}
 
function deploy_rabbitmq() {
    echo "[INFO] - Deploy RabbitMQ"
    kubectl create -n $1 -f rabbit_deploy.yaml
}

function deploy_service() {
    echo "[INFO] - Deploy service"
    kubectl create -n $1 -f srv_deploy.yaml
}

function generate_deploy_cookie() {
    echo "[INFO] -  Generate cookie"
    echo $(openssl rand -base64 32) > erlang.cookie
    echo "[INFO] -  deploy to namespace: $1"
    kubectl delete namespace $1
    sleep 30
    kubectl create namespace $1 
    kubectl -n $1 create secret generic erlang.cookie --from-file=erlang.cookie
}
function clean_old_image() {
    echo "[INFO] - Clean old Docker Image"
    docker rmi -f minikube/rabbitmq-autocluster
}

function build_docker_image() {
    echo "[INFO] -  Build Docker image"
    docker build -t minikube/rabbitmq-autocluster .
}

function validate_input() {
    if [ -z "$1" ]; then
        echo "[ERROR] - No namespace"
        exit 1
    fi
}

function main() {
    validate_input $@
    clean_old_image
    build_docker_image
    generate_deploy_cookie $1
    deploy_service $@
    deploy_rabbitmq $@
    get_cluster_info $@
}

main $@