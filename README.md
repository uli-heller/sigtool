[![GoDoc](https://godoc.org/github.com/opencoff/go-sign?status.svg)](https://godoc.org/github.com/opencoff/go-sign)

# README for sigtool


## What is this?
`sigtool` is an opinionated tool to generate keys, sign, verify, encrypt &
decrypt files using Ed25519 signature scheme.  In many ways, it is like
like OpenBSD's [signify][1] -- except written in Golang and definitely
easier to use. It can use SSH ed25519 public and private keys.

It can sign and verify very large files - it prehashes the files
with SHA-512 and then signs the SHA-512 checksum. The keys and signatures
are human readable YAML files.

It can encrypt files for multiple recipients - each of whom is identified
by their Ed25519 public key. The encryption generates ephmeral
Curve25519 keys and creates pair-wise shared secret for each recipient of
the encrypted file. The caller can optionally use a specific private key
during the encryption process - this has the benefit of also authenticating
the sender (and the receiver can verify the sender if they possess the
corresponding sender's public key).

The sign, verify, encrypt, decrypt operations can use OpenSSH Ed25519 keys
*or* the keys generated by sigtool. This means, you can send encrypted
files to any recipient identified by their comment in `~/.ssh/authorized_keys`.

## How do I build it?
You need two things:

1. Protobuf compiler:

   On Debian based systems: `apt install protobuf-compiler`

   Consult your OS's package manager to install protobuf tools;
   these are typically named 'protobuf' or 'protoc'.

2. go 1.13+ toolchain


Next, build sigtool:

    git clone https://github.com/opencoff/sigtool
    cd sigtool
    make

The binary will be in `./bin/$HOSTOS-$ARCH/sigtool`.
where `$HOSTOS` is the host OS where you are building (e.g., openbsd)
and `$ARCH` is the CPU architecture (e.g., amd64).

## How do I use it?
Broadly, the tool can:

- generate new key pairs (public key and private key)
- sign a file
- verify a file against its signature
- encrypt a file
- decrypt a file

### Generate Key pair
To start with, you generate a new key pair (a public key used for
verification and a private key used for signing). e.g.,

    sigtool gen /tmp/testkey

The tool then generates */tmp/testkey.pub* and */tmp/testkey.key*.  The secret
key (".key") can optionally be encrypted with a user supplied pass
phrase - which the user has to enter via interactive prompt:

    sigtool gen -p /tmp/testkey

### Sign a file
Signing a file requires the user to provide a previously generated
Ed25519 private key.  The signature (YAML) is written to STDOUT.
e.g.,  to sign `archive.tar.gz` with private key `/tmp/testkey.key`:

    sigtool sign /tmp/testkey.key archive.tar.gz

If *testkey.key* was encrypted without a user pass phrase:

    sigtool sign --no-password /tmp/testkey.key archive.tar.gz


The signature can also be written directly to a user supplied output
file.

    sigtool sign -o archive.sig /tmp/testkey.key archive.tar.gz


### Verify a signature against a file
Verifying a signature of a file requires the user to supply three
pieces of information:

- the Ed25519 public key to be used for verification
- the Ed25519 signature
- the file whose signature must be verified

e.g., to verify the signature of *archive.tar.gz* against
*testkey.pub* using the signature *archive.sig*

    sigtool verify /tmp/testkey.pub archive.sig archive.tar.gz


You can also pass a public key as a string (instead of a file):

    sigtool verify iF84Dymq/bAEnUMK6DRIHWAQDRD8FwDDDfsgFfzdjWM= archive.sig archive.tar.gz

Note that signing and verifying can also work with OpenSSH ed25519
keys.

### Encrypt a file by authenticating the sender
If the sender wishes to prove to the recipient that they  encrypted
a file:

    sigtool encrypt -s sender.key to.pub -o archive.tar.gz.enc archive.tar.gz


This will create an encrypted file *archive.tar.gz.enc* such that the
recipient in possession of *to.key* can decrypt it. Furthermore, if
the recipient has *sender.pub*, they can verify that the sender is indeed
who they expect.

### Decrypt a file and verify the sender
If the receiver has the public key of the sender, they can verify that
they indeed sent the file by cryptographically checking the output:

    sigtool decrypt -o archive.tar.gz -v sender.pub to.key archive.tar.gz.enc

Note that the verification is optional and if the `-v` option is not
used, then decryption will proceed without verifying the sender.

### Encrypt a file *without* authenticating the sender
`sigtool` can generate ephemeral keys for encrypting a file such that
the receiver doesn't need to authenticate the sender:

    sigtool encrypt to.pub -o archive.tar.gz.enc archive.tar.gz

This will create an encrypted file *archive.tar.gz.enc* such that the
recipient in possession of *to.key* can decrypt it.

### Encrypt a file to an OpenSSH recipient *without* authenticating the sender
Suppose you want to send an encrypted file where the recipient's
public key is in `~/.ssh/authorized_keys`. Such a recipient is identified
by their OpenSSH key comment (typically `name@domain`):

    sigtool encrypt user@domain -o archive.tar.gz.enc archive.tar.gz

If you have their public key in file "name-domain.pub", you can do:

    sigtool encrypt name-domain.pub -o archive.tar.gz.enc archive.tar.gz

This will create an encrypted file *archive.tar.gz.enc* such that the
recipient can decrypt using their private key.

## Technical Details

### How is the file encryption done?
The file encryption uses AES-GCM-256 in AEAD mode. The encryption uses
a random 32-byte AES-256 key. This root key is expanded via
HKDF-SHA256 into:

   - AES-GCM-256 key (32 bytes)
   - AES Nonce (12 bytes)
   - HMAC-SHA-256 key (32 bytes)

The input to the HKDF is the root-key, header-checksum ("salt") and
a context string.

The input is broken into chunks and each chunk is individually AEAD encrypted.
The default chunk size is 4MB (4 * 1048576 bytes). Each chunk generates
its own nonce: the top-4 bytes of the nonce is the chunk-number. The
actual chunk-length and EOF marker is used as additional data (the
"AD" of "AEAD").
The last block has its most-signficant-bit set to 1 to denote EOF. Thus, the
maximum chunk size is set to 1GB.

We calculate a running hmac of the plaintext blocks; when sender
identity is present, the final HMAC is signed via the sender's
Ed25519 key. This signature is appended as the "trailer" (last 64
bytes of the encrypted file are the Ed25519 signature).

When sender identity is not present, the last bytes are random
bytes.

### What is the public-key cryptography in sigtool?
`sigtool` uses ephemeral Curve25519 keys to generate shared secrets
between pairs of sender & one or more recipients. This pairwise shared
secret is used as a key-encryption-key (KEK) to wrap the
data-encryption key in AEAD mode. Thus, each recipient has their own
individual encrypted key blob - that **only** they can decrypt.

If the sender authenticates the encryption by providing their secret
key, the encryption key material is signed via Ed25519 and the signature
is encrypted (using the data-encryption key) and stored in the
header. If the sender opts to not authenticate, a "signature" of all
zeroes is encrypted instead.

The Ed25519 keys generated by `sigtool` or OpenSSH are transformed to their
corresponding Curve25519 points in order to generate the pair-wise shared secret.
This elliptic co-ordinate transform follows [FiloSottile's writeup][2].

### Format of the Encrypted File
Every encrypted file starts with a header and the header-checksum:

* Fixed-size header
* Variable-length header
* SHA256 sum of both of the above

The fixed length header is:

    7 byte magic ("SigTool")
    1 byte version number
    4 byte header length (big endian encoding)

The variable length header has the per-recipient wrapped keys. This is
described as a protobuf file (sign/hdr.proto):

```protobuf
    message header {
        uint32 chunk_size = 1;
        bytes  salt       = 2;
        bytes  pk         = 3;  // sender's ephemeral curve PK
        bytes  sender     = 4;  // ed25519 signature of key material
        repeated wrapped_key keys = 5;
    }

    /*
     * A file encryption key is wrapped by a recipient specific public
     * key. WrappedKey describes such a wrapped key.
     */
    message wrapped_key {
        bytes d_key = 1;
        bytes nonce = 2;
    }
```

The SHA256 sum covers the fixed-length and variable-length headers.

The encrypted data immediately follows the headers above. Each encrypted
chunk is encoded the same way:

```C
    4 byte chunk length (big endian encoding)
    AEAD encrypted chunk data
    AEAD tag
```

The chunk length does _not_ include the AEAD tag length; it is implicitly
computed. The chunk data and AEAD tag are treated as an atomic unit for AEAD
decryption.

### How is the private key protected?
The Ed25519 private key is encrypted in AES-GCM-256 mode using a key
derived from the user's pass-phrase. The user pass phrase is expanded via
SHA256; this expanded pass phrase is fed to `scrypt()` to
generate a key-encryption-key.  In pseudo code, this operation looks
like below:

    passphrase = get_user_passphrase()
    expanded   = SHA512(passphrase)
    salt       = randombytes(32)
    key        = Scrypt(expanded, salt, N, r, p)
    esk        = AES256_GCM(ed25519_private_key, key)

Where, ```N```, ```r```, ```p``` are Scrypt parameters. In our
implementation:

    N = 2^19 (1 << 19)
    r = 8
    p = 1


## Understanding the Code
The core logic is in `src/sign`: it is a library that exposes all the
functionality: key generation, key parsing, signing, encryption, decryption
etc.

* `src/encrypt.go` contains the core encryption, decryption code
* `src/sign.go`    contains the Ed25519 signing, verification code
* `src/keys.go`    contains key generation, serialization, de-serialization
* `src/ssh.go`     contains code to parse SSH Ed25519 key files
* `src/stream.go`  contains code that provides an `io.Reader` and `io.WriteCloser` interface
                   for encryption and decryption.
* `tests.sh`       simple round trip test using the tool; this is in addition to the tests in
                   `sign/`.


The generated keys and signatures are proper YAML files and human
readable.

The signature file contains a hash of the public key - so that at
verification time, the right private key may be used (in situations
where there are lots of keys).

Signatures on large files are calculated efficiently by reading them
in memory mapped mode (```mmap(2)```) and hashing the file contents
using SHA-512. The Ed25519 signature is calculated on the file-hash.

### Tests
The core library in `sign/` has extensive tests to verify signing and encryption.
Additionally, a simple shell script `tests.sh` does a full roundtrip of tests
using `sigtool`.

## Example of Keys, Signature

### Ed25519 Public Key
A serialized Ed25519 public key looks like so:

    pk: uxpDh+gqXojAmxA/6vxZHzA+Uk+8wogUwvEhPBlWgvo=

### Ed25519 Private Key
And, a serialized Ed25519 private key looks like so:

```yaml

    esk: t3vfqHbgUiA733KKPymFjWT8DdnBEkiMfsDHolPUdQWpvVn/F1Z4J6KYV3M5rGO9xgKxh5RAmqt+6LKgOiJAMQ==
    salt: pPHKG55UJYtJ5wU0G9hBvNQJ0DvT0a7T4Fmj4aPB84s=
    algo: scrypt-sha256
    Z: 131072
    r: 16
    p: 1
```


### Ed25519 Signature
A generated signature looks like below after serialization:

```yaml

    comment: inpfile=/tmp/file.txt
    pkhash: 36z9tCwTIVNwwDlExrB0SQ==
    signature: ow2oBP+buDbEvlNakOrsxgB5Yc/7PYyPVZCkfyu7oahw8BakF4Qf32uswPaKGZ8RVz4uXboYHdZtfrEjCgP/Cg==
```

Here, ```pkhash`` is a SHA256 of the public key needed to verify
this signature.

## Licensing Terms
The tool and code is licensed under the terms of the
GNU Public License v2.0 (strictly v2.0). If you need a commercial
license or a different license, please get in touch with me.

See the file ``LICENSE.md`` for the full terms of the license.

## Author
Sudhi Herle <sw@herle.net>

[1]: https://www.openbsd.org/papers/bsdcan-signify.html
[2]: https://blog.filippo.io/using-ed25519-keys-for-encryption/
