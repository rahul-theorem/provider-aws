#!/usr/bin/env bash

set -eux

# setting up colors
BLU='\033[0;34m'
YLW='\033[0;33m'
GRN='\033[0;32m'
RED='\033[0;31m'
NOC='\033[0m' # No Color
echo_info(){
    printf "\n${BLU}%s${NOC}" "$1"
}
echo_step(){
    printf "\n${BLU}>>>>>>> %s${NOC}\n" "$1"
}
echo_sub_step(){
    printf "\n${BLU}>>> %s${NOC}\n" "$1"
}

echo_step_completed(){
    printf "${GRN} [✔]${NOC}"
}

echo_success(){
    printf "\n${GRN}%s${NOC}\n" "$1"
}
echo_warn(){
    printf "\n${YLW}%s${NOC}" "$1"
}
echo_error(){
    printf "\n${RED}%s${NOC}" "$1"
    exit 1
}

# ------------------------------

# get the build environment variables from the special build.vars target in the main makefile
eval $(make --no-print-directory -C $(pwd) build.vars)

# ------------------------------

NAMESPACE="upbound-system"
K8S_CLUSTER="uptest"

# setup package cache
echo_step "setting up local package cache"
CACHE_PATH="${WORK_DIR}/package-cache"
mkdir -p "${CACHE_PATH}"
echo "created cache dir at ${CACHE_PATH}"

# ${PLATFORM} should be available here but for some reason it is not and all CI
# runners are linux_amd64 anyway.
"${UP}" xpkg xp-extract --from-xpkg "${OUTPUT_DIR}/xpkg/linux_amd64/${PROJECT_NAME}-${VERSION}.xpkg" -o "${CACHE_PATH}/${PROJECT_NAME}.gz" && chmod 644 "${CACHE_PATH}/${PROJECT_NAME}.gz"

# create kind cluster with extra mounts
echo_step "creating k8s cluster using kind"
KIND_CONFIG="$( cat <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
  - hostPath: "${CACHE_PATH}/"
    containerPath: /cache
EOF
)"
echo "${KIND_CONFIG}" | "${KIND}" create cluster --name="${K8S_CLUSTER}" --wait=5m --config=-

# tag controller image and load it into kind cluster
BUILD_REGISTRY="build-$(echo "${HOSTNAME}"-"$(pwd)" | shasum -a 256 | cut -c1-8)"

BUILD_IMAGE="${BUILD_REGISTRY}/${PROJECT_NAME}-${SAFEHOSTARCH}"
PACKAGE_IMAGE="${PROJECT_NAME}:${VERSION}"

docker tag "${BUILD_IMAGE}" "${PACKAGE_IMAGE}"
"${KIND}" load docker-image "${PACKAGE_IMAGE}" --name="${K8S_CLUSTER}"

echo_step "create ${NAMESPACE} namespace"
"${KUBECTL}" create ns ${NAMESPACE}

echo_step "create persistent volume and claim for mounting package-cache"
PV_YAML="$( cat <<EOF
apiVersion: v1
kind: PersistentVolume
metadata:
  name: package-cache
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 5Mi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/cache"
EOF
)"
echo "${PV_YAML}" | "${KUBECTL}" create -f -

PVC_YAML="$( cat <<EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: package-cache
  namespace: ${NAMESPACE}
spec:
  accessModes:
    - ReadWriteOnce
  volumeName: package-cache
  storageClassName: manual
  resources:
    requests:
      storage: 1Mi
EOF
)"
echo "${PVC_YAML}" | "${KUBECTL}" create -f -

# install universal-crossplane
echo_step "installing universal-crossplane from stable channel"
"${HELM3}" repo add upbound-stable https://charts.upbound.io/stable
chart_version="$("${HELM3}" search repo upbound-stable/universal-crossplane --devel  | awk 'FNR == 2 {print $2}')"
echo_info "using universal-crossplane version ${chart_version}"
echo
# we replace empty dir with our PVC so that the /cache dir in the kind node
# container is exposed to the crossplane pod
"${HELM3}" install uxp --namespace ${NAMESPACE} upbound-stable/universal-crossplane --devel --version ${chart_version} --wait --set packageCache.pvc=package-cache

# ----------- integration tests
echo_step "--- INTEGRATION TESTS ---"

# install package
echo_step "installing ${PROJECT_NAME} into \"${NAMESPACE}\" namespace"

CONFIG_YAML="$( cat <<EOF
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: config
spec:
  image: "${PACKAGE_IMAGE}"
  args: ["-d"]
EOF
)"

INSTALL_YAML="$( cat <<EOF
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: "${PROJECT_NAME}"
spec:
  package: "${PROJECT_NAME}"
  packagePullPolicy: Never
  controllerConfigRef:
    name: config
EOF
)"

echo "${CONFIG_YAML}" | "${KUBECTL}" apply -f -
echo "${INSTALL_YAML}" | "${KUBECTL}" apply -f -

# printing the cache dir contents can be useful for troubleshooting failures
echo_step "check kind node cache dir contents"
docker exec "${K8S_CLUSTER}-control-plane" ls -la /cache

echo_step "waiting for provider to be installed"

kubectl wait "provider.pkg.crossplane.io/${PROJECT_NAME}" --for=condition=healthy --timeout=180s

kubectl get deployment -n ${NAMESPACE}
