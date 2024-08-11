res=$(curl -s "http://localhost:8080/charge")

if [ -z "$res" ]; then
    echo "Could not call API"
    exit 1
fi

echo "API response: $res"
exit 1
