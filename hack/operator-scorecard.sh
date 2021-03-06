#!/bin/bash
#
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

OPERATOR_SDK_CMD=$(command -v operator-sdk)
TREE_CMD=$(command -v tree)
SCRIPTS_HOME=$(dirname ${BASH_SOURCE})

if [ -z ${OPERATOR_SDK_CMD} ]; then
    echo "'operator-sdk' command not found, exiting."
    exit 1
fi    

source ${SCRIPTS_HOME}/latest_tag

# create a test deployment dir in the format expected by operator-sdk scorecard and copy artifacts there
TEST_DEPLOY_DIR=`mktemp -d 2>/dev/null || mktemp -d -t 'mytmpdir'`

echo "Copying artifacts in $TEST_DEPLOY_DIR ..."
mkdir -p ${TEST_DEPLOY_DIR}/deploy/crds
cp ${SCRIPTS_HOME}/../deploy/crds/ibmcloud.ibm.com_accessgroup_crd.yaml ${TEST_DEPLOY_DIR}/deploy/crds/ibmcloud.ibm.com_accessgroup_crd.yaml
cp ${SCRIPTS_HOME}/../deploy/crds/ibmcloud.ibm.com_customrole_crd.yaml ${TEST_DEPLOY_DIR}/deploy/crds/ibmcloud.ibm.com_customrole_crd.yaml
cp ${SCRIPTS_HOME}/../deploy/crds/ibmcloud.ibm.com_accesspolicy_crd.yaml ${TEST_DEPLOY_DIR}/deploy/crds/ibmcloud.ibm.com_accesspolicy_crd.yaml
cp ${SCRIPTS_HOME}/../deploy/examples/ibmcloud.ibm.com_accessgroup_cr.yaml ${TEST_DEPLOY_DIR}/deploy/crds/accessgroup_cr.yaml
cp ${SCRIPTS_HOME}/../deploy/examples/ibmcloud.ibm.com_customrole_cr.yaml ${TEST_DEPLOY_DIR}/deploy/crds/customrole_cr.yaml 
cp ${SCRIPTS_HOME}/../deploy/examples/ibmcloud.ibm.com_accesspolicy_cr.yaml ${TEST_DEPLOY_DIR}/deploy/crds/accesspolicy_cr.yaml 

# now copy all other artifacts
cp ${SCRIPTS_HOME}/../releases/latest/*_service_account.yaml ${TEST_DEPLOY_DIR}/deploy/service_account.yaml
cp ${SCRIPTS_HOME}/../releases/latest/*_role_binding.yaml ${TEST_DEPLOY_DIR}/deploy/role_binding.yaml
cp ${SCRIPTS_HOME}/../releases/latest/*_role.yaml ${TEST_DEPLOY_DIR}/deploy/role.yaml
cp ${SCRIPTS_HOME}/../releases/latest/*_operator.yaml ${TEST_DEPLOY_DIR}/deploy/operator.yaml
cp ${SCRIPTS_HOME}/../olm/${TAG}/ibmcloud_iam_operator.v${TAG}.clusterserviceversion.yaml ${TEST_DEPLOY_DIR}/deploy/ibmcloud_iam_operator.v${TAG}.clusterserviceversion.yaml

# get namespace to use
NS=$(cat ${SCRIPTS_HOME}/../releases/latest/*operator.yaml | grep namespace: | awk '{print $2}')

# create the scorecard config file
cat > ${TEST_DEPLOY_DIR}/.osdk-scorecard.yaml <<EOS
scorecard:
  # Setting a global scorecard option
  output: json
  plugins:
    # `basic` tests configured to test 3 CRs
    - basic:
        namespace: $NS
        cr-manifest:
          - "${TEST_DEPLOY_DIR}/deploy/crds/accessgroup_cr.yaml"
          - "${TEST_DEPLOY_DIR}/deploy/crds/customrole_cr.yaml"
          - "${TEST_DEPLOY_DIR}/deploy/crds/accesspolicy_cr.yaml"
        crds-dir: "${TEST_DEPLOY_DIR}/deploy/crds"   
    # `olm` tests configured to test 3 CRs
    - olm:
        cr-manifest:
          - "${TEST_DEPLOY_DIR}/deploy/crds/accessgroup_cr.yaml"
          - "${TEST_DEPLOY_DIR}/deploy/crds/customrole_cr.yaml"
          - "${TEST_DEPLOY_DIR}/deploy/crds/accesspolicy_cr.yaml"
        csv-path: "${TEST_DEPLOY_DIR}/deploy/ibmcloud_iam_operator.v${TAG}.clusterserviceversion.yaml"
EOS

# show up files
if [ -z ${TREE_CMD} ]; then
    echo "'tree' command not found, consider installing it."
else
    tree $TEST_DEPLOY_DIR
fi    

# show config file
echo "Scorecard config file:"
cat ${TEST_DEPLOY_DIR}/.osdk-scorecard.yaml

# create install ns if needed
kubectl get ns ${NS} >/dev/null 2>&1
if [ $? -eq 0 ]
then
    echo "Namespace ${NS} already exists"
else
    echo "Creating namespace ${NS}"
    kubectl create ns ${NS} 
fi

cd $TEST_DEPLOY_DIR
operator-sdk scorecard

# clean up
rm -r $TEST_DEPLOY_DIR
