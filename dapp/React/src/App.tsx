import React, { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import { usdtAddress, sonmAddress, receiverAddress } from './constants';
import { erc20Abi, receiverAbi } from './abis';

declare global {
  interface Window {
    ethereum: any;
  }
}

const App: React.FC = () => {
  const [amount, setAmount] = useState<number>(0);
  const [minMintAmount, setMinMintAmount] = useState<number | null>(null);

  // Fetch minimum mint amount on component mount
  useEffect(() => {
    const fetchMinMintAmount = async () => {
      if (!window.ethereum) {
        console.error('MetaMask is not installed.');
        return;
      }

      try {
        const provider = new ethers.BrowserProvider(window.ethereum);
        const receiverContract = new ethers.Contract(receiverAddress, receiverAbi, provider);

        const minMintAmountWei = await receiverContract.getMinMintAmount();
        const minMintAmountValue = ethers.formatUnits(minMintAmountWei, 6); // Assuming USDT has 6 decimals
        setMinMintAmount(Number(minMintAmountValue));
      } catch (error) {
        console.error('Error fetching minimum mint amount:', error);
      }
    };

    fetchMinMintAmount();
  }, []);

  const mintForUSDT = async () => {
    if (!window.ethereum) {
      alert('MetaMask is not installed. Redirecting to MetaMask installation page.');
      window.location.href = 'https://metamask.io/download.html';
      return;
    }

    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      await provider.send('eth_requestAccounts', []);
      const signer = await provider.getSigner();

      const usdtContract = new ethers.Contract(usdtAddress, erc20Abi, signer);
      const receiverContract = new ethers.Contract(receiverAddress, receiverAbi, signer);

      const amountInWei = ethers.parseUnits(amount.toString(), 6);

      // Approve USDT transfer
      const approval = await usdtContract.approve(receiverAddress, amountInWei);
      await approval.wait();

      // Mint SONM tokens
      const tx = await receiverContract.mintForUSDT(amountInWei);
      await tx.wait();

      alert('Minting successful!');
    } catch (error: unknown) {
      if ((error as any).code === 4001) {
        alert('Transaction cancelled by user.');
      } else {
        console.error(error);
        alert('An error occurred. Please try again.');
      }
    }
  };

  return (
    <div>
      <h1>USDT Receiver DApp</h1>
      <p>
        {minMintAmount !== null
          ? `Minimum mint amount: ${minMintAmount} USDT`
          : 'Fetching minimum mint amount...'}
      </p>
      <input
        type="number"
        value={amount}
        onChange={(e) => setAmount(Number(e.target.value))}
        placeholder="Amount"
      />
      <button
        onClick={mintForUSDT}
        disabled={minMintAmount !== null && amount < minMintAmount}
      >
        Mint SONM for USDT
      </button>
    </div>
  );
};

export default App;
