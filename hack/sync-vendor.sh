#!/bin/bash

set -e

tmpDir="${tmpDir:-/tmp/sync-vendor}"
fromRepoName="${fromRepoName}"
fromRepo="${fromRepo}"
fromBranch="${fromBranch}"
fromDir="${fromDir}"
toRepoName="${toRepoName}"
toRepo="${toRepo}"
toBranch="${toBranch}"

fromRepoLocation="${tmpDir}/${fromRepoName}"
toRepoLocation="${tmpDir}/${toRepoName}"

mkdir -p "${fromRepoLocation}"
mkdir -p "${toRepoLocation}"

cd "${fromRepoLocation}"
# if we already have a git repo, assume it's the right one and get the requested branch
git rev-parse --is-inside-work-tree || git clone --single-branch --branch=${fromBranch} -- "${fromRepo}" ${fromRepoLocation}
git checkout ${fromBranch}
git reset --hard origin/${fromBranch}
firstFromSHA=$(git log --reverse --oneline  --format='%H' | head -n 1)
newFromSHA=$(git log --oneline --format='%H' -1)

cd "${toRepoLocation}"
# if we aren't in a git repo, clone it
git rev-parse --is-inside-work-tree || git clone --single-branch --branch=${toBranch} -- "${toRepo}" ${toRepoLocation}
git fetch --refmap -- origin refs/heads/${toBranch}:refs/remotes/origin/${toBranch}
git checkout ${toBranch}
git reset --hard origin/${toBranch}
startingFromSHA=${firstFromSHA}
if [ -f "${fromRepoName}.sha" ]; then
    startingFromSHA=$(cat ${fromRepoName}.sha)
fi
if [ "${newFromSHA}" == "${startingFromSHA}" ]; then
    echo "already at level: ${newFromSHA}, nothing to sync"
    exit 0
fi
echo "syncing from ${startingFromSHA}..${newFromSHA}"
# create a clean branch to start from
git branch -D "${toBranch}-working" || true
git checkout -b "${toBranch}-working"

pushd "${fromRepoLocation}"
patchDir="${tmpDir}/patches"
rm -rf ${patchDir} && mkdir -p "${patchDir}"
index=0
for commit in $(git log --format='%H' --no-merges --reverse -- "${startingFromSHA}..${newFromSHA}" "${fromDir}"); do
  git format-patch --raw --start-number=${index} --relative="${fromDir}" "${commit}^..${commit}" -o "${patchDir}"
  let index+=10
done
# remove all commits that had no entries
find "${patchDir}" -type f -size 0 -exec rm {} \;
popd

# only patch if we have patches
if [ "$(ls -A ${patchDir})" ]; then
    # apply the changes
    if ! git am -3 --ignore-whitespace ${patchDir}/*.patch; then
      echo 2>&1
      echo "++ Patches do not apply cleanly, continue with 'git am' in ${relativedir}" 2>&1
      exit 1
    fi
    echo 2>&1
    echo "++ All patches applied cleanly upstream" 2>&1
else
    echo "no patches to apply.  bumping sync commit."
fi

echo "${newFromSHA}" > ${fromRepoName}.sha
git add ${fromRepoName}.sha && git commit -m "sync(${fromRepo}): ${newFromSHA}"

# push the results
git push origin ${toBranch}-working:${toBranch}