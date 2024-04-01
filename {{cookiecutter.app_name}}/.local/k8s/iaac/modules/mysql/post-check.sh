#!/usr/bin/env bash

# Color variables
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
RED=$(tput setaf 1)
NC=$(tput sgr0) # No Color
BOLD=$'\033[1m'
BLUE=$'\e[34m'

NAMESPACE=${1:-"database"}
MYSQL_LABEL=${2:-"mysql"}
USERNAME=${3:-"app-user"}
PASSWORD=${4:-"password"} # Caution: Avoid hardcoding passwords in scripts

# Check if MySQL pod is running
function check_pod_running() {
    echo "${BOLD}${YELLOW}Checking if MySQL pod is running...${NC}"
    MYSQL_POD=$(kubectl get pods -n $NAMESPACE -l app=$MYSQL_LABEL -o jsonpath='{.items[0].metadata.name}')
    if [ -z "$MYSQL_POD" ]; then
        echo "${RED}MySQL pod not found! Status: Not Healthy${NC}"
        exit 1
    else
        echo "${GREEN}MySQL pod found: ${MYSQL_POD} Status: Healthy${NC}"
    fi
}

# Fetch logs of the MySQL pod
function fetch_pod_logs() {
    echo "${BOLD}${YELLOW}Fetching logs of the MySQL pod...${NC}"
    count=$(kubectl logs $MYSQL_POD -n $NAMESPACE | grep "MySQL Community Server - GPL" | grep "ready for connections" | wc -l)
    if [ $(echo $count ) -ne 2 ]; then
        echo "${RED}MySQL pod is not ready! Status: Not Healthy${NC}"
        exit 1
    else 
        echo "${GREEN}MySQL pod is ready! Status: Healthy${NC}"
    fi
}

# Access MySQL shell and create testdb
function access_mysql_shell_and_create_db() {
    echo "${BOLD}${YELLOW}Accessing MySQL shell and creating testdb...${NC}"
    kubectl exec -it $MYSQL_POD -n $NAMESPACE -- mysql -u$USERNAME -p$PASSWORD -e "CREATE DATABASE IF NOT EXISTS testdb; USE testdb; CREATE TABLE IF NOT EXISTS testtable (id INT, name VARCHAR(255)); INSERT INTO testtable (id, name) VALUES (1, 'test'); SELECT * FROM testtable;"
}

# Access MySQL shell and drop testdb
function access_mysql_shell_and_drop_db() {
    echo "${BOLD}${YELLOW}Accessing MySQL shell and dropping testdb...${NC}"
    kubectl exec -it $MYSQL_POD -n $NAMESPACE -- mysql -u$USERNAME -p$PASSWORD -e "DROP DATABASE IF EXISTS testdb;"
}

# Check service details
function check_service_details() {
    echo "${BOLD}${YELLOW}Checking service details...${NC}"
    kubectl get svc -n $NAMESPACE
}

# Check persistent storage
function check_persistent_storage() {
    echo "${BOLD}${YELLOW}Checking PV and PVC...${NC}"
    kubectl get pv,pvc -n $NAMESPACE
}

# Check resource usage
function check_resource_usage() {
    echo "${BOLD}${YELLOW}Checking resource usage for MySQL pod...${NC}"
    kubectl top pod $MYSQL_POD -n $NAMESPACE
}

# Main
function main() {
    check_pod_running
    fetch_pod_logs
    access_mysql_shell_and_create_db
    access_mysql_shell_and_drop_db
    check_service_details
    check_persistent_storage
    check_resource_usage
    echo "${GREEN}Done.${NC}"
}

main $@
