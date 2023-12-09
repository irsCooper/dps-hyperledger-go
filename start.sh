echo "[+] Starting project!

"

sleep 3;

cd ./test-network && bash startNetwork.sh

sleep 5;

echo "[+] Run web server"

cd ../app && go run ./cmd