from Crypto.Cipher import DES
from time import time

# 1. 입력값 설정
plaintext = b"ABCDEFGH"  # 정확히 8바이트
if len(plaintext) != 8:
    raise ValueError("Plaintext must be 8 bytes")

key = bytes.fromhex("133457799BBCDFF1")  # 64비트 키
block = int.from_bytes(plaintext, byteorder='big')
print(f"Original plaintext (hex): {block:016X}")

# 2. 암호화 객체 생성 (ECB)
cipher = DES.new(key, DES.MODE_ECB)

# 3. 암호화
encrypted = cipher.encrypt(plaintext)
print(f"Encrypted ciphertext (hex): {encrypted.hex().upper()}")

# 4. 복호화
decrypted = cipher.decrypt(encrypted)
decrypted_block = int.from_bytes(decrypted, byteorder='big')
print(f"Decrypted result (hex): {decrypted_block:016X}")
print(f"Recovered text: {decrypted.decode()}")

# 5. 성능 측정
start_enc = time()
for _ in range(100):
    cipher.encrypt(plaintext)
elapsed_enc = time() - start_enc

start_dec = time()
for _ in range(100):
    cipher.decrypt(encrypted)
elapsed_dec = time() - start_dec

print(f"\nEncryption 100x time: {elapsed_enc:.6f}s")
print(f"Decryption 100x time: {elapsed_dec:.6f}s")
