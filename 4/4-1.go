package main
import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"encoding/hex"
	"fmt"
)
func hash160(data []byte) []byte {
	// Step 1: SHA-256
	sha256Hasher := sha256.New()
	sha256Hasher.Write(data)
	sha256Hash := sha256Hasher.Sum(nil)
	
	// Step 2: RIPEMD-160
	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(sha256Hash)
	return ripemd160Hasher.Sum(nil)
}
func generateAddress(publicKey string) (string, error) {
	// 将公钥的16进制字符串转换为字节
	pubKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	// Step 1: 生成指纹 (Hash160)
	fingerprint := hash160(pubKeyBytes)

	// Step 2: 添加版本字节 (0x6f) 到 Fingerprint 前面
	versionedPayload := append([]byte{0x6f}, fingerprint...)

	// Step 3: 生成校验码 (Checksum)，对 versionedPayload 做 HASH256
	checksum := hash256(versionedPayload)[:4]

	// Step 4: 将 versionedPayload 和 checksum 连接起来
	fullPayload := append(versionedPayload, checksum...)

	// Step 5: Base58 编码生成最终地址
	address := base58Encode(fullPayload)
	return address, nil
}
// Hash256: 两次SHA-256
func hash256(data []byte) []byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return hash[:]
}

// Base58 编码函数 (可用第三方包实现)
func base58Encode(input []byte) string {
	// 根据 Base58 编码规则将 input 编码为 Base58 字符串
	// 这里可以使用第三方 Base58 编码库，或自行实现
	// 为简洁起见，假设我们已导入 base58 包
	return base58.Encode(input)
}
func main() {
	publicKey1 := "02b1ebcdbac723f7444fdfb8e83b13bd14fe679c59673a519df6a1038c07b719c6"
	publicKey2 := "036e69a3e7c303935403d5b96c47b7c4fa8a80ca569735284a91d930f0f49afa86"

	address1, err := generateAddress(publicKey1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Address for Public Key 1:", address1)

	address2, err := generateAddress(publicKey2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Address for Public Key 2:", address2)
}
