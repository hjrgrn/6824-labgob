# labgob

This library provides defensive wrappers around `encoding/gob`'s `Encoder`, `Decoder`, `Register`, and `RegisterName`.

These wrappers include safety checks, specifically:

- `checkValue()` warns the user if a struct type or field name starts with a lowercase letter. This library is designed for use with RPC systems, where sending unexported (lowercase) fields can cause silent failures or misbehavior.

- `checkDefault()` runs before decoding data. Its purpose is to warn if you are about to decode into a variable that already contains non-zero values. This warning is necessary because Go's `gob` decoder does not overwrite existing values with zero values if the field is missing from the encoded data (e.g., if it was omitted because it held the zero value). Instead, the decoder leaves the destination field unchanged, which can lead to stale data.

This library originates from the [MIT 6.824 Distributed Systems](https://pdos.csail.mit.edu/6.824/) laboratory. It has been extracted from the course source code and packaged as an external library to better serve my projects.
