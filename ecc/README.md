Potato Elliptic Curve Cryptography Wrapper
=========================================

This is a simple wrapper for `btcec`, that handles the specificities
of the format for keys in POC.

It was crafted in reference to `pcjs-ecc`, `pcjs-keygen` and the C++
codebase of Potato Coin Software.

This handles the `POC` prefix on public keys, manages the version and
checksums in public and private keys.
