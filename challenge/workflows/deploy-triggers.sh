BUILD_CONF='workflows/cloudbuild_template.yaml'
REPO_NAME="challenge"
REPO_OWNER="vik-vok"

array=(
  'challenge-create':'ChallengeCreate'
)

for i in "${array[@]}"; do
  IFS=":"
  set -- ${i}

  CLOUD_FUNC_NAME=${1}
  GO_FUNC_NAME=${2}
  TRIGGER_NAME="${CLOUD_FUNC_NAME}-trigger"
  echo "#### Generating Trigger ${TRIGGER_NAME}"

#  gcloud alpha builds triggers delete "${TRIGGER_NAME}" --quiet
  gcloud beta builds triggers create github \
    --repo-name="${REPO_NAME}" \
    --repo-owner="${REPO_OWNER}" \
    --included-files="functions/${GO_FUNC_NAME}.go" \
    --name="${TRIGGER_NAME}" \
    --branch-pattern="^master$" \
    --build-config=${BUILD_CONF} \
    --substitutions _CLOUD_FUNC_NAME="${CLOUD_FUNC_NAME}",_GO_FUNC_NAME="${GO_FUNC_NAME}"
done
