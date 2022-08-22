#/bin/bash
export FOO="CIAO"
OUT=$(go run main.go $1)
echo "======"
echo $OUT
echo "======"
if [ "$OUT" != "" ]; then
  eval $OUT
else
  echo "Nessuna selezione"
fi


