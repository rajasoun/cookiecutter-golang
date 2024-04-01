# Execution Sequence 


kubectl apply -k .local/k8s/iaac/modules/cert-manager
kubectl wait -n cert-manager --for=condition=ready pod --field-selector=status.phase=Running --timeout=120s
.local/scripts/wrapper.sh run wait_for_cert_manager_crds

kubectl apply -k .local/k8s/iaac/modules/opentelemetry-operator
kubectl wait -n opentelemetry-operator-system --for=condition=ready pod --field-selector=status.phase=Running --timeout=120s
.local/scripts/wrapper.sh run wait_for_open_telemetry_crds

.local/scripts/wrapper.sh run create_secrets_in_cluster_from_aws_credential_file
.local/scripts/wrapper.sh run check_aws_token_expiration_time
kubectl apply -k .local/k8s/iaac/modules/aws-otel-collector
kubectl wait -n aws-otel-collector --for=condition=ready pod --field-selector=status.phase=Running --timeout=120s


kubectl apply -k .local/k8s/iaac/modules/ingress-nginx
kubectl wait -n ingress-nginx --for=condition=ready pod --field-selector=status.phase=Running --timeout=120s


.local/scripts/wrapper.sh run delete_aws_credentials_from_cluster
