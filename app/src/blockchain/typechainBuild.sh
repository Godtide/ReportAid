#!/bin/sh

../../node_modules/.bin/typechain --target ethers  --outDir ./typechain ./build/contracts/*json
