#!/usr/bin/env bash

# This script downloads the Flux OpenAPI schemas, then it validates the
# Flux custom resources and the kustomize overlays using kubeconform.
# This script is meant to be run locally and in CI before the changes
# are merged on the main branch that's synced by Flux.

# Copyright 2022 The Flux authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is meant to be run locally and in CI to validate the Kubernetes
# manifests (including Flux custom resources) before changes are merged into
# the branch synced by Flux in-cluster.

# Prerequisites
# - yq v4.30
# - kustomize v4.5
# - kubeconform v0.5.0

GIT_BASE_PATH=$(git rev-parse --show-toplevel)
SCRIPT_LIB_DIR="$GIT_BASE_PATH/.local/scripts/lib"
source "$SCRIPT_LIB_DIR/os.sh"

set -o errexit

# Download the Flux OpenAPI schemas
function download_flux_schemas() {
  pretty_print "\t${YELLOW}INFO - Downloading Flux OpenAPI schemas\n${NC}"
  mkdir -p /tmp/flux-crd-schemas/master-standalone-strict
  curl -sL https://github.com/fluxcd/flux2/releases/latest/download/crd-schemas.tar.gz | tar zxf - -C /tmp/flux-crd-schemas/master-standalone-strict 
  line_separator
}

# validate the Flux custom resources
function validate_flux_crds() {
  pretty_print "\t${YELLOW}INFO - Validating Flux custom resources\n${NC}"
  find . -type f -name '*.yaml' -print0 | while IFS= read -r -d $'\0' file;
    do
      pretty_print "\tINFO - Validating $file\n"
      yq e 'true' "$file" > /dev/null || result=1
  done
  line_separator
}

# kubeconform validate Kubernetes manifests against the OpenAPI schemas
function validate_kubeconform() {
  kubeconform_config=("-strict" "-ignore-missing-schemas" "-schema-location" "default" "-schema-location" "/tmp/flux-crd-schemas" "-verbose")
  pretty_print "\t${YELLOW}INFO - Validating clusters using kubeconform\n${NC}"
  find ./gitops/clusters -maxdepth 2 -type f -name '*.yaml' -print0 | while IFS= read -r -d $'\0' file;
    do
      kubeconform "${kubeconform_config[@]}" "${file}"
      if [[ ${PIPESTATUS[0]} != 0 ]]; then
        exit 1
      fi
  done
  line_separator
}

# validate the kustomize overlays
function validate_kustomize_overlays() {
  # mirror kustomize-controller build options
  kustomize_flags=("--load-restrictor=LoadRestrictionsNone")
  kustomize_config="kustomization.yaml"
  
  #find . -type f -name $kustomize_config -print0 | while IFS= read -r -d $'\0' file;
  warn "Skipping kustomize validation for .resources now"
  pretty_print "\t${YELLOW}INFO - Validating kustomize overlays \n${NC}"
  find . -type f -name kustomization.yaml | grep -v ".resources" | while IFS= read -r -d $'\0' file;
    do
      pretty_print "\t${YELLOW}INFO - Validating kustomization ${file/%$kustomize_config}\n${NC}"
      kustomize build "${file/%$kustomize_config}" "${kustomize_flags[@]}" | kubeconform "${kubeconform_config[@]}"
      if [[ ${PIPESTATUS[0]} != 0 ]]; then
        exit 1
      fi
  done
  line_separator
}

# main
function main() {
  download_flux_schemas
  validate_flux_crds
  validate_kubeconform
  validate_kustomize_overlays
}

#ToDo: FluxCD specific
#main $@





