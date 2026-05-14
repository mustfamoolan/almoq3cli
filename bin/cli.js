#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const os = require('os');

// Determine binary name based on OS
const platform = os.platform();
let binaryName = 'almoq3';

if (platform === 'win32') {
    binaryName = 'almoq3.exe';
}

// Path to the binary (assumed to be in the same folder as this script or project root)
const binaryPath = path.join(__dirname, '..', binaryName);

const args = process.argv.slice(2);

const child = spawn(binaryPath, args, {
    stdio: 'inherit',
    shell: false
});

child.on('error', (err) => {
    console.error(`❌ Failed to start almoq3 CLI: ${err.message}`);
    console.log(`Please ensure the ${binaryName} binary exists at: ${binaryPath}`);
    process.exit(1);
});

child.on('exit', (code) => {
    process.exit(code);
});
