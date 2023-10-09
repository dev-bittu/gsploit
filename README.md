# Gsploit
Gsploit is an exploit framework written in Go (Golang).
It allows users to easily exploit vulnerabilities, create and encode payloads, and post exploits to networks. 
It offers a wide range of exploits to effectively penetrate target systems and access sensitive data. 

Additionally, Gsploit provides auxiliary modules such as port scanners and fingerprinting tools for gathering information about target systems. 
Users can create custom payloads using techniques like buffer overflows and shellcode injection, and the framework includes advanced encoding techniques to evade antivirus software and other security measures.
Gsploit also offers post-exploitation modules for maintaining access to compromised systems, including tools for persistence and privilege escalation. 

Overall, Gsploit is a highly capable framework for security professionals and penetration testers seeking to exploit vulnerabilities or gather information about potential attack vectors.

# Features
- Auxiliar systems
- Exploit vulnerabilities
- Create payloads
- Encode payloads
- Post exploits to networks

# Installation
1. Install Golang and Git:
```bash
sudo apt install golang git -y
```
2. Clone gsploit git repo:
```bash
git clone https://github.com/dev-bittu/gsploit.git
```
3. Go to gsploit:
```bash
cd gsploit
```
4. Setup:
- In Linux
```bash
sh scripts/setup.sh
```
- In Termux:
```bash
sh scripts/setup-in-termux.sh
```
- In windows
```bash
./scripts/setup.bat
```

## Usage
Gsploit provides a command-line interface for interacting with its various tools.
To see a list of available commands,
run ``gsconsole --help.`` OR ``./bin/gsconsole``.

## Contributing
Contributions to Gsploit are welcome and encouraged!
If you would like to contribute, please fork the repository and submit a pull request.

## Author
- dev-bittu

### Contact With Us
  - [GitHub](https://github.com/dev-bittu "dev-bittu")

## License
Gsploit is released under the [MIT License](LICENSE "License File").
See the LICENSE file for more details.
