package main
import (
	"fmt"
	"os"
	"strings"
	"github.com/ecies/go"
)

func genPrivKey() *eciesgo.PrivateKey {
	var input string
	fmt.Println("do you have private key ? (y/n): ")
	fmt.Scan(&input)

	var priv *eciesgo.PrivateKey

	if input != "y" {
		priv, _ = eciesgo.GenerateKey()
		os.WriteFile("private.key", []byte(priv.Hex()), 0644)
		fmt.Println("Private key generated\nTake it from private.key and keep it in secret\ndelete private.key after copying")
	} else {
		fmt.Println("enter private key: ")
		fmt.Scan(&input)
		priv, _ = eciesgo.NewPrivateKeyFromHex(input)
		fmt.Println("Private key loaded")
	}
	return priv
}

func crypt(name string) {
	priv := genPrivKey()
	pub := priv.PublicKey
	data, err := os.ReadFile(name)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	enc, _ := eciesgo.Encrypt(pub, data)
	os.WriteFile(strings.Split(name, ".")[0] + ".vlt", enc, 0644)
	fmt.Println(name + " encrypted")
}

func decrypt(name string, format string) {
	var priv *eciesgo.PrivateKey
	var input string
	fmt.Println("enter private key: ")
	fmt.Scan(&input)
	priv, _ = eciesgo.NewPrivateKeyFromHex(input)
	data, err1 := os.ReadFile(name)

	if err1 != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	dec, err := eciesgo.Decrypt(priv, data)

	if err != nil {
		fmt.Println("Error decrypting file")
		os.Exit(1)
	}

	os.WriteFile(strings.Split(name, ".")[0] + "." + format, dec, 0644)
	fmt.Println(name + " decrypted")

}

func main() {
	if os.Args[1] == "crypt" {
		crypt(os.Args[2])
	} else if os.Args[1] == "decrypt" {
		decrypt(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Usage: go run main.go crypt <file>|decrypt <file> [format]")
	}
}