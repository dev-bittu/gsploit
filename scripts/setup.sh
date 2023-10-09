echo "[+] Installing Required Packages"
sudo apt install golang -y

echo "[+] Installing Required Modules"
go mod tidy

echo "[+] Building gsploit"
go build -o bin/gsploit cmd/gsploit/gsploit.go

echo "[+] Installing gsploit"
go install bin/gsploit

echo "[+] Sucessfully Install gsploit"
echo "[!] Please make sure your GOPATH is in your PATH"

