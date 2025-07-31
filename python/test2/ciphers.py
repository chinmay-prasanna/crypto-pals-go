from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from Crypto.Random import get_random_bytes
import base64


class ECB:
    def __init__(self, block_size=16, key=None):
        if not key:
            self.key = get_random_bytes(block_size)
        else:
            self.key = key
        self.block = AES.new(self.key, AES.MODE_ECB)
        self.block_size = block_size

    def encrypt(self, bstr):
        bstr = pad(bstr, self.block_size)
        output = b''
        for i in range(0, len(bstr), self.block_size):
            end = i+self.block_size
            encrypted = self.block.encrypt(bstr[i:end])
            output += encrypted

        return output
    
    def decrypt(self, bstr):
        output = b''
        for i in range(0, len(bstr), self.block_size):
            end = i+self.block_size
            decrypted = self.block.decrypt(bstr[i:end])
            output += decrypted

        return output
    

class CBC:
    def __init__(self, block_size=16, key=None, iv=None):
        if key:
            self.key = key
        else:
            self.key = get_random_bytes(block_size)
        if iv:
            self.iv = iv
        else:
            self.iv = get_random_bytes(block_size)
        self.block_size = block_size
        self.block = AES.new(self.key, AES.MODE_ECB)

    def xor(self, b1, b2):
        return bytes([x ^ y for x, y in zip(b1, b2)])
    
    def encrypt(self, bstr):
        bstr = pad(bstr, self.block_size)
        output = b''
        prev = self.iv
        ecb = ECB(block_size=self.block_size, key=self.key)
        for i in range(0, len(bstr), self.block_size):
            end = i+self.block_size
            e_str = self.xor(bstr[i:end], prev)
            encrypted = ecb.encrypt(e_str)
            prev = encrypted
            output += encrypted

        return output
    
    def decrypt(self, bstr):
        output = b''
        prev = self.iv
        ecb = ECB(block_size=self.block_size, key=self.key)
        for i in range(0, len(bstr), self.block_size):
            end = i+self.block_size
            decrypted = ecb.decrypt(bstr[i:end])
            d_str = self.xor(decrypted, prev)
            prev = bstr[i:end]
            output += d_str

        return output
