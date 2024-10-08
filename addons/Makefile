# Copyright 2020 The Kubermatic Kubernetes Platform contributors.
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

DOCKER_REPO ?= "quay.io/kubermatic"

.PHONY: docker
docker:
	docker build -t $(DOCKER_REPO)/addons:$(TAG) .

.PHONY: aws-node-termination-handler
aws-node-termination-handler: OUTPUT_FILE=aws-node-termination-handler/aws-node-termination-handler.yaml
aws-node-termination-handler:
	helm repo add eks https://aws.github.io/eks-charts/ || true
	helm repo update
	cat aws-node-termination-handler/_header.txt > $(OUTPUT_FILE)
	helm template aws-node-termination-handler eks/aws-node-termination-handler \
	  --version 0.21.0 \
	  --namespace kube-system \
	  --values values-aws-node-termination-handler.yaml \
	  >> $(OUTPUT_FILE)
	cat aws-node-termination-handler/_footer.txt >> $(OUTPUT_FILE)
	./templatify-images.sh $(OUTPUT_FILE)

.PHONY: aws-ebs-csi-driver
aws-ebs-csi-driver: OUTPUT_FILE=csi/aws-ebs/driver.yaml
aws-ebs-csi-driver:
	helm repo add aws-ebs-csi-driver https://kubernetes-sigs.github.io/aws-ebs-csi-driver || true
	helm repo update
	mkdir -p csi/aws-ebs
	cat csi/aws-ebs/_header.txt > $(OUTPUT_FILE)
	helm --namespace kube-system template aws-ebs-csi-driver aws-ebs-csi-driver/aws-ebs-csi-driver \
	  --version 2.26.0 \
	  --set 'controller.k8sTagClusterId=\{{ "{{ .Cluster.Name }}" }}' \
	  --set 'node.securityContext.seccompProfile.type=RuntimeDefault' \
	  --set 'controller.securityContext.seccompProfile.type=RuntimeDefault' \
	  --set 'controller.podAnnotations.cluster-autoscaler\.kubernetes\.io/safe-to-evict-local-volumes=socket-dir' \
	  --api-versions 'policy/v1/PodDisruptionBudget' \
	  --skip-tests \
	  >> $(OUTPUT_FILE)
	cat csi/aws-ebs/_footer.txt >> $(OUTPUT_FILE)
	./templatify-images.sh $(OUTPUT_FILE)

.PHONY: gcp-csi-driver
gcp-csi-driver: OUTPUT_FILE=csi/gcp/driver.yaml
gcp-csi-driver:
	mkdir -p csi/gcp
	cp gcp-csi-driver-kustomization.yaml csi/gcp/kustomization.yaml
	cat csi/gcp/_header.txt > $(OUTPUT_FILE)
	kubectl kustomize csi/gcp/ >> $(OUTPUT_FILE)
	cat csi/gcp/_footer.txt >> $(OUTPUT_FILE)
	rm csi/gcp/kustomization.yaml
	./templatify-images.sh $(OUTPUT_FILE)

.PHONY: nutanix-csi
nutanix-csi: nutanix-csi-storage nutanix-csi-snapshotter

.PHONY: nutanix-csi-storage
nutanix-csi-storage: OUTPUT_FILE=csi/nutanix/csi-driver.yaml
nutanix-csi-storage:
	helm repo add nutanix https://nutanix.github.io/helm/ || true
	helm repo update
	mkdir -p csi/nutanix
	cat csi/nutanix/_header.txt > $(OUTPUT_FILE)
	helm --namespace kube-system template nutanix-csi-storage nutanix/nutanix-csi-storage \
	  --version 2.6.6 \
	  --set 'createSecret=false' \
	  >> $(OUTPUT_FILE)
	cat csi/nutanix/_footer.txt >> $(OUTPUT_FILE)
	./templatify-images.sh $(OUTPUT_FILE)

nutanix-csi-snapshotter: OUTPUT_FILE=csi/nutanix/snapshot-controller.yaml
nutanix-csi-snapshotter:
	helm repo add nutanix https://nutanix.github.io/helm/ || true
	helm repo update
	mkdir -p csi/nutanix
	cat csi/nutanix/_header.txt > $(OUTPUT_FILE)
	helm --namespace kube-system template nutanix-csi-snapshot nutanix/nutanix-csi-snapshot \
	  --version 6.3.2 \
	  >> $(OUTPUT_FILE)
	cat csi/nutanix/_footer.txt >> $(OUTPUT_FILE)
	./templatify-images.sh $(OUTPUT_FILE)

.PHONY: metallb
metallb: OUTPUT_FILE=metallb/00_metallb.yaml
metallb:
	cat metallb/_header.txt > $(OUTPUT_FILE)
	wget -q -O- https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml >> $(OUTPUT_FILE)
	./templatify-images.sh $(OUTPUT_FILE)
