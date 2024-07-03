import React, { useState } from 'react';
import { ethers } from 'ethers';
import { usdtAddress, sonmAddress, receiverAddress } from './constants';
import { erc20Abi, receiverAbi } from './abis';

const App: React.FC = () => {
  const [amount, setAmount] = useState<number>(0);

  const mintForUSDT = async () => {
    if (!window.ethereum) return alert('Please install MetaMask');

    const provider = new ethers.JsonRpcProvider(window.ethereum);
    await provider.send('eth_requestAccounts', []);
    //const signer = provider.getSigner();
    
    const usdtContract = new ethers.Contract(usdtAddress, erc20Abi, provider);
    const receiverContract = new ethers.Contract(receiverAddress, receiverAbi, provider);

    const amountInWei = ethers.parseUnits(amount.toString(), 6);

    // Approve USDT transfer
    const approval = await usdtContract.approve(receiverAddress, amountInWei);
    await approval.wait();

    // Mint SONM tokens
    const tx = await receiverContract.mintForUSDT(amountInWei);
    await tx.wait();

    alert('Minting successful!');
  };

  const burnAndSendUSDT = async () => {
    if (!window.ethereum) return alert('Please install MetaMask');

    const provider = new ethers.JsonRpcProvider(window.ethereum);
    await provider.send('eth_requestAccounts', []);
    
    const sonmContract = new ethers.Contract(sonmAddress, erc20Abi, provider);
    const receiverContract = new ethers.Contract(receiverAddress, receiverAbi, provider);

    const amountInWei = ethers.parseUnits(amount.toString(), 6); 

    // Approve SONM burn
    const approval = await sonmContract.approve(receiverAddress, amountInWei);
    await approval.wait();

    // Burn SONM tokens and send USDT
    const tx = await receiverContract.burnAndSendUSDT(amountInWei);
    await tx.wait();

    alert('Burning successful!');
  };

  return (
    <div>
      <h1>USDT Receiver DApp</h1>
      <input
        type="number"
        value={amount}
        onChange={(e) => setAmount(Number(e.target.value))}
        placeholder="Amount"
      />
      <button onClick={mintForUSDT}>Mint SONM for USDT</button>
      <button onClick={burnAndSendUSDT}>Burn SONM for USDT</button>
    </div>
  );
};

export default App;