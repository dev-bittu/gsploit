echo "[+] Installing Required Packages"
apt install golang -y

echo "[+] Installing Required Modules"
go mod tidy

echo "[+] Building gsploit"
go build -o bin/gsploit cmd/gsploit/gsploit.go

echo "[+] Installing gsploit"
go install bin/gsploit

echo "[+] Sucessfully Install gsploit"

echo "[+] Adding GOPATH to PATH"
echo "export PATH=$PATH:$HOME/go/bin/" >> .bashrc

