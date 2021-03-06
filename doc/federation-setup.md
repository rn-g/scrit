Setting up a federation of Scrit mints
--------------------------------------

This walkthrough shows how one would set up a federation of 3 Scrit
mints.

First each of the three mints has to generate their mint identity key:

    $ scrit-mint keygen

Afterwards they all display their identity key and send it to the
federation coordinator (one of the mints should manage the
[Codechain](https://github.com/frankbraun/codechain) for the new Scrit
federation).

    $ scrit-mint identity

Typical output looks like this:

    ed25519-fZLPEvdKwhvxU_asrnqbR9t1PV0FukT71f1iwExX_ic

Let's say we have the following three mint identity keys:

    ed25519-vVqGX7eEyH5DNxO_UHm2k8iJAvf-NNv2g1UbZnTnu44
    ed25519-boVnUGMNKkI1Pe72m8Kf_9KljL4DBvsOGxbr1wi9flo
    ed25519-er0Phn1PjBzbz3gBUEbFQUIbexZxufELZyzCyfT4A5U

Now we can setup the federation (2-of-3):

    $ scrit-gov start -m 2 -n 3 ed25519-vVqGX7eEyH5DNxO_UHm2k8iJAvf-NNv2g1UbZnTnu44 ed25519-boVnUGMNKkI1Pe72m8Kf_9KljL4DBvsOGxbr1wi9flo ed25519-er0Phn1PjBzbz3gBUEbFQUIbexZxufELZyzCyfT4A5U

Define first DBC types (in denominations of 1, 2, and 5 EUR):

    $ scrit-gov dbctype add -currency EUR -amount 100000000
    $ scrit-gov dbctype add -currency EUR -amount 200000000
    $ scrit-gov dbctype add -currency EUR -amount 500000000

Now each of the three mints creates their key lists:

    $ scrit-mint keylist create -desc mint_name https://mint.example.com

Define the second signing epoch

    $ scrit-gov epoch add

Now each of the three mints extend their key lists (can be skipped if
`scrit-gov epoch add` is called _before_ `scrit-mint keylist create`):

    $ scrit-mint keylist extend

Now the configuration directory for the federation looks like this:

    $ tree
    .
    ├── federation.json
    └── mints
        ├── ed25519-boVnUGMNKkI1Pe72m8Kf_9KljL4DBvsOGxbr1wi9flo.json
        ├── ed25519-er0Phn1PjBzbz3gBUEbFQUIbexZxufELZyzCyfT4A5U.json
        └── ed25519-vVqGX7eEyH5DNxO_UHm2k8iJAvf-NNv2g1UbZnTnu44.json

With `federation.json` generally defining the network (including the
siginig epochs, DBC types, and mints) and a key list for each individual
mint in the `mints` subdirectory.

This directory structure would typically be checked into a Codechain
with the mint identity keys as signers and the same quorum for changes
(2-of-3 in this example).

This Scrit federation config directory is then distributed and
automatically updated via Codechain's `secpkg` tool to Scrit wallet
users.

To test a configuration a user would execute the following in the
configuration directory:

    $ scrit-engine validateconf

To be continued...
