Invoice-generator
============

invoice-generator is simple support tool for making invoice.
Translate english product name to japanese product name on amazon.

Usage
-----

    > ./main -wishlist [wishlist ID]

For example,

    > ./main -wishlist 26NVSI4REFIQY

Output: 

```json
...
{"ID":"4797374144","JpName":"数学ガールの秘密ノート/式とグラフ","EnName":"Sugaku Girl No Himitsu No Note / Shiki To Graph"}
...
```