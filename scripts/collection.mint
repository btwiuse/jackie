#!/usr/bin/env bash

set -uo pipefail

# cd $(dirname $(realpath $0))

dir=$(keypair.parse "$1")
if [[ $? -ne 0 ]]; then
  echo "$dir"
  exit 1
fi


# echo dir: $dir
# rm -r $dir
# exit 0

template="NFTSample"
cname="$2"
method="mint"
args="${3}"

ABI=$PWD/template/$template.abi

>$dir/log
if ! [[ -f $ABI ]]; then
  solc --base-path $PWD/template/node_modules/ template/$template.sol -o template --abi --overwrite >>$dir/log
fi

echo $ABI >> $dir/log
echo template=$template >> $dir/log
echo cname=$cname >> $dir/log
echo method=$method >> $dir/log
echo args="$args" >> $dir/log
# xchain-cli --keys $dir account new --account $account --fee 1000 > $dir/log

echo xchain-cli --keys $dir evm invoke --method "$method" -a "$args" "$cname" --abi "$ABI" --fee 1000 >> $dir/log
xchain-cli --keys $dir evm invoke --method "$method" -a "$args" "$cname" --abi "$ABI" --fee 1000 >> $dir/log
exit_code=$?
if [[ $exit_code -ne 0 ]]; then
	cat $dir/log | tail -n1
	exit $exit_code
fi

# xchain-cli evm query --method "$method" -a "$args" "$cname" --abi "$ABI" | sed -e 's,^contract response: ,,g'

cat<<EOF
response: $(cat $dir/log | grep '^contract response:' | sed 's,contract response: ,,g')
collection: "$cname"
template: $template
tx: "$(cat $dir/log | grep 'Tx id:' | sed 's,Tx id: ,,g')"
EOF

# rm -r $dir
