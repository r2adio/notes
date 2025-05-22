# CAT FILE:

- cat-file, git built-in command, allows to see the content of a commit w/o needing to futz around with object files directly.

<!-- NOTE: `git cat-file -p 5f28`
    - this outputs the tree, parent, author, gpgsig (if generated), commit message
    - output:

    ❯ git cat-file -p 5f28
    tree faa2158db05d8a9545b103ec594c9cd3b730a5a2
    parent 7f126534076c15552c5df11c71a150ee89891809
    author GauraVashisth <miya.musashi.moto1645@gmail.com> 1747718145 +0530
    committer GauraVashisth <miya.musashi.moto1645@gmail.com> 1747718145 +0530
    gpgsig -----BEGIN PGP SIGNATURE-----

     iQIzBAABCAAdFiEEy2WUdbWgnSa2QIt89Tj+OoYtEWkFAmgsEAEACgkQ9Tj+OoYt
     EWlupg/9FtqWRj8OVyYDliMoQpMPlcGXOc5du=1EOzjw/D1y6ddwZ0VN4flMLx7f
     ubTsHT+uvqBRBxp3CXnrjJm+snZH9HkD8q+O+uIk4+RgWEhetcj68+wIy6i9Prt8
     +Cs63X/6D+WjQhA6HdWPG5Sqeu6jr6fNhZmkxCssBAjwoRjUmffGcU6YcXr4uZZo
     VZ4g9G35ppIv3aSz10gTFdAb/hLwbEjmQtDXn2coZFnNN6PmRf4VXhaQ7LGn1cNL
     GCtw3ezA16mglk99fVbZP4V7VZb2T7UkiTBfWC3N1sV5UCGDP+JDlquXrAyAZmgv
     pxn0UTVmY6OyLCMPzAkQz00CnQvm0mqbdFY9ZMU4DXDEGr6WNBk0tYoicBEwG3K0
     c2JVX2psRtMrSaYx1axYjPS9cg9aSNcnte87m3Tpikku1L+PDYKtVlI2MyPdqU49
     vI5+QuQPmIEA-+YBnuRkaC2Ztf9+6pJe806v1Nq5hsDK6TTbQnpSUimpJpNRUYZ1
     QFGo/iKXmQVUczlFChKe7k4JJphEhpwjBMol5Th1/7Zhj/M2UkIa6DU84lHQSmgU
     mQCuxWeHykcYJYJo+jwgILlsJLBSTKaAbzw9Ht8qF+v6P42zxyHCWojW49F3Y92h
     eUhG1w5JS/z5dy8wN1Eh6P7GOMdksv0qxenVyGBBvY/4E0Kh8a4=
     =eDJg
     -----END PGP SIGNATURE-----

    testing md comments
-->

## TREE AND BLOBS
