#!/usr/bin/env bash

# Color variables
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
NC=$(tput sgr0) # No Color
BOLD=$'\033[1m'
BLUE=$'\e[34m'

# Retrieves pod information including names, container names, and ports
function get_pod_details() {
  local namespace=${1:-default}
  kubectl get pods --namespace "$namespace" -o jsonpath='{range .items[*]}'"${GREEN}Pod Name: ${NC}"'{.metadata.name}{"\n"}{range .spec.containers[*]}'"${YELLOW} Container Name: ${NC}"'{.name}{"\n"}'"${YELLOW} Port(s): ${NC}"'{range .ports[*]}{.containerPort}{","}{end}{"\n"}{end}{"\n"}{end}' 
}

# Retrieves service information including local endpoint (DNS name and port)
function get_service_details() {
  local namespace=${1:-default}
  local service_names=($(kubectl get services --namespace "$namespace" -o jsonpath='{range .items[*].metadata.name}{@}{"\n"}{end}'))
  printf "${BOLD}Services:${NC}\n"
  for service_name in ${service_names[@]}; do
    local service_info=$(kubectl get service "$service_name" --namespace "$namespace" -o jsonpath='{.metadata.name}.{.metadata.namespace}.svc.cluster.local:{range .spec.ports[*]}{.port}{";"}{end}{"\n"}')
    printf "${BLUE}%s${NC} ${YELLOW}%s${NC}\n" "$service_name" "$service_info"
  done
}

# Print adminer URL from ingress service
function print_adminer_url() {
  local namespace=${1:-default}
  local ingress_host=$(kubectl get ingress adminer-ingress -n database -o jsonpath='{.spec.rules[0].host}')

  local query_params="?server=mysql&username=app-user&db=mysql&select=user"
  local full_url="http://${ingress_host}${query_params}"

  printf "\n${GREEN}URL: ${NC}" 
  printf "\nAdminer ${BLUE}%s\n" "$full_url${NC}"

}



# Main function that orchestrates the execution flow
function main() {
  local namespace=${1:-default}
  printf "\n${GREEN}Namespace: %s${NC}\n" "$namespace"
  get_pod_details "$namespace"
  get_service_details "$namespace"
  print_adminer_url "$namespace"
}

main "$@"
