from ciphers import ECB
from Crypto.Util.Padding import pad, unpad

ecb = ECB()

def profile_for(email: str):
    email = email.replace("&", "")
    email = email.replace("=", "")
    email_str = f'email={email}&uid=10&role=user'
    email_bytes = bytes(email_str, 'utf-8')
    padded = pad(email_bytes, block_size=16)
    encoded = ecb.encrypt(padded)
    return encoded

encoded = profile_for("foo@bar.com")
padding_len = 16 - len('email=')
email1 = b'A'*padding_len + pad(b'admin', 16)
c1 = profile_for(email1.decode("latin1"))
a_block = c1[16:32]
email2 = "A"*13
c2 = profile_for(email2)
forged = c2[:32] + a_block

dc = ecb.decrypt(forged)
print(unpad(dc,16).decode())
