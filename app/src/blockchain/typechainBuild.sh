#!/bin/sh

../../node_modules/.bin/typechain --target=ethers  --outDir ./build/contracts/typechain ./build/contracts/*json
